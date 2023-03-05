package handlers

import (
	"fmt"
	"net/http"
)

func (h *Handlers) HandleHealth(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Running! :)")
	default:
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Method not allowed")
	}
}
