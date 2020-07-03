package snowflake

import (
	"sync"
	"time"
)

const (
	MaxSerialID = 4096 << 1
)

var SerialCh chan int
var OffsetTS int64

type SnowID struct {
	machineID  uint8
	mTimeStamp int64
}

func GID(machineID uint8) uint64 {
	sid := SnowID{machineID, -1}
	var realID uint64 = 0
	realID ^= uint64(sid.getMicroTS()) << 22
	realID ^= uint64(sid.machineID) << 13
	serialID := sid.getSerialNum()
	realID ^= uint64(serialID)
	return realID
}

func (sid *SnowID) getMicroTS() int64 {
	if sid.mTimeStamp == -1 {
		sid.mTimeStamp = getCurrentTS()
	}
	return sid.mTimeStamp
}

func getCurrentTS() int64 {
	return (time.Now().UnixNano() / 1000000) - OffsetTS
}

func InitSnowSerial(ts int64) {

	var single sync.Once
	single.Do(func() {
		// init vars and serial-id-gen-goroutine
		OffsetTS = ts
		SerialCh = make(chan int, 0)
		go func() {
			var lastTS int64
			for i := 0; i < MaxSerialID; i++ {
				if getCurrentTS() != lastTS {
					i = 0
					lastTS = getCurrentTS() // update
					SerialCh <- i
				} else {
					SerialCh <- i // ret
				}
			}
			panic(`serial ID max value arrived, generate failed!`)
		}()
	})

}

func (sid *SnowID) getSerialNum() int {
	if SerialCh == nil {

	}

	return <-SerialCh
}
