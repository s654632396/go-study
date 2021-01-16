package main

import (
	"bytes"
	"fmt"
)

type Empty struct{}
type Set map[interface{}]Empty

func (this Set) Add(vs ...interface{}) Set {
	for _, v := range vs {
		this[v] = Empty{}
	}
	return this
}

func (this Set) String() {
	var buf bytes.Buffer
	for _, v := range this {
		if buf.Len() > 0 {
			buf.WriteString(",")
		}
		buf.WriteString(fmt.Sprintf("%v", v))
	}
}

func NewSet() Set {
	return make(map[interface{}]Empty)
}

func main() {
	set := NewSet().Add(1,2,3,4,5,6,7,5,6,10)
	fmt.Println(set)
}
