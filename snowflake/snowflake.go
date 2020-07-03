package snowflake

import (
	"sync"
	"time"
)

const (
	TimeStampShiftNum int64 = 22
	MachineShiftNum   int64 = 13
	MaxSerialID       int64 = 1 << MachineShiftNum
)

var offsetTS int64

type SnowID struct {
	machineID int64
	mts       int64
	serialID  int64
	lastTs    int64
	lck       sync.Mutex
}

func NewSnowID(ts int64, machineID int64) *SnowID {
	offsetTS = ts
	var sid *SnowID
	sid = new(SnowID)
	sid.machineID, sid.mts = machineID, getCurrentTS()
	return sid
}

func (sid *SnowID) GID() int64 {
	sid.lck.Lock()
	defer sid.lck.Unlock()

	sid.mts = getCurrentTS()
	sid.getSerialNum()
	if sid.serialID >= MaxSerialID {

	}
	return (sid.mts << TimeStampShiftNum) | (sid.machineID << MachineShiftNum) | sid.serialID
}

func getCurrentTS() int64 {
	return (time.Now().UnixNano() / 1e6) - offsetTS
}

func (sid *SnowID) getSerialNum() int64 {
	if sid.mts == sid.lastTs {
		sid.serialID ++
	} else {
		sid.serialID = 0
	}
	sid.lastTs = sid.mts
	return sid.serialID
}

func ParseSnowID(sid int64) (microTS time.Time, serial int64) {
	const tsMask = 0x1ffffffffff << TimeStampShiftNum
	var microTimeStamp = int64((tsMask&sid)>>TimeStampShiftNum) + offsetTS
	microTS, serial = time.Unix(microTimeStamp/1e3, 1e6*(microTimeStamp%1e3)), sid&(MaxSerialID-1)

	return
}
