package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
	"sync"
	"sync/atomic"
	"unsafe"
)

type MyConn struct {
	id   int64
	conn net.Conn
	oCh  <-chan []byte
}

func NewMyConn(c net.Conn, id int64) *MyConn {
	return &MyConn{
		id:   id,
		conn: c,
	}
}

type MyChannel struct {
	iCh  chan []byte
	oChs []chan []byte
}

func (mc *MyChannel) addListener(c *MyConn) {
	var lck sync.Mutex
	lck.Lock()
	defer lck.Unlock()
	cch := make(chan []byte)
	mc.oChs = append(mc.oChs, cch)
	c.oCh = cch
}

func (mc *MyChannel) start() {
	for {
		select {
		case pdata := <-mc.iCh:
			log.Println("channel read data:", string(pdata))
			mc.broadcast(pdata)
		}
	}
}

func (mc *MyChannel) broadcast(data []byte) {
	for _, outCh := range mc.oChs {
		outCh <- data
	}
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
		oChs: make([]chan []byte, 0),
	}
	go mc.start()

	var autoID int64 = 1
	for {
		conn, err := lsn.Accept()
		if err != nil {
			log.Println(err)
		}
		c := NewMyConn(conn, autoID)
		atomic.AddInt64(&autoID, 1)
		log.Println(fmt.Sprintf("connected[%d]: %s", c.id, conn.RemoteAddr()))

		mc.addListener(c)
		go handleConnRead(c, mc.iCh)
		go handleConnWrite(c)
	}

}

func handleConnRead(mc *MyConn, out chan<- []byte) {
	defer func() {
		log.Println("closed:", mc.conn.RemoteAddr())
		_ = mc.conn.Close()
	}()
	reader := bufio.NewReader(mc.conn)
	for {
		cmd, err := reader.ReadBytes('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			log.Println(err)
			break
		}
		log.Println("get req command:", cmd)
		out <- pack(mc.id, cmd)
	}

}

func handleConnWrite(mc *MyConn) {
	for {
		select {
		case pdata := <-mc.oCh:
			id, data := unpack(pdata)
			if id == mc.id {
				continue
			}
			log.Println("send data to:", mc.id, string(data), "pdata.id=", id)
			mc.conn.Write(data)
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
	var fb = [8]byte{pdata[0], pdata[1], pdata[2], pdata[3], pdata[4], pdata[5], pdata[6], pdata[7]}
	id = *(*int64)(unsafe.Pointer(&fb))
	data = pdata[8:]
	return
}

// 测试data pack转换
func _() {
	var id int64 = 45648112154654515
	log.Println(id)
	pdata := pack(id, []byte{'a', 'b', 'c'})
	log.Println(unpack(pdata))
	return
}
