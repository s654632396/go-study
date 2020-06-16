package main

import (
	"encoding/gob"
	"fmt"
	"os"
)

type PA struct {
	X, Y, Z int
	Name    string
}

func main() {

	file, err := os.Create("gob.dat")
	defer file.Close()
	if err != nil {
		panic(err)
	}

	//p := make([]PA, 0)
	//p = append(p, PA{5, 6, 7, "Hello young man"})
	//p = append(p, PA{3, 4, 5, "Pythagoras"})
	var p = [...]PA{
		PA{5, 6, 7, "Hello young man"},
		PA{3, 4, 5, "Pythagoras"},
	}
	enc := gob.NewEncoder(file)

	for _, pp := range p {
		fmt.Println("write to gob.dat:", pp)
		if err := enc.Encode(pp); err != nil {
			panic(fmt.Errorf("Encode ERR: %s\n", err))
		}
	}

}
