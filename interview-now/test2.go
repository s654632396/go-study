package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	ch1 := make(chan struct{})
	ch2 := make(chan struct{})
	ch3 := make(chan struct{})



	rand.Seed(time.Now().UnixNano())
	for i := 1; i <= 3; i++ {
		go func(cid int) {
			sec := time.Duration(rand.Int31n(10)) * time.Second
			fmt.Println("cid=", cid, " sleep ", sec, " s")
			time.Sleep(sec)
			switch cid {
			case 1:
				ch1 <- struct{}{}
			case 2:
				ch2 <- struct{}{}
			case 3:
				ch3 <- struct{}{}
			}
		}(i)
	}


	select {
	case <-ch1:
		fmt.Println("ch1 done first.")
	case <-ch2:
		fmt.Println("ch2 done first.")
	case <-ch3:
		fmt.Println("ch3 done first.")
	}

}
