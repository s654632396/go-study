package snowflake

import (
	"sync"
	"testing"
)

type BMap struct {
	data     map[int64]int
	lck      sync.Mutex
	dupCnt   int
	totalCnt int
}

func initSID() *SnowID {
	// 初始化
	// 全局偏移量
	var gts int64 = 1593770436295
	var mID int64 = 1
	return NewSnowID(gts, mID)
}

func TestNewSF(t *testing.T) {
	sid := initSID()
	m := new(BMap)
	m.data = make(map[int64]int)

	var wg sync.WaitGroup
	pn := 16
	loop := 4e7 / pn
	wg.Add(pn)

	for i := 0; i < pn; i++ {
		go func(wg *sync.WaitGroup) {
			for i := 0; i < loop; i++ {
				var gid = sid.GID()
				m.lck.Lock()
				if _, ok := m.data[gid]; ok {
					m.data[gid] += 1
					m.dupCnt += 1
					t.Fail()
				} else {
					m.data[gid] = 1
				}
				m.totalCnt += 1
				m.lck.Unlock()
			}
			wg.Done()
		}(&wg)
	}
	wg.Wait()

	t.Log("dup:", m.dupCnt, "/", m.totalCnt)
}

func BenchmarkNewSnowID(b *testing.B) {
	sid := initSID()
	var wg sync.WaitGroup
	pn := 16
	loop := 5e5 / pn
	wg.Add(pn)
	for i := 0; i < pn; i++ {
		go func(wg *sync.WaitGroup) {
			for i := 0; i < loop; i++ {
				var _ = sid.GID()
			}

			wg.Done()
		}(&wg)
	}
	wg.Wait()
}
