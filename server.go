package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.Handle("/", http.FileServer(http.Dir(".")))

	err := http.ListenAndServe(":4000", mux)
	if err != nil {
		log.Fatalf("Failed to listen:  %s", err)
	}
}
