package hashmap

import (
	"context"
	"fmt"
	"log"
	"sync"
	"testing"
	"time"
)

// 容量测试
func BenchmarkHashMap_Store(b *testing.B) {
	hm := NewHashMap(0)
	for i := 1; i <= 1000000; i++ {
		_ = hm.Store("key"+string(i), "value"+string(i))
	}
	log.Println(fmt.Sprintf("Hash Map expending capacity to %d", hm.cap))
	log.Println("[BenchmarkHashMap_Store] Over.")
}

// 并发测试
func BenchmarkNewHashMap_Concurrency(b *testing.B) {
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
		wg     sync.WaitGroup
		ch     = make(chan string, 10)
		ch2    = make(chan string, 10)
		doneCh = make(chan int, 2)
	)

	// ctx, cancel := context.WithCancel(context.Background())
	d := time.Now().Add(5000 * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), d)
	defer cancel()

	wg.Add(3)

	// 测试方法:
	// 开启4个goroutine, 2个并发写入, 1个读取, 1个删除
	// 关闭读取和删除协程使用上下文timeout控制

	go func(hm *HashMap, wg *sync.WaitGroup) {
		for _, data := range dataCollection {
			log.Printf("[proc1]adding k=%s \n", data[0])
			_ = hm.Store(data[0], data[1])
			ch <- data[0]
		}
		doneCh <- 1
		wg.Done()
	}(hm, &wg)
	go func(hm *HashMap, wg *sync.WaitGroup) {
		for _, data := range dataCollection2 {
			log.Printf("[proc2]adding k=%s \n", data[0])
			_ = hm.Store(data[0], data[1])
			ch <- data[0]
		}
		doneCh <- 1
		wg.Done()
	}(hm, &wg)

	go func(hm *HashMap, wg *sync.WaitGroup, ctx context.Context) {
		var done int = 0
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
			case <-doneCh:
				done++
				// 所有Store操作的协程都完成了
				if done >= 2 {
					break END
				}
			case <-ctx.Done():
				// 超时关闭
				fmt.Println("timeout, close Get loop.")
				time.Sleep(500 * time.Millisecond)
				break END
			}
		}
		wg.Done()
	}(hm, &wg, ctx)

	wg.Wait()
	// hm.print()
}
