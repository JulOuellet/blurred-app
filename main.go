package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, SportLight! test")
	})

	fmt.Println("Server ir running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
