package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	mw "school-management/internal/api/middlewares"
	"strings"
	// "time"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	City string `json:"city"`
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Hello Root Route")
	w.Write([]byte("Hello Root Route"))

}

func teachersHandler(w http.ResponseWriter, r *http.Request) {

	// PATH PARAMS --> teachers/{id}
	fmt.Println("/teachers  handler", r.Method, r.URL.Path)
	path := strings.TrimPrefix(r.URL.Path, "/teachers/")
	userId := strings.TrimSuffix(path, "/")
	fmt.Println("The id is: ", userId)

	// QUERY PARAMS --> teachers/?key=value&query=value2&sortby=email&sortorder=ASC
	fmt.Println("Query Params: ", r.URL.Query())
	queryParams := r.URL.Query()
	sortBy := queryParams.Get("sortby")
	sortOrder := queryParams.Get("sortorder")
	fmt.Println("Sort By: ", sortBy)
	fmt.Println("Sort Order: ", sortOrder)

	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("Hello GET on Teachers Route"))
	case http.MethodDelete:
		w.Write([]byte("Hello DELETE on Teachers Route"))
	case http.MethodPatch:
		w.Write([]byte("Hello PATCH on Teachers Route"))
	case http.MethodPost:
		w.Write([]byte("Hello POST on Teachers Route"))
	case http.MethodPut:
		w.Write([]byte("Hello PUT on Teachers Route"))
	}

	// fmt.Fprintf(w, "Hello Root Route")
	// w.Write([]byte("Hello Teachers Route"))
}

func studentsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("/students  handler", r.Method)

	w.Write([]byte("Hello Students Route"))
}

func execsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("/execs handler", r.Method)

	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("Hello GET on Execs Route"))
	case http.MethodDelete:
		w.Write([]byte("Hello DELETE on Execs Route"))
	case http.MethodPatch:
		w.Write([]byte("Hello PATCH on Execs Route"))
	case http.MethodPost:
		fmt.Println("Query:", r.URL.Query())
		fmt.Println("Query:", r.URL.Query().Get("name"))

		//Parse form data(necessary for x-www-form-urlencoded)
		err := r.ParseForm()
		if err != nil {
			return
		}
		fmt.Println("Form from POST method:", r.Form)
		w.Write([]byte("Hello POST on Execs Route"))
	case http.MethodPut:
		w.Write([]byte("Hello PUT on Execs Route"))
	}
}

func main() {
	port := ":3030"

	// cert := "cert.pem"
	// key := "key.pem"

	mux := http.NewServeMux()

	mux.HandleFunc("/", rootHandler)

	mux.HandleFunc("/teachers", teachersHandler)

	mux.HandleFunc("/students", studentsHandler)

	mux.HandleFunc("/execs", execsHandler)

	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS12,
	}

	// rl := mw.NewRateLimiter(5, time.Minute)

	// hppOptions := mw.HPPOptions{
	// 	CheckQuery:                  true,
	// 	CheckBody:                   true,
	// 	CheckBodyOnlyForContentType: "application/x-www-form-urlencoded",
	// 	Whitelist:                   []string{"sortBy", "name", "age", "class"},
	// }

	// secureMux := mw.Cors(rl.Middleware(mw.ResponseTime(mw.SecurityHeaders(mw.Compression(mw.Hpp(hppOptions)(mux))))))
	// function to properly chain middlewares
	// secureMux := applyMiddlewares(mux, mw.Hpp(hppOptions), mw.Compression, mw.SecurityHeaders, mw.ResponseTime, rl.Middleware, mw.Cors)
	secureMux := mw.SecurityHeaders(mux)

	//custom server
	server := &http.Server{
		Addr:      port,
		Handler:   secureMux,
		TLSConfig: tlsConfig,
	}

	fmt.Println("Server is running on port: ", port)
	err := server.ListenAndServe()
	// err := server.ListenAndServeTLS(cert, key)
	if err != nil {
		log.Fatalln("Error starting the server: ", err)
	}
}

// Middleware is a function that properly wraps an http.Handler with additional functionality (for chaining without cluttering)
type Middleware func(http.Handler) http.Handler

func applyMiddlewares(handler http.Handler, middlewares ...Middleware) http.Handler {
	for _, middleware := range middlewares {
		handler = middleware(handler)
	}
	return handler
}
