package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", router)

	err := http.ListenAndServe("0.0.0.0:9090", nil)

	if err != nil {
		log.Fatal("server fatal:", err)
	}
}

func router(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println("GET Request------------------------")
	fmt.Println(r.Form)
	fmt.Println(r.URL.Path)
	fmt.Println(r.URL.Scheme)
	for i, v := range r.Form {
		fmt.Println("key", i, "value", v)
	}
	fmt.Fprintf(w, "hello, go~")
	fmt.Println("END Request------------------------")
}
