package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello K8S, request path: %s\n", r.URL.Path)
	})

	log.Println("listening on 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
