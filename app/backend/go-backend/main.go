package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello from Go backend!")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Go backend running on :8080")
	http.ListenAndServe(":8080", nil)
}
