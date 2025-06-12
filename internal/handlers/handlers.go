package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	storage "httpapi/internal/storage"
)

func HandlerTask(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	switch r.Method {
	case "POST":
		HPOSTTask(w, r)
	case "GET":
		HGETTask(w, r)
	}
}

func HPOSTTask(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{"error": "Only POST method allowed"}

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}

	var data storage.Data

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response = map[string]string{"error": "Bad request, use seconds."}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}

	id := data.AddTask()

	w.WriteHeader(http.StatusOK)
	response = map[string]string{"message": fmt.Sprintf("Task with ID %d successfully added", id)}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func HGETTask(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{"error": "Only GET method allowed"}

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}

	tasks := storage.GetTasks()

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks) // curl http://localhost:8080/task | jq - format output
}

func HandlerDELETETask(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{"error": "Only DELETE method allowed"}

	if r.Method != http.MethodDelete {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}

	idstr := strings.TrimPrefix(r.URL.Path, "/task/")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Header().Set("Content-Type", "application/json")
		response = map[string]string{"error": "Bad option. You should use the ID number that needs to be deleted."}
		json.NewEncoder(w).Encode(response)
		return
	}

	if v, exist := storage.MapByID[id]; exist {
		v.MapDelete()
	} else {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		response = map[string]string{"error": fmt.Sprintf("Task with ID %d isn't found. Try another.", id)}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	response = map[string]string{"message": fmt.Sprintf("Task with ID %d was removed", id)}

	json.NewEncoder(w).Encode(response)
	fmt.Printf("Task with ID %d was removed\n", id)
}
