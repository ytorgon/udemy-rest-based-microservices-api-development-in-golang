package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
)

const contentType = "Content-Type"
const applicationXml = "application/xml"
const applicationJson = "application/json"

type Customer struct {
	Name string `json:"full_name" xml:"full_name"`
	Addr string `json:"city" xml:"city"`
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello there!")
}

func getCustomers(w http.ResponseWriter, r *http.Request) {
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
