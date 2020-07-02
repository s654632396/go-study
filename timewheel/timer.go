package timewheel

import (
	"github.com/satori/go.uuid"
	"sync"
	"time"
)

type BucketType int

const (
	BucketTypeAfter  BucketType = 1
	BucketTypeTicker BucketType = 2
)

// bucket node (a linked list)
type bucket struct {
	LinkNode
	j         Job
	jArgs     JobArgs
	t         BucketType    // 类型
	startTime time.Time     // 开始时间
	interval  time.Duration // 间隔
	round     int           // 需要转动多少轮
}

type timer struct {
	ticker       *time.Ticker  // 定时器
	wheels       []*bucket     // 时间轮
	size         int           // 时间轮大小
	tickDuration time.Duration // 轮盘刻度间隔时间
	tick         int           // 轮盘指针
}

var once sync.Once
var gTimer *timer

func Timer() *timer {
	once.Do(func() {
		gTimer = new(timer)
		gTimer.size = 1024
		gTimer.tickDuration = time.Millisecond // recommend larger than millisecond
		gTimer.wheels = make([]*bucket, gTimer.size)
		gTimer.tick = 0
		gTimer.start()
	})
	return gTimer
}

func (t *timer) start() {
	if t == nil {
		t = Timer()
	}
	t.ticker = time.NewTicker(t.tickDuration)
	//var debugTicker = time.NewTicker(time.Second)
	go func() {
		var mCtr = 0
		for {
			select {
			case <-t.ticker.C:
				mCtr++
				t.runTask(t.tick)
				if t.tick+1 >= t.size {
					t.tick = 0
				} else {
					t.tick++
				}
			//case <-debugTicker.C:
				// log.Println("mCtr=", mCtr, "tick=", t.tick)
				// mCtr = 0
			}
		}
	}()
}

type JobArgs []interface{}
type Job func(JobArgs)

func (t *timer) addTask(interval time.Duration, noDelay bool, job Job, jArgs JobArgs) {
	if interval <= 0 {
		panic(`invalid args to create task`)
	}
	b := new(bucket)
	// create bucket ID
	b.ID = uuid.NewV4().String()
	b.next = nil
	b.t = BucketTypeTicker
	// start_time = current_time + delay_time
	b.startTime = time.Now().Add(interval)
	b.interval = interval
	b.j = job
	b.jArgs = jArgs
	if noDelay {
		go b.j(b.jArgs)
	}
	hashIndex := t.getBucketTickIndex(b)
	// log.Printf("insert task with hi=%d\n", hashIndex)
	t.pushBucket(b, hashIndex)
}

func (t *timer) pushBucket(bkt *bucket, idx int) {

	var currBucket = t.wheels[idx]
	if currBucket == nil {
		t.wheels[idx] = bkt
		return
	}
	currBucket.append(bkt)
}

func (t *timer) getBucketTickIndex(b *bucket) int {
	var hashIndex int
	b.round, hashIndex = t.calculateRoundAndTick(b.interval)
	if t.tick+hashIndex >= t.size {
		hashIndex = t.tick + hashIndex - t.size
	} else {
		hashIndex = t.tick + hashIndex
	}
	return hashIndex
}

func (t *timer) calculateRoundAndTick(in time.Duration) (r, h int) {
	r = int(in/t.tickDuration) / t.size
	h = int(in/t.tickDuration) % t.size
	//log.Printf("curTick=%d, in=%d, td=%d, tl=%d, r=%d, h=%d\n", t.tick, in, t.tickDuration, t.size, r, h)
	return
}

func (t *timer) runTask(idx int) {
	var bkt = t.wheels[idx]
	for bkt != nil {
		if bkt.round > 0 {
			bkt.round--
			if bkt.getNext() != nil {
				bkt = bkt.getNext().(*bucket)
				continue
			} else {
				break
			}
		} else {
			// async running jobs
			bkt.j(bkt.jArgs)
			// remove this bucket from linklist
			t.removeTaskFromIndex(idx, bkt.ID)
			// re-insert bkt
			hashIndex := t.getBucketTickIndex(bkt)
			nextBkt := bkt.getNext()
			bkt.next = nil // 设置该bucket的next为nil,防止链表循环
			t.pushBucket(bkt, hashIndex)
			if nextBkt == nil {
				break
			}
			bkt = nextBkt.(*bucket)
		}
	}
}

func (t *timer) removeTaskFromIndex(index int, ID string) {
	if t.wheels[index] == nil {
		return
	}

	var lock sync.Mutex
	lock.Lock()
	defer lock.Unlock()

	rootID := t.wheels[index].getID()
	pNode := t.wheels[index]
	currNode := t.wheels[index]
	for {
		if currNode.getID() == ID {
			// matched
			if currNode.getID() == rootID {
				// is root bucket
				if currNode.getNext() == nil {
					t.wheels[index] = nil
				} else {
					t.wheels[index] = currNode.getNext().(*bucket)
				}
			} else {
				if currNode.getNext() == nil {
					pNode.setNext(nil)
				} else {
					pNode.setNext(currNode.getNext())
				}
			}
		}
		if currNode.getNext() == nil {
			break
		}
		currNode = currNode.getNext().(*bucket)
	}
}
