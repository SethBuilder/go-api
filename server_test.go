package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetWines(t *testing.T) {
	t.Run("returns 200 on /wines", func(t *testing.T) {
		r, _ := http.NewRequest(http.MethodGet, "/wines", nil)
		w := httptest.NewRecorder()

		wineHandlers := newWineHandlers()
		wineHandlers.get(w, r)
		t.Helper()
		if w.Code != http.StatusOK {
			t.Errorf("did not get correct status, got %d, want %d", w.Code, http.StatusOK)
		}
	})

	t.Run("returns a list of wines", func(t *testing.T) {
		r, _ := http.NewRequest(http.MethodGet, "/wines", nil)
		w := httptest.NewRecorder()

		wineHandlers := newWineHandlers()
		wineHandlers.get(w, r)

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
}
