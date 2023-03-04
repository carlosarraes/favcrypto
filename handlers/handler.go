package handlers

import (
	"fmt"
	"net/http"
)

func HandleRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Request: %q", r.Method)
	fmt.Fprintf(w, "Hello, World!")
}
