package main

import (
	"log"
	"runtime"
	"time"
)

type Sample struct {
	Name string
}

func main() {

	for i := 0; i < 2; i++ {
		go test()
	}
	time.Sleep(1* time.Second)

	var cnt = 1024
	var m = make([][]int8, cnt)
	for ; cnt > 0; cnt-- {
		m[cnt-1] = make([]int8, 1024 * 4) // 4kB
	}
	//time.Sleep(2* time.Second)
}

func test() {
	sobj := &Sample{"test"}

	// 为对象设置回收时的回调
	runtime.SetFinalizer(sobj, func(o *Sample) {
		log.Println("sample object is collected!")
	})

	log.Printf(sobj.Name)
}
