package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"runtime/pprof"

	"strings"
	"syscall"
	"time"

	"github.com/urfave/cli"
	"gopkg.in/godo.v2/glob"
)

//    use system file change time to determine wheather file changed.
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
	// Command output in log file when Mode=background
	LogFile string `json:"log_file"`
	// regist commands
	Commands map[string]string `json:"commands"`
	// Command excution path
	CmdPath string `json:"cmd_path"`

	// excludes files by patterns or filename
	Excludes []string `json:"excludes"`

	// file changed identify map
	datas map[string]changeIdentify
	// scan times
	scanCounter int
	// scan ticker
	ScanTicker int64 `json:"scan_ticker"`
}

var needCPUProfile bool

func main() {

	var (
		configFile string
	)

	// parse cli command
	app := cli.NewApp()
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "file",
			Value:       "",
			Required:    true,
			Usage:       "specify the configuration file.",
			Destination: &configFile,
		},
		cli.BoolFlag{
			Name:        "prof",
			Required:    false,
			EnvVar:      "",
			Usage:       "whether use pprof to analysis performance.",
			Destination: &needCPUProfile,
		},
	}

	app.Action = func(c *cli.Context) error {
		if c.Bool("prof") {
			needCPUProfile = true
		}

		Setup(configFile)
		return nil
	}

	app.Run(os.Args)
	//  godaemon.MakeDaemon(&godaemon.DaemonAttr{})
}

// Setup run service.
func Setup(cf string) {
	//   是否启用pprof分析程序cpu使用
	if needCPUProfile {
		log.Println("Start CPU profile logger.")

		f, err := os.Create("./ghrun_cpu_profile")
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}
	// read config file
	var (
		data []byte
		err  error
	)
	if data, err = ioutil.ReadFile(cf); err != nil {
		panic(err)
	}
	var config string
	config = string(data)
	cbs := parseCfg(config)
	// log.Println(cbs)

	var ctx context.Context
	ctx = context.Background()
	ctx, cancel := context.WithCancel(ctx)

	defer func() {
		log.Println("stop children process..")
		cancel()
		time.Sleep(1 * time.Second)
		log.Println("Stopped.")
	}()

	for _, cb := range cbs {
		ListenBlock(ctx, cb)
	}

	var ch = make(chan os.Signal)

	signal.Notify(ch, syscall.SIGINT, syscall.SIGKILL, syscall.SIGTERM, syscall.SIGHUP)
	signal := <-ch // blocking
	log.Println("get shutdown signal:", signal)
}

func parseCfg(s string) []CfgBlock {
	var cbs []CfgBlock
	cbs = make([]CfgBlock, 0)
	if err := json.Unmarshal([]byte(s), &cbs); err != nil {
		log.Fatalf("Config block parse failed: %s", err)
	}
	// log.Printf("%+v \n", cbs)
	for idx := range cbs {
		cbs[idx].datas = make(map[string]changeIdentify)
	}
	return cbs
}

// ListenBlock 监听配置的block
// 改为了传值, 因为不允许多个协程读写一个map(不考虑线程安全,配置无修改操作)
func ListenBlock(ctx context.Context, cb CfgBlock) {

	go func() {
		var duration int64
		if cb.ScanTicker < 1 {
			duration = int64(5 * time.Second)
		} else {
			duration = cb.ScanTicker * int64(time.Second)
		}
		log.Println("scan_ticker =", duration)
		var cancel context.CancelFunc
	END:
		for {
			// log.Println("counter run..")

			if isUpdated, c := cb.readNestedDirs(ctx, cancel); isUpdated && c != nil {
				cancel = c
			}

			select {
			case <-ctx.Done():
				cancel()
				time.Sleep(1 * time.Second)
				break END
			default:
				time.Sleep(time.Duration(duration))
			}
		}

	}()

	// // !!!!! 使用ticker的for-select goroutine会大幅度提高cpu占用率 !!!!!
	// ticker := time.NewTicker(time.Duration(10 * time.Second))

	// go func(ticker *time.Ticker) {
	// 	var cancel context.CancelFunc

	// END:
	// 	for {
	// 		select {
	// 		case <-ticker.C:
	// 			//定时执行
	// 			log.Println("counter run..")
	// 			if isUpdated, c := cb.readNestedDirs(ctx, cancel); isUpdated && c != nil {
	// 				cancel = c
	// 			}
	// 		case <-ctx.Done():
	// 			cancel()
	// 			log.Println("counter run..")
	// 			time.Sleep(1 * time.Second)
	// 			ticker.Stop()
	// 			break END
	// 		default:
	// 		}
	// 	}
	// }(ticker)

	return
}

func (cb *CfgBlock) readNestedDirs(ctx context.Context, stopLast context.CancelFunc) (isUpdated bool, cancel context.CancelFunc) {
	defer func() {
		// 目录扫描次数up
		cb.scanCounter++
		// log.Println("scan dir --- ", cb.Path, "====>", cb.scanCounter)
	}()

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
					// log.Printf("[%s]match pattern: %+v\n", p, regexp)
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
				log.Printf("修改文件[%s]...\n", p)
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
			log.Printf("新增文件[%s]...\n", p)
		}
		return nil
	})

	if total < len(cb.datas) {
		isUpdated = true
		for path, ci := range cb.datas {
			if ci.SyncTime != ctime {
				delete(cb.datas, path)
				log.Printf("删除文件[%s]...\n", path)
			}
		}
	}

	if isUpdated && cb.scanCounter > 0 {
		// log.Printf(">>> listen path:(%s) is changed!\n", cb.Path)
		if stopLast != nil {
			stopLast()
			time.Sleep(10 * time.Microsecond)
		}
		ctx, cancel = context.WithCancel(ctx)
		cb.Exec(ctx)
	}

	return
}

// Exec executing config block in goroutine
func (cb *CfgBlock) Exec(ctx context.Context) {

	var shellCmd string = cb.Commands[cb.Execute]
	var mode string = cb.Mode
	if !Contains([]string{"foreground", "background"}, mode) {
		log.Fatalf("config[%s]err: mode(%s) not supported.\n", cb.Name, cb.Mode)
	}

	cmd := exec.CommandContext(ctx, "/bin/sh", "-c", shellCmd)

	if cb.CmdPath != "" {
		cmd.Dir = cb.CmdPath
	} else {
		// 进入监听的目录来执行
		cmd.Dir = cb.Path
	}
	// !! 设置进程组属性
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}

	// if foreground can use output
	if cb.Mode == "background" {
		if cb.LogFile == "" {
			log.Fatalf("backgound mode must setting a log_file first!")
		}
		if f, err := os.OpenFile(cb.LogFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666); err != nil {
			log.Fatalln(err)
		} else {
			cmd.Stdout = f
		}
	} else {
		cmd.Stdout = os.Stdout
	}
	// log.Println(cmd.Path, cmd.Args)

	errCh := make(chan error, 1)
	pgidCh := make(chan int, 1)

	go func(cmd *exec.Cmd) {
		if err := cmd.Start(); err != nil {
			errCh <- err
		}
		log.Println(cmd.Process.Pid, "creating CMD Process.")
		if pgid, err := syscall.Getpgid(cmd.Process.Pid); err != nil {
			log.Fatal(err)
		} else {
			pgidCh <- pgid
		}

		go func() {
			errCh <- cmd.Wait()
			log.Println(cmd.Process.Pid, "CMD Process done.")
		}()
	}(cmd)

	var pgid int = <-pgidCh

	go func() {
	END:
		for {
			select {
			case <-ctx.Done():
				// log.Println(cmd.Process.Pid, ">>>>>> stop cmd process.")
				// cmd.Process.Kill()
				// @see: https://stackoverflow.com/questions/22470193/why-wont-go-kill-a-child-process-correctly
				// @see: https://stackoverflow.com/questions/24982845/process-kill-on-child-processes
				// cmd.Process.Kill()无法关闭子进程, 所以设置pgid,通过关闭进程组来正确关闭
				syscall.Kill(-pgid, 15)
				break END
			case err := <-errCh:
				if err != nil {
					log.Println(err)
				}
				break END
			}
		}
		// log.Println(cmd.Process.Pid, "for-select ctr exited.")
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
