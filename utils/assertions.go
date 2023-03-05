package utils

import (
	"testing"
)

func AssertResponse(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func AssertStatus(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func AssertContentType(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func AssertLength(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
