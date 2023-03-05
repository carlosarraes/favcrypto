package handlers

import (
	"encoding/json"
	"favcrypto/data"
	"fmt"
	"log"
	"net/http"
)

func HandleRequest(w http.ResponseWriter, _ *http.Request) {
	currencies := data.DB.GetDataFromDB()
	currenciesData := currencies.ToCurrenciesData()

	jsonData, err := json.Marshal(currenciesData)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"message": "Error marshalling data"}`)
		log.Fatalf("Error marshalling data: %q", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_, err = w.Write(jsonData)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"message": "Error writing data"}`)
		log.Fatalf("Error writing data: %q", err)
	}
}
