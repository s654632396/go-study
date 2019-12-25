package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

//读写锁
var rwLock sync.RWMutex

func testRWLock() {
	var a map[int]int
	a = make(map[int]int, 5)
	a[8] = 10
	a[3] = 10
	a[2] = 10
	a[1] = 10
	a[18] = 10
	for i := 0; i < 2; i++ {
		go func(b map[int]int) {
			rwLock.Lock()
			b[8] = rand.Intn(100)
			rwLock.Unlock()
		}(a)
	}
	for i := 0; i < 10; i++ {
		go func(b map[int]int) {
			rwLock.RLock() //读锁
			fmt.Println(a)
			rwLock.RUnlock()
		}(a)
	}
	time.Sleep(time.Second * 2)

}
func main() {

	testRWLock()
	//读多写少的时候，用读写锁
}
