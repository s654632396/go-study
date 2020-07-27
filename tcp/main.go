package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
	"strconv"
	"sync"
	"sync/atomic"
	"unsafe"
)

type MyClient struct {
	id   int64
	conn net.Conn
	oCh  <-chan []byte
}

func NewMyClient(c net.Conn, id int64) *MyClient {
	return &MyClient{
		id:   id,
		conn: c,
	}
}

type MyChannel struct {
	iCh  chan []byte
	oChs map[int64]chan []byte
}

func (mc *MyChannel) addListener(c *MyClient) {
	var lck sync.Mutex
	lck.Lock()
	defer lck.Unlock()
	cch := make(chan []byte)
	// mc.oChs = append(mc.oChs, cch)
	mc.oChs[c.id] = cch
	c.oCh = cch
	mc.broadcast(c.id, []byte("\033[34m["+strconv.FormatInt(c.id, 10)+"]\033[0m 进入了聊天室\n"))
}

func (mc *MyChannel) rmListener(c *MyClient) {
	delete(mc.oChs, c.id)
}

func (mc *MyChannel) start() {
	for {
		select {
		case pdata := <-mc.iCh:
			id, data := unpack(pdata)
			if len(data) >= 2 && string(data[:2]) == "\\C" {
				mc.useCommand(id, string(data[3:7]))
			} else {
				mydata := unionBytes(
					[]byte("\033[34m["+strconv.FormatInt(id, 10)+"]\033[0m "),
					data,
				)
				mc.broadcast(id, []byte(mydata))
			}
		}
	}
}

func (mc *MyChannel) broadcast(id int64, data []byte) {
	for oid, outCh := range mc.oChs {
		if oid != id {
			outCh <- data
		}
	}
}

func (mc *MyChannel) useCommand(id int64, cmd string) {
	if cmd == "QUIT" {
		// 离开
		log.Println(fmt.Sprintf("clientID[%d] leave chat channel.", id))
		delete(mc.oChs, id)
		mc.broadcast(id, []byte("\033[34m["+strconv.FormatInt(id, 10)+"]\033[0m 离开了聊天室\n"))
	}
	if cmd == "LIST" {
		log.Println("call command: LIST")
		for sid, _ := range mc.oChs {
			mc.sendTo(id, []byte(strconv.FormatInt(sid, 10)+"\n"))
		}
	}
}

func (mc *MyChannel) sendTo(id int64, data []byte) {
	log.Println("sent", data)
	mc.oChs[id] <- data
}

// 简易群聊 based on TCP
//
func main() {
	lsn, err := net.Listen("tcp", "127.0.0.1:4300")
	if err != nil {
		panic(err)
	}
	mc := &MyChannel{
		iCh:  make(chan []byte),
		oChs: make(map[int64]chan []byte),
	}
	go mc.start()

	var autoID int64 = 1
	for {
		conn, err := lsn.Accept()
		if err != nil {
			log.Println(err)
		}
		c := NewMyClient(conn, autoID)
		atomic.AddInt64(&autoID, 1)
		log.Println(fmt.Sprintf("connected[%d]: %s", c.id, conn.RemoteAddr()))

		mc.addListener(c)
		go handleConnRead(c, mc.iCh)
		go handleConnWrite(c)
	}
}

func handleConnRead(mc *MyClient, out chan<- []byte) {
	defer func() {
		log.Println("closed:", mc.conn.RemoteAddr())
		out <- pack(mc.id, []byte("\\C QUIT"))
		_ = mc.conn.Close()
	}()
	reader := bufio.NewReader(mc.conn)
	for {
		mc.conn.Write([]byte{'>', ' '})
		cmd, err := reader.ReadBytes('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			log.Println(err)
			break
		}
		if len(bytes.Trim(cmd, " \r\n\t")) == 0 {
			continue
		}
		// log.Println("get req command:", cmd)
		out <- pack(mc.id, cmd)
	}
}

func handleConnWrite(mc *MyClient) {
	for {
		select {
		case data := <-mc.oCh:
			mc.conn.Write(data)
			mc.conn.Write([]byte{'>', ' '})
		}
	}
}

// pack: 打包数据
// cast(int64, [8]byte) && unshift到data里
func pack(id int64, data []byte) []byte {
	ptr := uintptr(unsafe.Pointer(&id))
	var bs [8]byte
	//noinspection GoVetUnsafePointer
	bs = *(*[8]byte)(unsafe.Pointer(ptr))
	var pdata bytes.Buffer
	pdata.Write(bs[:8])
	pdata.Write(data)
	return pdata.Bytes()
}

// unpack: 解包数据
func unpack(pdata []byte) (id int64, data []byte) {
	if len(pdata) == 0 {
		log.Println("err package data..")
		return 0, nil
	}
	// var fb = [8]byte{pdata[0], pdata[1], pdata[2], pdata[3], pdata[4], pdata[5], pdata[6], pdata[7]}
	var fb [8]byte
	for i := 0; i < 8; i++ {
		fb[i] = pdata[i]
	}
	id = *(*int64)(unsafe.Pointer(&fb))
	data = pdata[8:]
	return
}

func unionBytes(data ...[]byte) []byte {
	var l int
	for _, dat := range data {
		l += len(dat)
	}
	var ret []byte
	ret = make([]byte, l)
	var offset = 0
	for _, dat := range data {
		copy(ret[offset:offset+len(dat)], dat)
		offset += len(dat)
	}
	return ret
}

// 测试data pack转换
func _() {
	var id int64 = 45648112154654515
	log.Println(id)
	pdata := pack(id, []byte{'a', 'b', 'c'})
	log.Println(unpack(pdata))
	return
}
