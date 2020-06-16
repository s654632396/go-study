package pool

import (
	"runtime"
	"sync"
	"testing"
)

var procs = runtime.GOMAXPROCS(-1)

func BenchmarkPool_Pool(b *testing.B) {
	var p Pool
	p.New = func() interface{} {
		return 1
	}
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			p.Put(1)
			p.Get()
		}
	})

}

func BenchmarkPool_ShardPool(b *testing.B) {
	var p = NewShardPool(procs, func() interface{} {
		return 1
	})

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			p.Put(1)
			p.Get()
		}
	})

}

func BenchmarkPool_SyncPool(b *testing.B) {
	var p sync.Pool
	p.New = func() interface{} {
		return 1
	}

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			p.Put(1)
			p.Get()
		}
	})
}

func BenchmarkPool_ChanPool(b *testing.B) {
	var p = NewChanPool(10000)
	p.New = func() interface{} {
		return 1
	}

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			p.Put(1)
			p.Get()
		}
	})
}

func BenchmarkPoolOverflow_Pool(b *testing.B) {
	var p Pool
	p.New = func() interface{} {
		return 1
	}

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			for b := 0; b < 1000; b++ {
				p.Put(1)
			}
			for b := 0; b < 1000; b++ {
				p.Get()
			}
		}
	})
}

func BenchmarkPoolOverflow_ShardPool(b *testing.B) {
	var p = NewShardPool(procs, func() interface{} {
		return 1
	})

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			for b := 0; b < 1000; b++ {
				p.Put(1)
			}
			for b := 0; b < 1000; b++ {
				p.Get()
			}
		}
	})
}

func BenchmarkPoolOverflow_SyncPool(b *testing.B) {
	var p sync.Pool
	p.New = func() interface{} {
		return 1
	}

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			for b := 0; b < 1000; b++ {
				p.Put(1)
			}
			for b := 0; b < 1000; b++ {
				p.Get()
			}
		}
	})
}

func BenchmarkPoolOverflow_ChanPool(b *testing.B) {
	var p = NewChanPool(10000)
	p.New = func() interface{} {
		return 1
	}

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			for b := 0; b < 1000; b++ {
				p.Put(1)
			}
			for b := 0; b < 1000; b++ {
				p.Get()
			}
		}
	})
}
