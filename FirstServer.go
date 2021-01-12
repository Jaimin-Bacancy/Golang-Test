package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/index", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "index handler")
	})

	http.HandleFunc("/contact", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "contact handler")
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Root handler")
	})

	fmt.Println("Server started at 8080")
	http.ListenAndServe(":8080", nil)

}
