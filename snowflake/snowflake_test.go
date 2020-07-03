package snowflake

import (
	"fmt"
	"sync"
	"testing"
)

type BMap struct {
	data map[uint64]bool
	lck  sync.Mutex
}

func TestNewSF(t *testing.T) {
	// 全局偏移量
	var gts int64 = 1593770436295
	InitSnowSerial(gts)
	m := new(BMap)
	m.data = make(map[uint64]bool)

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			for i := 0; i < 100000; i++ {
				// time.Sleep(time.Millisecond)
				var gid = GID(1)
				m.lck.Lock()
				if _, ok := m.data[gid]; ok {
					m.lck.Unlock()
					panic(fmt.Sprintf("duplicate ID occurred:%d", gid))
				} else {
					m.data[gid] = true
				}
				m.lck.Unlock()
			}
			wg.Done()
		}(&wg)
	}
	wg.Wait()
}
