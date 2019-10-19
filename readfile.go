package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("./input.sock")

	if err != nil {
		fmt.Printf("An error occurred on opening the inputfile\n" +
			"Does the file exist?\n" +
			"Have you got acces to it?\n")
		return // exit the function on error
	}
	defer file.Close()
	inputReader := bufio.NewReader(file)

	for {
		inputString, readerError := inputReader.ReadString('\n')
		if readerError != nil {
			// fmt.Println("err:", readerError)
			// goto HERE
		} else {
			fmt.Println(inputString)
			if inputString == "close\n" {
				fmt.Println("end!")
				break
			}
		}
	}
	// HERE:
}
