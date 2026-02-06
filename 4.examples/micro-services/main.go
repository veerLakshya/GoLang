package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/veerlakshya/my-go-all/4.examples/micro-services/handlers"
)

func main() {
	l := log.New(os.Stdout, "product-api ", log.LstdFlags)

	// Create a new serve mux
	sm := http.NewServeMux()

	// hh := handlers.NewHello(l)
	// gh := handlers.NewGoodBye(l)

	// Create handlers
	ph := handlers.NewProducts(l)

	// *Register Handlers
	// sm.Handle("/", hh)
	// sm.Handle("/goodbye", gh)
	sm.Handle("/", ph)

	// Create a new server
	s := &http.Server{
		Addr:         ":9090",
		Handler:      sm,                // default handler
		ErrorLog:     l,                 // set logger for the server
		ReadTimeout:  5 * time.Second,   // max time to read request from client
		WriteTimeout: 10 * time.Second,  // max time to write response to client
		IdleTimeout:  120 * time.Second, // max time to for connections using TCP keep-Alive
	}

	// starting server
	go func() {
		l.Println("Starting server on port 9090")
		err := s.ListenAndServe()
		if err != nil {
			log.Printf("Error starting server: %s\n", err)
			os.Exit(1)
		}
	}()

	// trap sigterm or interupt and gracefully shutdown the server
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	// Block until a signal is received.
	sig := <-c
	l.Println("Recieved terminate, graceful shutdown", sig)

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	ctx, _ := context.WithTimeout(context.Background(), time.Second*30)
	s.Shutdown(ctx)
}
