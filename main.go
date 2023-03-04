package main

import (
	"favcrypto/data"
	"favcrypto/handlers"
	"fmt"
	"log"
	"net/http"
)

func init() {
	currencies := data.GetDataFromDB()
	data.UpdateDataInDB(currencies)
}

func main() {
	http.HandleFunc("/", handlers.HandleRequest)
	fmt.Println("Server is running on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Error starting server: %q", err)
	}
}
