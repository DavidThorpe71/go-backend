package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETplayers(t *testing.T) {
	t.Run("Returns Pepper's score", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/players/Pepper", nil)
		response := httptest.NewRecorder()

		PlayerServer(response, request)

		got := response.Body.String()
		want := "20"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("Returns Floyd's score", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/players/Floyd", nil)
		response := httptest.NewRecorder()

		PlayerServer(response, request)

		got := response.Body.String()
		want := "50"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
