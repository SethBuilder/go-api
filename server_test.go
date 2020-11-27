package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetWines(t *testing.T) {
	t.Run("returns 200 on GET /wines", func(t *testing.T) {
		r, _ := http.NewRequest(http.MethodGet, "/wines", nil)
		w := httptest.NewRecorder()

		wineHandlers := newWineHandlers()
		wineHandlers.wines(w, r)

		if w.Code != http.StatusOK {
			t.Errorf("did not get correct status, got %d, want %d", w.Code, http.StatusOK)
		}
	})

	t.Run("returns a list of wines", func(t *testing.T) {
		r, _ := http.NewRequest(http.MethodGet, "/wines", nil)
		w := httptest.NewRecorder()

		wineHandlers := newWineHandlers()
		wineHandlers.wines(w, r)

		var wine []Wine

		err := json.NewDecoder(w.Body).Decode(&wine)

		if err != nil {
			t.Fatalf("Unable to parse response from server %q into slice of Player, '%v'", w.Body, err)
		}

		got := fmt.Sprintf("%T", wine)
		want := "[]main.Wine"

		if got != want {
			t.Errorf("did not get correct status, got %q, want %q", got, want)
		}
	})

	t.Run("returns 201 on POST /wines", func(t *testing.T) {

		const jsonStream = `
		{"name": "Saint-Émilion Grand Cru (Premier Grand Cru Classé)",
		"year": 2006,
		"price": 2524.68,
		"region": "Saint-Émilion Grand Cru",
		"country": "France"}
		`
		requestBody := strings.NewReader(jsonStream)
		r, _ := http.NewRequest(http.MethodPost, "/wines", requestBody)
		w := httptest.NewRecorder()
		r.Header.Set("content-type", "application/json")
		wineHandlers := newWineHandlers()
		wineHandlers.wines(w, r)

		if w.Code != http.StatusCreated {
			t.Errorf("did not get correct status, got %d, want %d", w.Code, http.StatusCreated)
		}
	})
}
