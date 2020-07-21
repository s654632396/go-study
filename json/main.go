package main

import (
	"encoding/json"
	"fmt"
	"io"
	//"io"
	"log"
	"os"
)

func main() {
	type User struct {
		FName  string `json:"FirstName"`
		LName  string `json:"LastName"`
		Remark string `json:"Remark,omitempty"`
	}

	f, err := os.Open("./json/data.json")
	defer f.Close()
	if err != nil {
		panic(err)
	}
	// dec := json.NewDecoder(f)
	var data []byte
	var buf []byte = make([]byte, 1024)
	for count, err := f.Read(buf); err != io.EOF; count, err = f.Read(buf) {
		if err != nil {
			log.Fatal(err)
		}
		// fmt.Printf("read %d bytes: %q\n", count, buf[:count])
		for i := 0; i < count; i++ {
			data = append(data, buf[i])
		}
	}

	var user User
	var userMap map[string]interface{}
	if err := json.Unmarshal([]byte(data), &user); err != nil {
		log.Println(err)
	}
	fmt.Printf("%#v \n", user)

	if err := json.Unmarshal([]byte(data), &userMap); err != nil {
		log.Println(err)
	}
	fmt.Printf("%+v \n", userMap)

	return
}
