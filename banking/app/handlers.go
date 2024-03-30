package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

const contentType = "Content-Type"
const applicationXml = "application/xml"
const applicationJson = "application/json"
const customerId = "customerId"

type Customer struct {
	Name string `json:"full_name" xml:"full_name"`
	Addr string `json:"city" xml:"city"`
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello there!")
}

func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers := []Customer{
		{"Bob", "New York"},
		{"Alice", "Manchester"},
	}

	if w.Header().Get(contentType) == applicationXml {
		w.Header().Add(contentType, applicationXml)
		xml.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Add(contentType, applicationJson)
		json.NewEncoder(w).Encode(customers)
	}
}

func getCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprint(w, vars[customerId])
}

func createCustomer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Post request invoked")
}
