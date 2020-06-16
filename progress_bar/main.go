package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	// Clear the screen by printing \x0c.

	runBar()
}

func runBar() {
	const col = 30

	bar := fmt.Sprintf("[%%-%vs]\r", col)
	for i := 0; i < col; i++ {
		fmt.Printf(bar, strings.Repeat("=", i)+">")
		time.Sleep(100 * time.Millisecond)
	}
	fmt.Printf(bar+" Done!", strings.Repeat("=", col))
}
