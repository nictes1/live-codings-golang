package main

import (
	"fmt"
	"net/http"
)

func holaHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hola\n")
}

func main() {
	http.HandleFunc("/hola", holaHandler)
	http.ListenAndServe(":8080", nil)
}
