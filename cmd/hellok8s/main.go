package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		name, err := os.Hostname()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Fprintf(w, "Hello K8S\n---------\nHost: %s\nRequest path: %s\n", name, r.URL.Path)
	})

	log.Println("listening on 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
