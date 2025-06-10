package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	cors "httpapi/internal/cors"
	storage "httpapi/internal/storage"
)

func HandlerTask(w http.ResponseWriter, r *http.Request) {
	cors.EnableCors(w)
	fmt.Println(r.Method)
	switch r.Method {
	case "POST":
		HPOSTTask(w, r)
	case "GET":
		HGETTask(w, r)
	}
}

func HPOSTTask(w http.ResponseWriter, r *http.Request) {
	cors.EnableCors(w)
	response := map[string]string{"error": "Only POST method allowed"}

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}

	var data struct {
		Duration int64 `json:"duration"`
	}

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response = map[string]string{"error": "Bad request, use seconds."}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}

	storage.MapMutex.Lock() //lock for ID
	task := storage.Task{
		ID:          len(storage.MapByID) + 1,
		Date:        time.Now().Unix(),
		AllDuration: data.Duration,
		Remaining:   data.Duration,
		IsComplete:  false,
	}
	task.MapCreate()
	storage.MapMutex.Unlock()

	go func() { // remaining -= 2;update.
		for task.Remaining > 0 {
			if _, exists := storage.MapByID[task.ID]; exists {
				time.Sleep(time.Second * 2)
				task.Remaining = task.Remaining - 2
				task.MapUpdate()
			} else {
				break
			}
		}
		if task.Remaining <= 0 {
			fmt.Printf("Task %d successfully end\n", task.ID)
			task.IsComplete = true
			task.MapUpdate()
		} else {
			fmt.Printf("Task %d was deleted\n", task.ID)
		}
	}()

	w.WriteHeader(http.StatusOK)
	response = map[string]string{"message": fmt.Sprintf("Task with ID %d successfully added", task.ID)}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func HGETTask(w http.ResponseWriter, r *http.Request) {
	cors.EnableCors(w)
	response := map[string]string{"error": "Only GET method allowed"}

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}
	var tasks []storage.TaskOutput

	for _, v := range storage.MapByID {
		percent := (1 - (float32(v.Remaining) / float32(v.AllDuration))) * 100
		task := storage.TaskOutput{
			ID:         v.ID,
			Date:       time.Unix(v.Date, 0).Format("2006-01-02 15:04"),
			Remaining:  v.Remaining,
			Percent:    int(percent),
			IsComplete: v.IsComplete,
		}
		tasks = append(tasks, task)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks) // curl http://localhost:8080/task | jq - format output
}

func HandlerDELETETask(w http.ResponseWriter, r *http.Request) {
	cors.EnableCors(w)
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
	response = map[string]string{"message": fmt.Sprintf("Task with ID %d was removed", deleting.ID)}
	storage.MapMutex.Lock()
	deleting.MapDelete()
	storage.MapMutex.Unlock()
	json.NewEncoder(w).Encode(response)
	fmt.Printf("Task with ID %d was removed\n", deleting.ID)
}
