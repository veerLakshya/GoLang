package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello World!")
	})

	const serverAddr string = "127.0.0.1:4000"

	fmt.Println("Starting server at", serverAddr)
	err := http.ListenAndServe(serverAddr, nil)
	if err != nil {
		log.Fatal("error starting server: ", err)
	}
}
