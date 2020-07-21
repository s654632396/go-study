package main

import (
	"fmt"
	"log"
	"os"
	"syscall"
)

func main() {

	var dbFilePath = "test.db"
	// 打开文件句柄
	f, err := os.OpenFile(dbFilePath, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// 获取文件信息
	fInfo, err := os.Stat(dbFilePath)
	if err != nil {
		panic(err)
	}

	var size = int(fInfo.Size())
	// Mmap的长度必须大于0
	if size == 0 {
		size = 32
		// expand file
		syscall.Truncate(dbFilePath, int64(size))
	}
	log.Println(int(fInfo.Size()))
	var data []byte
	data, err = syscall.Mmap(int(f.Fd()), 0, size, syscall.PROT_READ|syscall.PROT_WRITE|syscall.PROT_EXEC, syscall.MAP_SHARED)
	if err != nil {
		panic(err)
	}

	addr := &data[0]
	fmt.Println("mmap success,addr=", addr, "size=", len(data))
	var myData []byte = ([]byte)("奇迹2:高尔夫")

	// 写入到映射好的内存中
	for i, b := range myData {
		if i > len(data) {
			break
		}
		data[i] = b
	}

	//取消映射
	syscall.Munmap(data)
}
