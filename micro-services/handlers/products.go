package handlers

import (
	"log"
	"my-go-files/micro-services/data"
	"net/http"
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

	if r.Method == http.MethodPut {
		p.updateProducts(rw, r)
		return
	}

	// Catch All
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

// getProducts returns the products from the data source
func (p *Products) getProducts(rw http.ResponseWriter, r *http.Request) {
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

func (p *Products) updateProducts(rw http.ResponseWriter, r *http.Request) {

}
