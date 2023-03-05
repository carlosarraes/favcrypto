package main

import (
	"favcrypto/data"
	"favcrypto/handlers"
	"fmt"
	"log"
	"net/http"
	"os"
)

func init() {
	data.DB.InitDB()
	currencies := data.DB.GetDataFromDB()
	data.DB.UpdatePrices(currencies)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	h := handlers.NewHandlers(&data.DB)

	http.HandleFunc("/upvote/", h.HandleUpvote)
	http.HandleFunc("/getcoins/", h.HandleRequest)
	http.HandleFunc("/", h.HandleHealth)

	fmt.Printf("Server is running on port%s\n", port)

	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Error starting server: %q", err)
	}
}
