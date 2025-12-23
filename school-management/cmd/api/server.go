package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := ":3000"

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("/  handler", r.Method)

		// fmt.Fprintf(w, "Hello Root Route")
		w.Write([]byte("Hello Root Route"))

	})
	http.HandleFunc("/teachers", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("/teachers  handler", r.Method)

		if r.Method == http.MethodGet {
			w.Write([]byte("Hello GET on Root Route"))
			return
		}

		// fmt.Fprintf(w, "Hello Root Route")
		w.Write([]byte("Hello Teachers Route"))
	})

	http.HandleFunc("/students", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("/students  handler", r.Method)

		// fmt.Fprintf(w, "Hello Root Route")
		w.Write([]byte("Hello Students Route"))
	})

	http.HandleFunc("/execs", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("/execs  handler", r.Method)

		// fmt.Fprintf(w, "Hello Root Route")
		w.Write([]byte("Hello Execs Route"))
	})

	fmt.Println("Server is running on port: ", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatalln("Error starting the server: ", err)
	}
}
