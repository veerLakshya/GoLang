package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Println(resp, "Hello World!")
	})

	const serverAddr string = "127.0.0.1:4000"

	err := http.ListenAndServe(serverAddr, nil)
	if err != nil {
		log.Fatal("error starting server: ", err)
	}
}
