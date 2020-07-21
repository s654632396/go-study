package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type Drink struct {
	Name     string
	TypeName string
}

func main() {
	write()
}

func write() {
	coffee := Drink{"luckin coffee", "coffee"}
	data, _ := json.Marshal(coffee)

	f, err := os.OpenFile("drink.json", os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = f.Truncate(0)
	if err != nil {
		log.Println(err)
	}
	f.WriteString(string(data))
}

func read() {

	f, err := os.OpenFile("drink.json", os.O_CREATE|os.O_RDONLY, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	data, _ := ioutil.ReadAll(f)
	log.Println(string(data))
	var coffee Drink
	json.Unmarshal(data, &coffee)
	log.Println(data)
}
