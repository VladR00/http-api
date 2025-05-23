package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	cors "webrestapi/internal/cors"
	storage "webrestapi/internal/storage"
)

func HandlerGETRandomQuote(w http.ResponseWriter, r *http.Request) {
	cors.EnableCors(w)
	fmt.Println("gigi")
}

func HandlerQuote(w http.ResponseWriter, r *http.Request) {
	cors.EnableCors(w)

	switch r.Method {
	case "POST":
		HPOSTQuote(w, r)
	case "GET":
		HGETQuote(w, r)
	case "DELETE":
		HDELETEQuote(w, r)
	}

}

func HPOSTQuote(w http.ResponseWriter, r *http.Request) {
	cors.EnableCors(w)
	response := map[string]string{"error": "Only POST method allowed"}

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}

	data := storage.Quotes{}

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response = map[string]string{"error": "Bad request"}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}
}

func HGETQuote(w http.ResponseWriter, r *http.Request) {
	cors.EnableCors(w)

}

func HDELETEQuote(w http.ResponseWriter, r *http.Request) {
	cors.EnableCors(w)

}
