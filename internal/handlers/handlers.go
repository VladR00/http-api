package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	cors "webrestapi/internal/cors"
	storage "webrestapi/internal/storage"
)

func HandlerGETRandomQuote(w http.ResponseWriter, r *http.Request) {
	cors.EnableCors(w)
	fmt.Println("gigi")
}

func HandlerQuote(w http.ResponseWriter, r *http.Request) {
	cors.EnableCors(w)
	fmt.Println("Handler Quote")
	switch r.Method {
	case "POST":
		fmt.Println("POST request")
		HPOSTQuote(w, r)
	case "GET":
		fmt.Println("GET request")
		HGETQuote(w, r)
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

	var data struct {
		Author string `json:"author"`
		Quote  string `json:"quote"`
	}

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response = map[string]string{"error": "Bad request"}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		fmt.Println(err)
		return
	}

	storage.MapMutex.Lock()
	quote := storage.Quotes{
		ID:     len(storage.MapByID) + 1,
		Author: data.Author,
		Quote:  data.Quote,
	}
	quote.MapCreate()
	storage.MapMutex.Unlock()

	w.WriteHeader(http.StatusOK)
	response = map[string]string{"message": "Quote successfully added"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
	fmt.Println(storage.MapByID)
}

func HGETQuote(w http.ResponseWriter, r *http.Request) {
	cors.EnableCors(w)
	response := map[string]string{"error": "Only GET method allowed"}

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}
	var quotes []storage.Quotes

	for _, v := range storage.MapByID {
		quotes = append(quotes, v)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(quotes) // curl http://localhost:8080/quotes | jq - format output
}

func HandlerDELETEQuote(w http.ResponseWriter, r *http.Request) {
	cors.EnableCors(w)
	response := map[string]string{"error": "Only DELETE method allowed"}

	if r.Method != http.MethodDelete {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}

	idstr := strings.TrimPrefix(r.URL.Path, "/quotes/")
	fmt.Println("delete")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Header().Set("Content-Type", "application/json")
		response = map[string]string{"error": "Bad option. You should use the ID number that needs to be deleted."}
		json.NewEncoder(w).Encode(response)
		return
	}

	deleting, exist := storage.MapByID[id]

	if !exist {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		response = map[string]string{"error": "ID isn't found. Try another."}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	response = map[string]string{"message": fmt.Sprintf("%s quote was removed with id: %d", deleting.Author, deleting.ID)}
	storage.MapMutex.Lock()
	deleting.MapDelete()
	storage.MapMutex.Unlock()
	json.NewEncoder(w).Encode(response)
}
