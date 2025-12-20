package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/orders", func(http.ResponseWriter, *http.Request) {
		log.Println("Handling incoming orders")
	})
	http.HandleFunc("/users", func(http.ResponseWriter, *http.Request) {
		log.Println("Handling users")
	})
	port := 3000

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		log.Fatalln("Could not start server: ", err)
	}
	log.Println("server is running on port: ", port)
}
