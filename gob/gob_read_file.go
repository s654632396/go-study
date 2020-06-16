package main

import (
	"encoding/gob"
	"fmt"
	"io"
	"log"
	"os"
)

type QA struct {
	X, Y, Z int
	Name    string
}

func main() {

	file, err := os.Open("gob.dat")
	defer file.Close()
	if err != nil {
		panic(err)
	}
	dec := gob.NewDecoder(file)
	var q []QA = make([]QA, 5)
	var qPtr *QA
	for i := 0; ; i++ {
		qPtr = &q[i]
		if err := dec.Decode(qPtr); err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal("decode error:", err)
		}
		fmt.Printf("%q: {%d,%d,%d}\n", q[i].Name, q[i].X, q[i].Y, q[i].Z)
	}
}
