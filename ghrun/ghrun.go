package main

import (
	"context"
	"fmt"
	"gopkg.in/godo.v2/glob"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

type changeIdentify struct {
	SyncTime    time.Time
	changeValue string
}

// CfgBlock reading config block from config file
type CfgBlock struct {
	// block name
	name string
	// listen path
	path string
	// execute specific command when path which listened changed.
	execute string
	// execute mode ,valid values: forceground(前台) background(后台,用于启动守护进程之类的)
	mode string
	// regist commands
	commands map[string]string

	// excludes files by patterns or filename
	excludes []string

	// file changed identify map
	datas map[string]changeIdentify
	// scan times
	scanCounter int
	// if background...
	runningCtx context.CancelFunc
}

func main() {

	cb := &CfgBlock{
		name:    "test",
		path:    "/opt/practice/go_server",
		execute: "run",
		commands: map[string]string{
			"run":   "go run /opt/practice/go_server/main.go",
			"build": "go build /opt/practice/go_server/main.go",
		},
		mode:  "background",
		datas: make(map[string]changeIdentify),
		excludes: []string{
			"test/*",
			"{main}",
		},
	}
	//fmt.Printf("%#v\n", cb)

	ch := ListenBlock(cb)

	for {
		time.Sleep(30 * time.Second)
	}
	fmt.Println("for loop is done")
	ch <- 1
}

// ListenBlock 监听配置的block
func ListenBlock(cb *CfgBlock) (ch chan int) {
	ticker := time.NewTicker(time.Duration(1000 * time.Millisecond))
	ch = make(chan int, 1)

	var lock sync.Mutex

	go func(ticker *time.Ticker) {
		for {
			select {
			case <-ticker.C:
				//定时执行
				lock.Lock()
				// fmt.Println("counter run..")
				cb.readNestedDirs()
				lock.Unlock()
			case <-ch:
				fmt.Println("stop ticker..")
				ticker.Stop()
				break
			}
		}
	}(ticker)

	return
}

func (cb *CfgBlock) readNestedDirs() (isUpdated bool) {
	isUpdated = false

	var total int = 0
	var ctime time.Time = time.Now()
	_, regexps, _ := glob.Glob(cb.excludes)

	// fmt.Printf("%+v, %+v\n", regexps, cb.excludes)
	// return

	filepath.Walk(cb.path, func(p string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if strings.HasPrefix(p, cb.path) {
			shortP := strings.Trim(p[len(cb.path):], "/")
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

	if isUpdated {
		// fmt.Printf(">>> listen path:(%s) is changed!\n", cb.path)
		cb.exec()
	}

	cb.scanCounter++

	return isUpdated
}

func (cb *CfgBlock) exec() {
	// parse run
	var cmd string = cb.commands[cb.execute]
	fmt.Println("需要执行命令: ", cb.execute, cmd)
	var mode string = cb.mode
	fmt.Println("运行模式: ", mode)

	// parse cmd
	var parsed []string
	parsed = strings.SplitN(cmd, " ", 3)

	// fmt.Printf("%#v \n", parsed)

	// cancel last Cmd
	if cb.runningCtx != nil {
		fmt.Printf("Stop latest running...\n")
		cb.runningCtx()
	}

	ctx := context.Background()
	ctx, cb.runningCtx = context.WithCancel(ctx)

	// ex := exec.Command("bash", "-c", cmd)
	ex := exec.CommandContext(ctx, parsed[0], parsed[1], parsed[2])
	// if forceground can use output
	// TODO
	// ex.Stdout = os.Stdout

	// !! 默认进入监听的目录来执行
	ex.Dir = cb.path

	fmt.Println(ex.Path, ex.Args)

	exChan := make(chan error)
	go func(ex *exec.Cmd, exChan chan error) {
		exChan <- ex.Run()
		fmt.Printf("execute command, PID=%d\n", ex.Process.Pid)
	}(ex, exChan)

	for {
		select {
		case <-time.After(1 * time.Second):
			return
		case err := <-exChan:
			if err != nil {
				fmt.Printf("ex.Run() failed with %#v ...\n", err)
				return
			}
			fmt.Println("ex run done.")
		}
	}
}
