package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/http2"
)

func main() {

	http.HandleFunc("/orders", func(http.ResponseWriter, *http.Request) {
		log.Println("Handling incoming orders")
	})
	http.HandleFunc("/users", func(http.ResponseWriter, *http.Request) {
		log.Println("Handling users")
	})
	port := 3000

	// load the TLS cert and key
	cert := "cert.pem"
	key := "key.pem"

	// configure TLS
	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS12,
	}

	// create a custom server
	server := &http.Server{
		Addr:      fmt.Sprintf(":%d", port),
		Handler:   nil,
		TLSConfig: tlsConfig,
	}

	// Enable http2
	http2.ConfigureServer(server, &http2.Server{})

	log.Printf("server is running on port: %d", port)
	err := server.ListenAndServeTLS(cert, key)
	if err != nil {
		log.Fatalln("Could not start server: ", err)
	}

	// === HTTP 1.1 Server without TLS
	// err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	// if err != nil {
	// 	log.Fatalln("Could not start server: ", err)
	// }
	// log.Println("server is running on port: ", port)
}
