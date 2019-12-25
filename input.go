package main

import (
	"fmt"
	"github.com/gen2brain/dlgs"
)

func main() {
	item, _, err := dlgs.List("List", "Select item from list:", []string{"Bug", "New Feature", "Improvement"})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", item)
}
