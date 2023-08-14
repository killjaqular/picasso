package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Picasso")
}

func main() {
	http.HandleFunc("/", helloHandler)
	fmt.Println("Picasso server on port 9000...")
	http.ListenAndServe(":9000", nil)
}
