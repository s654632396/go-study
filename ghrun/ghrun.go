package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"strings"
	"sync"
	"syscall"
	"time"

	"gopkg.in/godo.v2/glob"
)

type changeIdentify struct {
	SyncTime    time.Time
	changeValue string
}

// CfgBlock reading config block from config file
type CfgBlock struct {
	// block name
	Name string `json:"name"`
	// listen path
	Path string `json:"path"`
	// execute specific command when path which listened changed.
	Execute string `json:"execute"`
	// execute mode ,valid values: forceground(前台) background(后台,用于启动守护进程之类的)
	Mode string `json:"mode"`
	// regist commands
	Commands map[string]string `json:"commands"`

	// excludes files by patterns or filename
	Excludes []string `json:"excludes"`

	// file changed identify map
	datas map[string]changeIdentify
	// scan times
	scanCounter int
}

func main() {
	var testCfg = `
	 [
	 	{
	 		"name" : "test",
	 		"path" : "/opt/practice/go_server",
	 		"execute" : "run",
	 		"commands": {
	 			"run" :   "go run /opt/practice/go_server/main.go",
	 			"build": "go build /opt/practice/go_server/main.go"
	 		},
	 		"mode" : "background",
	 		"excludes": [
	 			"test/*",
	 			"{main}",
	 			".git/*"
	 		]
	 	},
	 	{
	 		"name" : "test2",
	 		"path" : "/opt/practice/goapp",
	 		"execute" : "run",
	 		"commands": {
	 			"run" : "go run /opt/practice/goapp/test_btree.go"
	 		},
	 		"mode" : "background",
	 		"excludes": []
	 	}
	 ]
	 `
	// 	var testCfg = `
	// [
	// 	{
	// 		"name" : "test",
	// 		"path" : "/opt/practice/go_server",
	// 		"execute" : "run",
	// 		"commands": {
	// 			"run" :   "go run /opt/practice/go_server/main.go",
	// 			"build": "go build /opt/practice/go_server/main.go"
	// 		},
	// 		"mode" : "background",
	// 		"excludes": [
	// 			"test/*",
	// 			"{main}",
	// 			".git/*"
	// 		]
	// 	}
	// ]
	// `

	cbs := parseCfg(testCfg)

	var ctx context.Context
	ctx = context.Background()
	ctx, cancel := context.WithCancel(ctx)

	for _, cb := range cbs {
		ListenBlock(ctx, cb)
	}

	var ch = make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch // blocking
	fmt.Println("Read of interrupt signal...")
	cancel()
	time.Sleep(1 * time.Second)
	fmt.Println("Stopped.")
}

func parseCfg(s string) []CfgBlock {
	var cbs []CfgBlock
	cbs = make([]CfgBlock, 0)
	if err := json.Unmarshal([]byte(s), &cbs); err != nil {
		log.Fatalf("Config block parse failed: %s", err)
	}
	for idx := range cbs {
		cbs[idx].datas = make(map[string]changeIdentify)
	}
	return cbs
}

// ListenBlock 监听配置的block
// 改为了传值, 因为不允许多个协程读写一个map
func ListenBlock(ctx context.Context, cb CfgBlock) {
	ticker := time.NewTicker(time.Duration(1000 * time.Millisecond))

	var lock sync.Mutex

	go func(ticker *time.Ticker) {
		var cancel context.CancelFunc
	END:
		for {
			select {
			case <-ticker.C:
				//定时执行
				lock.Lock()
				// fmt.Println("counter run..")
				if isUpdated, c := cb.readNestedDirs(ctx, cancel); isUpdated && c != nil {
					cancel = c
				}
				lock.Unlock()
			case <-ctx.Done():
				fmt.Println("stop ticker && call cancel func..")
				cancel()
				ticker.Stop()
				break END
			}
		}
	}(ticker)

	return
}

func (cb *CfgBlock) readNestedDirs(ctx context.Context, stopLast context.CancelFunc) (isUpdated bool, cancel context.CancelFunc) {
	isUpdated = false
	cancel = nil

	var total int = 0
	var ctime time.Time = time.Now()
	_, regexps, _ := glob.Glob(cb.Excludes)

	filepath.Walk(cb.Path, func(p string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if strings.HasPrefix(p, cb.Path) {
			shortP := strings.Trim(p[len(cb.Path):], "/")
			for _, regexp := range regexps {
				if regexp.MatchString(shortP) {
					// fmt.Printf("[%s]match pattern: %+v\n", p, regexp)
					return nil
				}
			}
		}

		// files in activate listened
		total++
		curtimeStr := info.ModTime().Format(string(time.RFC3339Nano))
		if pv, ok := cb.datas[p]; ok {
			if pv.changeValue != curtimeStr {
				// is updated !
				cb.datas[p] = changeIdentify{ctime, curtimeStr}
				isUpdated = true
				fmt.Printf("修改文件[%s]...\n", p)
			} else {
				pv.SyncTime = ctime
				cb.datas[p] = pv
			}
		} else {
			cb.datas[p] = changeIdentify{ctime, curtimeStr}
			isUpdated = true
			if cb.scanCounter == 0 {
				return nil
			}
			fmt.Printf("新增文件[%s]...\n", p)
		}

		return nil
	})

	if total < len(cb.datas) {
		isUpdated = true
		for path, ci := range cb.datas {
			if ci.SyncTime != ctime {
				delete(cb.datas, path)
				fmt.Printf("删除文件[%s]...\n", path)
			}
		}
	}

	cb.scanCounter++
	if isUpdated {
		// fmt.Printf(">>> listen path:(%s) is changed!\n", cb.Path)
		if stopLast != nil {
			stopLast()
			time.Sleep(1 * time.Second)
		}
		ctx, cancel = context.WithCancel(ctx)
		cb.Exec(ctx)
	}

	return
}

// Exec executing config block in goroutine
func (cb *CfgBlock) Exec(ctx context.Context) {

	// parse run
	var unparsed string = cb.Commands[cb.Execute]
	// fmt.Println("需要执行命令: ", cb.Execute, cmd)
	var mode string = cb.Mode
	if !Contains([]string{"foreground", "background"}, mode) {
		log.Fatalf("config[%s]err: mode(%s) not supported.\n", cb.Name, cb.Mode)
	}

	// parse cmd
	var parsed []string
	parsed = strings.SplitN(unparsed, " ", 3)

	// fmt.Printf("%#v \n", parsed)

	cmd := exec.CommandContext(ctx, parsed[0], parsed[1], parsed[2])

	// if foreground can use output
	// TODO
	if cb.Mode == "foreground" {
		cmd.Stdout = os.Stdout
	}

	// !! 默认进入监听的目录来执行
	cmd.Dir = cb.Path
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}

	// fmt.Println(ex.Path, ex.Args)

	errCh := make(chan error, 1)
	pgidCh := make(chan int, 1)

	go func(cmd *exec.Cmd) {
		if err := cmd.Start(); err != nil {
			errCh <- err
		}
		fmt.Println(cmd.Process.Pid, "creating CMD Process.")
		if pgid, err := syscall.Getpgid(cmd.Process.Pid); err != nil {
			log.Fatal(err)
		} else {
			pgidCh <- pgid
		}

		errCh <- cmd.Wait()
		fmt.Println(cmd.Process.Pid, "CMD Process done.")
	}(cmd)

	var pgid int = <-pgidCh

	go func() {
	END:
		for {
			select {
			case <-ctx.Done():
				fmt.Println(cmd.Process.Pid, ">>>>>> stop cmd process.")
				// cmd.Process.Kill()
				// @see: https://stackoverflow.com/questions/22470193/why-wont-go-kill-a-child-process-correctly
				// cmd.Process.Kill()无法关闭子进程, 所以设置pgid,通过关闭进程组来正确关闭
				syscall.Kill(-pgid, 15)
				break END
			case err := <-errCh:
				fmt.Errorf("CMD Err: ", err)
				break END
			}
		}
		fmt.Println(cmd.Process.Pid, "for-select ctr exited.")
	}()

}

// Contains check string exists in []string
func Contains(a []string, x string) bool {
	for _, elem := range a {
		if elem == x {
			return true
		}
	}
	return false
}
