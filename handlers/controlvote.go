package handlers

import (
	"fmt"
	"net/http"
	"strings"
)

func (h *Handlers) HandleUpvote(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/upvote/")
	check := h.db.UpdateFavorite(strings.ToUpper(path))

	switch r.Method {
	case http.MethodPatch:
		writeHeader(w, check)
	default:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, `{"message": "Method not allowed"}`)
	}

	fmt.Printf("Upvoted: %q\n", strings.ToUpper(path))
}

func writeHeader(w http.ResponseWriter, c int64) {
	if c > 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		s := `{"message": "Upvoted"}`
		fmt.Fprint(w, s)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, `{"message": "Invalid coin"}`)
	}
}
