// gob1.go
package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
)

type P struct {
	X, Y, Z int
	Name    string
}

type Q struct {
	X, Y *int32
	Name string
}

func main() {
	// Initialize the encoder and decoder.  Normally enc and dec would be
	// bound to network connections and the encoder and decoder would
	// run in different processes.

	var msg bytes.Buffer // Stand-in for a network connection

	// Encode (send) the value.
	var ps []P = []P{
		{3, 4, 5, "Pythagoras"},
		{5, 6, 7, "Hello young man"},
	}

	for _, p1 := range ps {
		enc := gob.NewEncoder(&msg) // Will write to network.
		if err := enc.Encode(p1); err != nil {
			log.Fatal("encode error:", err)
		}

		dec := gob.NewDecoder(&msg)
		var q Q
		if err := dec.Decode(&q); err != nil {
			log.Fatal("decode error:", err)
		}
		msg.Truncate(msg.Len())

		fmt.Printf("%q: {%d,%d}\n", q.Name, *q.X, *q.Y)
	}

}

// Output:   "Pythagoras": {3,4}
