package handlers_test

import (
	"encoding/json"
	"favcrypto/data"
	"favcrypto/handlers"
	"favcrypto/utils"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func TestHandleUpvoteAndGetCoins(t *testing.T) {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %q", err)
	}
	data.DB.InitDB()
	currencies := data.DB.GetDataFromDB()
	data.DB.UpdatePrices(currencies)

	h := handlers.NewHandlers(&data.DB)

	t.Run("Tests if upvoting works", func(t *testing.T) {
		req, err := http.NewRequest("PATCH", "/upvote/btc", nil)
		if err != nil {
			t.Fatalf("Failed to create request: %v", err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(h.HandleUpvote)

		handler.ServeHTTP(rr, req)

		gotCode := rr.Code
		wantCode := http.StatusOK

		utils.AssertStatus(t, gotCode, wantCode)

		gotCt := rr.Header().Get("Content-Type")
		wantCt := "application/json"

		utils.AssertContentType(t, gotCt, wantCt)

		got := rr.Body.String()
		want := `{"message": "Upvoted"}`

		utils.AssertResponse(t, got, want)
	})

	t.Run("Tests if upvoting doesnt work if the coin doesnt exist", func(t *testing.T) {
		req, err := http.NewRequest("PATCH", "/upvote/abc", nil)
		if err != nil {
			t.Fatalf("Failed to create request: %v", err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(h.HandleUpvote)

		handler.ServeHTTP(rr, req)

		gotCode := rr.Code
		wantCode := http.StatusNotFound

		utils.AssertStatus(t, gotCode, wantCode)

		gotCt := rr.Header().Get("Content-Type")
		wantCt := "application/json"

		utils.AssertContentType(t, gotCt, wantCt)

		got := rr.Body.String()
		want := `{"message": "Invalid coin"}`

		utils.AssertResponse(t, got, want)
	})

	t.Run("Tests if getCoins, get the correct length", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/getcoins", nil)
		if err != nil {
			t.Fatalf("Failed to create request: %v", err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(h.HandleRequest)

		handler.ServeHTTP(rr, req)

		gotCode := rr.Code
		wantCode := http.StatusOK

		utils.AssertStatus(t, gotCode, wantCode)

		gotCt := rr.Header().Get("Content-Type")
		wantCt := "application/json"

		utils.AssertContentType(t, gotCt, wantCt)

		// convert *bytes.Buffer to slice of maps
		var gotSlice []map[string]interface{}
		if err := json.NewDecoder(rr.Body).Decode(&gotSlice); err != nil {
			t.Fatalf("Failed to decode response body: %v", err)
		}
		got := len(gotSlice)
		want := 10

		utils.AssertLength(t, got, want)
	})

	t.Run("Tests if getCoins throws error if wrong method is used", func(t *testing.T) {
		req, err := http.NewRequest("PATCH", "/getcoins", nil)
		if err != nil {
			t.Fatalf("Failed to create request: %v", err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(h.HandleRequest)

		handler.ServeHTTP(rr, req)

		gotCode := rr.Code
		wantCode := http.StatusMethodNotAllowed

		utils.AssertStatus(t, gotCode, wantCode)

		gotCt := rr.Header().Get("Content-Type")
		wantCt := "application/json"

		utils.AssertContentType(t, gotCt, wantCt)

		got := rr.Body.String()
		want := `{"message": "Method not allowed"}`

		utils.AssertResponse(t, got, want)
	})
}
