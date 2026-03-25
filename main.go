package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello World.")
	})

	addr := ":" + port
	log.Printf("Listening on %s...", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
