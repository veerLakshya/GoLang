package handlers

import (
	"log"
	"net/http"

	"github.com/veerLakshya/my-go-all/4.examples/micro-services/data"
)

// Products is a http.Handler
type Products struct {
	l *log.Logger
}

// NewProducts creates a products handler with a logger
func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

// Main entry point for the handler and satisfies the http.Handler interface
func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.getProducts(rw, r)
		return
	}

	if r.Method == http.MethodPost {
		p.addProduct(rw, r)
		return
	}

	if r.Method == http.MethodPut {
		// expect id in uri
		// p := r.URL.Path
		// p.updateProduct(rw, r)
		return
	}

	// Catch All
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

// getProducts returns the products from the data source
func (p *Products) getProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET Products")
	lp := data.GetProducts()
	// jsonlp, err := json.Marshal(lp)

	// if err != nil {
	// 	p.l.Println("Error marshalling data: ", err)
	// 	http.Error(rw, "Unable to marshal Json", http.StatusInternalServerError)
	// 	return
	// }
	// rw.Write(jsonlp)

	// BETTER IN TERMS OF OPTIMISATION than json.Marshal
	err := lp.ToJson(rw)

	if err != nil {
		p.l.Println("Error marshalling data: ", err)
		http.Error(rw, "Unable to marshal Json", http.StatusInternalServerError)
		return
	}
}

func (p *Products) addProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST Product")

	prod := &data.Product{}
	err := prod.FromJSON(r.Body)
	if err != nil {
		p.l.Println("Unable to unmarshal Json", err)
		// rw.WriteHeader(http.StatusBadRequest)
		http.Error(rw, "Unable to unmarshal Json", http.StatusBadRequest)
		return
	}

	p.l.Printf("Prod: %#v", prod)

	data.AddProduct(prod)
}
