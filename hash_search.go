package main

//@see: https://segmentfault.com/a/1190000022288951

//noinspection GoUnresolvedReference
import (
	"bytes"
	"context"
	"fmt"
	"github.com/OneOfOne/xxhash"
	"log"
	"math"
	"runtime"
	"strings"
	"sync"
	"time"
)

const (
	DefaultCap uint64 = 1 << 4
)

type item struct {
	key   string      // 原生key
	value interface{} // 节点存储值
	next  *item       // 指向下一个Item, 为nil则是末端节点
}

type HashMap struct {
	len          uint64  // 当前长度length
	cap          uint64  // 仓库容量capacity
	repo         []*item // 数据仓库
	expendFactor uint64  // 当仓库大小达到该值时,进行扩容
	lock         sync.Mutex
	wait         bool
	waitCh       chan int
}

// NewHashMap
func NewHashMap(cap uint64) (hm *HashMap) {
	hm = new(HashMap)
	if cap <= DefaultCap {
		hm.cap = DefaultCap
	} else {
		hm.cap = 1 << int(math.Ceil(math.Log2(float64(cap))))
	}
	hm.len = 0
	hm.expendFactor = uint64(math.Floor(float64(hm.cap) * 0.75))
	hm.repo = make([]*item, hm.cap, hm.cap)

	return
}

//extend capacity
func (hm *HashMap) extend() (error error) {
	defer runtime.GC()
	// lock by channel
	hm.waitCh, hm.wait = make(chan int, 1), true

	defer func() {
		hm.waitCh <- 1
		hm.wait = false
	}()
	var growCap uint64 = 1 << int(math.Ceil(math.Log2(float64(hm.cap)))+1)
	nhm := NewHashMap(growCap)
	/**
	 * Q&A:
	 * Q: 如果在扩展中发生了Store或者Del操作,造成数据丢失等,应该如何处理这部分问题?
	 * A: 对原型hashMap加锁,但是不能影响到Store操作
	 * TODO:
	 */
Crash:
	for _, item := range hm.repo {
		if item == nil {
			continue
		}
		for {
			error = nhm.Store(item.key, item.value)
			if error != nil {
				break Crash
			}
			if item.next == nil {
				break
			}
			item = item.next
		}
		// time.Sleep(100 * time.Microsecond) // 模拟大数据扩容
	}
	// notice: 这里只进行对等属性的copy, 不要直接*hm=*nhm
	// *hm=*nhm会导致waitCh的丢失
	hm.repo = nhm.repo
	hm.len = nhm.len
	hm.cap = nhm.cap
	hm.expendFactor = nhm.expendFactor

	return error
}

// Store
func (hm *HashMap) Store(k string, v interface{}) (error error) {
	if hm.wait {
		log.Println("call:Store, waiting for expending..")
		<-hm.waitCh
	}
	if hm.len+1 >= hm.expendFactor {
		log.Println(fmt.Sprintf("(len=%d cap=%d), start expend..", hm.len, hm.cap))
		if error = hm.extend(); error != nil {
			return error
		}
		log.Println(fmt.Sprintf("expend done, current capacity=%d", hm.cap))
		return hm.Store(k, v)
	} else {
		hm.lock.Lock()
		defer hm.lock.Unlock()

		hash := hashKey(k)
		index := hm.index(hash)
		item := &item{key: k, value: v}
		if nil != hm.repo[index] {
			// 该索引存在item了
			ptr := hm.repo[index]
			for {
				if ptr.key == k {
					// 存在相同key,则覆盖更新value
					ptr.value = v
					goto END
				}
				if ptr.next != nil {
					ptr = ptr.next
				} else {
					break
				}
			}
			ptr.next = item
		} else {
			hm.repo[index] = item
		}
		hm.len++
	END:
		return nil
	}
}

// Get: get value by key
func (hm *HashMap) Get(k string) (v interface{}, hit bool) {
	if hm.wait {
		log.Println("call:Get, waiting for expending..")
		<-hm.waitCh
	}

	hit = false
	hash := hashKey(k)
	index := hm.index(hash)
	if hm.repo[index] == nil {
		v = nil
		return
	}
	item := hm.repo[index]
	for {
		if item.key == k {
			v, hit = item.value, true
			break
		}
		if item.next == nil {
			break
		}
		item = item.next
	}
	return
}

func (hm *HashMap) Del(k string) (v interface{}, ok bool) {
	if hm.wait {
		log.Println("call:Del, waiting for expending..")
		<-hm.waitCh
	}

	hash := hashKey(k)
	index := hm.index(hash)
	if hm.repo[index] == nil {
		v = nil
		return
	}

	var (
		ptr    = &(hm.repo[index]) // 取指针(*item)的指针: **item
		item   = hm.repo[index]    // *item
		isRoot = true
	)
	for {
		if item == nil {
			break
		}
		if item.key == k {
			// hit
			v, ok = item.value, true
			// remove
			if item.next == nil && isRoot {
				*ptr = nil
			}
			if item.next == nil && !isRoot {
				(*ptr).next = nil
			}
			if item.next != nil && isRoot {
				*ptr = item.next
			}
			if item.next != nil && !isRoot {
				(*ptr).next = item.next
			}
			break
		} else {
			// to next
			if isRoot {
				isRoot = false
			} else {
				ptr = &item
			}
			item = item.next
		}
	}
	return
}

// 打印结构图
func (hm *HashMap) print() {
	fmt.Println(strings.Repeat("=", 20))
	fmt.Println("打印HashMap结构图")
	fmt.Println(strings.Repeat("=", 20))

	for _, link := range hm.repo {
		var buffer bytes.Buffer
		if link == nil {
			continue
		}
		for {
			_, _ = buffer.WriteString(fmt.Sprintf("N[%s=%+v]", link.key, link.value))
			if link.next != nil {
				link = link.next
				_, _ = buffer.WriteString("-->")
			} else {
				break
			}
		}
		if buffer.Len() > 0 {
			fmt.Println(buffer.String())
		}
	}

	fmt.Println(strings.Repeat("=", 20))

}

func hashKey(k string) (hash uint64) {
	var h = xxhash.New64()
	if _, err := h.Write([]byte(k)); err != nil {
		log.Fatal(err)
	}
	hash = h.Sum64()
	return
}

func (hm *HashMap) index(hash uint64) (index uint64) {
	index = hash & (hm.cap - 1)
	return
}

func (hm *HashMap) debug() {

	fmt.Println(hm.repo)
}

func main() {

	var dataCollection = [...][2]string{
		{"key1", "this is a string"},
		{"key2", "为什么你这么熟练啊"},
		{"key3", "你不要过来啊"},
		{"key4", "jojo,我不做人啦!"},
		{"key5", "炸哇陆多!"},
		{"key6", "kksk"},
		{"key7", "ko~ ko~ da~ yo~"},
		{"key8", "404"},
		{"key9", "打死白学家"},
		{"key10", "1111"},
		{"key11", "22222"},
		{"key12", "3333"},
		{"key13", "吃我压路机~~~"},
		{"key14", "+++++"},
		{"key2", "为什么你这么熟练啊??"},   // duplicate key
		{"key4", "jojo,我不做人啦~~!"}, // duplicate key
	}
	var dataCollection2 = [...][2]string{
		{"key1111", "this is a string"},
		{"key2222", "为什么你这么熟练啊"},
		{"key33333", "你不要过来啊"},
		{"key44444", "jojo,我不做人啦!"},
		{"key5333", "炸哇陆多!"},
		{"key644", "kksk"},
		{"key733", "ko~ ko~ da~ yo~"},
		{"key81", "404"},
		{"key9232", "打死白学家"},
		{"key102", "1111"},
		{"key114", "22222"},
		{"key112", "3333"},
		{"key12223", "吃我压路机~~~"},
		{"key1334", "+++++"},
		{"key2444", "为什么你这么熟练啊??"},  // duplicate key
		{"key455", "jojo,我不做人啦~~!"}, // duplicate key
	}

	var hm = NewHashMap(0)

	var (
		wg  sync.WaitGroup
		ch  = make(chan string, 10)
		ch2 = make(chan string, 1)
	)
	// ctx, cancel := context.WithCancel(context.Background())
	d := time.Now().Add(10000 * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), d)
	defer cancel()

	wg.Add(4)

	// 测试方法:
	// 开启4个goroutine, 2个并发写入, 1个读取, 1个删除
	// 关闭读取和删除协程使用上下文timeout控制

	go func(hm *HashMap, wg *sync.WaitGroup) {
		for _, data := range dataCollection {
			log.Printf("adding k=%s \n", data[0])
			_ = hm.Store(data[0], data[1])
			ch <- data[0]
		}
		wg.Done()

	}(hm, &wg)
	go func(hm *HashMap, wg *sync.WaitGroup) {
		for _, data := range dataCollection2 {
			log.Printf("adding k=%s \n", data[0])
			_ = hm.Store(data[0], data[1])
			ch <- data[0]
		}

		wg.Done()
	}(hm, &wg)

	go func(hm *HashMap, wg *sync.WaitGroup, ctx context.Context) {
	END:
		for {
			select {
			case key := <-ch:
				if key == "key4" {
					value, _ := hm.Get(key)
					fmt.Println(fmt.Sprintf("GET HIT: k=%s, v=%s", key, value))
				}
				if key == "key733" {
					value, _ := hm.Get(key)
					fmt.Println(fmt.Sprintf("GET HIT: k=%s, v=%s", key, value))
				}
				if key == "key8" {
					value, _ := hm.Get(key)
					fmt.Println(fmt.Sprintf("GET HIT: k=%s, v=%s", key, value))
				}
				ch2 <- key // 读完了就写
			case <-ctx.Done():
				fmt.Println("timeout, close Get loop.")
				break END
			default:
			}
		}
		wg.Done()
	}(hm, &wg, ctx)

	go func(hm *HashMap, wg *sync.WaitGroup, ctx context.Context) {
		END:
		for {
			select {
			case rk := <-ch2:
				for _, key := range []string{"key102", "key9232", "key9"} {
					if key == rk {
						if deletedValue, ok := hm.Del(key); ok {
							fmt.Printf("deleted: k=%s, v=%s\n", key, deletedValue)
						} else {
							fmt.Printf("delete failed, not hit key (%s)\n", key)
						}
					}
				}

			case <-ctx.Done():
				fmt.Println("timeout, close Del loop.")
				break END
			}
		}
		wg.Done()
	}(hm, &wg, ctx)

	wg.Wait()

	hm.print()

}
