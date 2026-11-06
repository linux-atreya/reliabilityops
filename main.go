package main

import (
	"fmt"
	"log"
	"net/http"
)

func healthzHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Ok")
}

func main() {
	http.HandleFunc("/healthz", healthzHandler)

	log.Println("server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}