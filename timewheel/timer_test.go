package timewheel

import (
	"fmt"
	"log"
	"testing"
	"time"
)

func TestTimer(t *testing.T) {
	tw := Timer()
	var (
		timeA = time.Now()
		cntA  = 0
		//timeB = time.Now()
		//cntB  = 0
		//timeC = time.Now()
		//cntC  = 0
		//timeD = time.Now()
		//cntD  = 0
	)

	tw.addTask(3*time.Second, true, func(args JobArgs) {
		fmt.Println("!!! called \033[31mAAAA\033[0m task >>>>>>>>>", time.Now().Sub(timeA))
		timeA = time.Now()
		cntA++
	}, nil)
	//tw.addTask(4*time.Second, false, func(args JobArgs) {
	//	fmt.Println("!!! called \u001B[31mBBBB\u001B[0m task >>>>>>>>>", time.Now().Sub(timeB))
	//	timeB = time.Now()
	//	cntB++
	//}, nil)
	//tw.addTask(10*time.Second, false, func(args JobArgs) {
	//	fmt.Println("!!! called \u001B[31mCCCC\u001B[0m task >>>>>>>>>", time.Now().Sub(timeC))
	//	timeC = time.Now()
	//	cntC++
	//}, nil)
	//tw.addTask(30*time.Second, false, func(args JobArgs) {
	//	fmt.Println("!!! called \u001B[31mDDDD\u001B[0m task >>>>>>>>>", time.Now().Sub(timeD))
	//	timeD = time.Now()
	//	cntD++
	//}, nil)

	time.Sleep(121 * time.Second)
	log.Println(
		fmt.Sprintf("cntA = %d", cntA),
		//fmt.Sprintf("cntB = %d", cntB),
		//fmt.Sprintf("cntC = %d", cntC),
		//fmt.Sprintf("cntD = %d", cntD),
	)
}

func debugTimeWheels() {
	for idx, ln := range gTimer.wheels {
		for ln != nil {
			log.Println("debug wheels --------------->", idx, "|", ln.getID())
			if ln.getNext() != nil {
				ln = ln.getNext().(*bucket)
			} else {
				break
			}
		}
	}
}
