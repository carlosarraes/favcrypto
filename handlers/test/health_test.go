package handlers_test

import (
	"favcrypto/handlers"
	"favcrypto/utils"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleHealth(t *testing.T) {
	h := handlers.NewHandlers(nil)

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(h.HandleHealth)

	handler.ServeHTTP(rr, req)

	gotCode := rr.Code
	wantCode := http.StatusOK

	utils.AssertStatus(t, gotCode, wantCode)

	gotCt := rr.Header().Get("Content-Type")
	wantCt := "text/plain"

	utils.AssertContentType(t, gotCt, wantCt)

	got := rr.Body.String()
	want := "Running! :)"

	utils.AssertResponse(t, got, want)
}

func TestHandleHealth_MethodNotAllowed(t *testing.T) {
	h := handlers.NewHandlers(nil)

	req, err := http.NewRequest(http.MethodPost, "/", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(h.HandleHealth)

	handler.ServeHTTP(rr, req)

	gotCode := rr.Code
	wantCode := http.StatusMethodNotAllowed

	utils.AssertStatus(t, gotCode, wantCode)

	gotCt := rr.Header().Get("Content-Type")
	wantCt := "text/plain"

	utils.AssertContentType(t, gotCt, wantCt)

	got := rr.Body.String()
	want := "Method not allowed"

	utils.AssertResponse(t, got, want)
}
