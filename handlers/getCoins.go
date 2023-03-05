package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func (h *Handlers) HandleRequest(w http.ResponseWriter, r *http.Request) {
	currencies := h.db.GetDataFromDB()
	currenciesData := currencies.ToCurrenciesData()

	var jsonData []byte
	var err error

	switch r.Method {
	case http.MethodGet:
		jsonData, err = json.Marshal(currenciesData)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, `{"message": "Error marshalling data"}`)
			log.Printf("Error marshalling data: %q", err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	default:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, `{"message": "Method not allowed"}`)
	}

	_, err = w.Write(jsonData)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"message": "Error writing data"}`)
		log.Printf("Error writing data: %q", err)
	}
}
