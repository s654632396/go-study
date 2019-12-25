package main

import (
	"io"
	"os"
)

func main() {
	CopyFile("source.txt", "target.txt")
}

func CopyFile(srcName, dstName string) (wirtten int64, err error) {

	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()

	dst, err := os.OpenFile(dstName, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return
	}
	defer dst.Close()

	return io.Copy(dst, src)
}
