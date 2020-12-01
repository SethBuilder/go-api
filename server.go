package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"
)

// Wine defines what attributes each wine must include
type Wine struct {
	ID      string  `json:"id"`
	Name    string  `json:"name"`
	Year    int     `json:"year"`
	Price   float32 `json:"price"`
	Region  string  `json:"region"`
	Country string  `json:"country"`
}

type wineHandlers struct {
	sync.Mutex
	store map[string]Wine
}

func (h *wineHandlers) wines(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		h.get(w, r)
		return
	case "POST":
		h.post(w, r)
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("method not allowed"))
		return
	}
}
func (h *wineHandlers) get(w http.ResponseWriter, r *http.Request) {
	wines := make([]Wine, len(h.store))
	h.Lock()
	i := 0
	for _, wine := range h.store {
		wines[i] = wine
		i++
	}
	h.Unlock()
	jsonBytes, err := json.Marshal(wines)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(([]byte(err.Error())))
		return
	}
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func (h *wineHandlers) post(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := ioutil.ReadAll((r.Body))
	defer r.Body.Close()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(([]byte(err.Error())))
		return
	}

	ct := r.Header.Get("content-type")
	if ct != "application/json" {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		w.Write(([]byte(fmt.Sprintf("need content-type 'application/json', but got '%s'", ct))))
		return
	}

	var wine Wine
	err = json.Unmarshal(bodyBytes, &wine)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(([]byte(err.Error())))
		return
	}
	wine.ID = fmt.Sprintf("%d", time.Now().UnixNano())
	h.Lock()
	h.store[wine.ID] = wine
	w.WriteHeader(http.StatusCreated)
	w.Write(([]byte("success")))
	h.Unlock()
}

func newWineHandlers() *wineHandlers {
	return &wineHandlers{
		store: map[string]Wine{},
	}
}

func main() {
	listenAddr := ":8000"

	wineHandlers := newWineHandlers()
	http.HandleFunc("/wines", wineHandlers.wines)

	log.Println("Server is ready to handle requests at", listenAddr)

	if err := http.ListenAndServe(listenAddr, nil); err != nil {
		log.Fatalf("Could not listen on %s: %v\n", listenAddr, err)
	}
}
