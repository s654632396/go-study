package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	// var fname, lname string
	// fmt.Println("Please enter your name:")
	// fmt.Scanln(&fname, &lname)
	// var name string = fname + lname
	// fmt.Printf("Hi~ %s, welcome!", name)

	inputReader := bufio.NewReader(os.Stdin)

	fmt.Println("Please type something:")

	input, isPrefix, err := inputReader.ReadLine()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("isPrefix=%v\n", isPrefix)
	fmt.Printf("your input content is :\n%s\n", input)

	var input2 string
	// ReadString() 读取到参数到delim字符，则返回
	input2, err = inputReader.ReadString('\n')
	fmt.Printf("your input content is :\n%s\n", input2)
}
