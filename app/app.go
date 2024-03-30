package app

import (
	"log"
	"net/http"
)

func Start() {
	mux := http.NewServeMux()
	mux.HandleFunc("/greet", greet)
	mux.HandleFunc("/customers", getCustomers)

	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}
