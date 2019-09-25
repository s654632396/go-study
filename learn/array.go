package main

import "fmt"

type PersonInfo struct {
	ID   string
	Name string
	desc string
}

func main() {

	var PersonDB map[string]PersonInfo
	PersonDB = make(map[string]PersonInfo)

	PersonDB["dong"] = PersonInfo{"10001", "dongchaofeng", "????"}
	PersonDB["poi"] = PersonInfo{"10002", "poi2", "kc"}

	person, ok := PersonDB["poi"]
	if ok {
		fmt.Println("found :", person)
	} else {
		fmt.Println("not found. ")
	}
}
