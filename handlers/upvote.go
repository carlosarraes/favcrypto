package handlers

import (
	"favcrypto/data"
	"fmt"
	"net/http"
	"strings"
)

func HandleUpvote(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/upvote/")
	check := data.DB.UpdateFavorite(true, strings.ToUpper(path))

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
		fmt.Fprintf(w, `{"message": "Upvoted"}`)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, `{"message": "Invalid coin"}`)
	}
}
