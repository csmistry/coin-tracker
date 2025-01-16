package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Received request")
	})

	fmt.Println("Serving requests on port :8080")
	http.ListenAndServe(":8080", nil)
}
