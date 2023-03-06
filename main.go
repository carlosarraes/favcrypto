package main

import (
	"favcrypto/data"
	"favcrypto/handlers"
	"fmt"
	"log"
	"net/http"
)

func init() {
	data.DB.InitDB()
	currencies := data.DB.GetDataFromDB()
	data.DB.UpdatePrices(currencies)
}

func main() {
	port := ":8080"

	h := handlers.NewHandlers(&data.DB)

	http.HandleFunc("/upvote/", handleCORS(h.HandleUpvote))
	http.HandleFunc("/downvote/", handleCORS(h.HandleDownvote))
	http.HandleFunc("/getcoins/", handleCORS(h.HandleRequest))
	http.HandleFunc("/", handleCORS(h.HandleHealth))

	fmt.Printf("[GO] Server is running on port%s\n", port)

	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Error starting server: %q", err)
	}
}

func handleCORS(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next(w, r)
	}
}
