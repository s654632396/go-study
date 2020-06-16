package main

import (
	"bytes"
	"fmt"
	"github.com/golang/protobuf/proto"
	"log"
	"net/http"
	"study-go/grpc/world"
)

func main() {
	request()

	return
}

func request() {

	resp, err := http.Get("http://127.0.0.1:10011")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	//fmt.Println(resp, err)

	var hello = world.Hello{}
	var data []byte

	var bf bytes.Buffer
	if bc, err := bf.ReadFrom(resp.Body); err == nil {
		log.Println(fmt.Sprintf("reading %d byte from response body", bc))
		data = bf.Bytes()
	}
	// or use ioutil:
	// data, err = ioutil.ReadAll(resp.Body)

	err = proto.Unmarshal(data, &hello)
	if err != nil {
		panic(err)
	}
	// fmt.Println(data)
	fmt.Println(hello.String())
}
