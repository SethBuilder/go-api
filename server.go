package main

import (
	"encoding/json"
	"net/http"
)

// Wine has the attributes of each wine
type Wine struct {
	Name    string  `json:"name"`
	Year    int     `json:"year"`
	Price   float32 `json:"price"`
	Region  string  `json:"region"`
	Country string  `json:"country"`
}

type wineHandlers struct {
	store map[string]Wine
}

func (h *wineHandlers) get(w http.ResponseWriter, r *http.Request) {
	wines := make([]Wine, len(h.store))
	i := 0
	for _, wine := range h.store {
		wines[i] = wine
	}
	jsonBytes, err := json.Marshal(wines)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(([]byte(err.Error())))
	}
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func newWineHandlers() *wineHandlers {
	return &wineHandlers{
		store: map[string]Wine{
			"id1": Wine{
				Name:   "Saint-Émilion Grand Cru (Premier Grand Cru Classé)",
				Year:   2006,
				Price:  2524.68,
				Region: "Saint-Émilion Grand Cru",
			},
		},
	}
}

func main() {
	wineHandlers := newWineHandlers()
	http.HandleFunc("/wines", wineHandlers.get)
	http.ListenAndServe(":8000", nil)
}
