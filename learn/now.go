package main

import (
	"fmt"
	"time"
)

func main() {
	var now time.Time = time.Now()

	fmt.Printf("current local time: %v \n", now)
	fmt.Printf("UTC time: %v \n", now.UTC())
	fmt.Printf("Unix timestamp: %v \n", now.Unix())

	fmt.Printf("current year:%v \n", now.Year())
	fmt.Printf("current month:%v | %d | %#v \n", now.Month(), now.Month(), now.Month())
	fmt.Printf("current day:%v \n", now.Day())
	fmt.Printf("current Hour:%v \n", now.Hour())
	fmt.Printf("current minute:%v \n", now.Minute())
	fmt.Printf("current second:%v \n", now.Second())
	fmt.Printf("current Nano second:%v \n", now.Nanosecond())
}
