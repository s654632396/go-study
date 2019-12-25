package main

// 使用写入缓冲区的方式来进行文件的写入
// bufio的写入器
import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var file string = "output.dat"

	// os.O_RDONLY 只读
	// os.O_WRONLY 只写
	// os.O_CREATE 创建, 当文件不存在时候生效
	// os.O_TRUNC 截断, 当文件存在时清空内容
	// os.O_APPEND 追加到文件尾部
	// os.O_EXCL 与O_CREATE合用,用于判断文件存在否,存在则返回错误 (用于某些需要原子性文件处理的场景)
	// ----
	// os.O_SYNC 将文件系统缓冲区数据同步到磁盘的标识 = write + fsync
	outputFile, err := os.OpenFile(file, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0666)
	if err != nil {
		fmt.Printf("An error occurred with file opening or creation:\n")
		fmt.Println(fmt.Errorf("err_msg: %s", err))
		return
	}
	defer outputFile.Close()

	wb := bufio.NewWriter(outputFile)

	for i := 0; i < 10; i++ {
		wb.WriteString("how to write file in golang with using bufio\n")
	}

	wb.Flush()

}
