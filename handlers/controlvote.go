package handlers

import (
	"fmt"
	"net/http"
	"strings"
)

func (h *Handlers) HandleUpvote(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/upvote/")
	check := h.db.UpdateFavorite(true, strings.ToUpper(path))

	switch r.Method {
	case http.MethodPatch:
		writeHeader(w, true, check)
	default:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, `{"message": "Method not allowed"}`)
	}

	fmt.Printf("Upvoted: %q\n", strings.ToUpper(path))
}

func (h *Handlers) HandleDownvote(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/downvote/")
	check := h.db.UpdateFavorite(false, strings.ToUpper(path))

	switch r.Method {
	case http.MethodPatch:
		writeHeader(w, false, check)
	default:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, `{"message": "Method not allowed"}`)
	}

	fmt.Printf("Upvoted: %q\n", strings.ToUpper(path))
}

func writeHeader(w http.ResponseWriter, t bool, c int64) {
	if c > 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		s := `{"message": "Upvoted"}`
		if !t {
			s = `{"message": "Downvoted"}`
		}
		fmt.Fprint(w, s)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, `{"message": "Invalid coin"}`)
	}
}
