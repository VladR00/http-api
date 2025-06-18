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
	if r.Method != http.MethodPost {
		DefaultResponse{Type: "", Message: r.Method}.Response(w, 0)
		return
	}

	var data storage.Data

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		DefaultResponse{Type: "Error", Message: "Bad request, use seconds."}.Response(w, http.StatusBadRequest)
		return
	}

	id := data.AddTask()

	DefaultResponse{Type: "Message", Message: fmt.Sprintf("Task with ID %d successfully added", id)}.Response(w, http.StatusOK)
}

func HGETTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		DefaultResponse{Type: "", Message: r.Method}.Response(w, 0)
		return
	}

	tasks := storage.GetTasks()
	TaskOutputResopnse{tasks}.Response(w)
}

func HandlerDELETETask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		DefaultResponse{Type: "", Message: r.Method}.Response(w, 0)
		return
	}

	idstr := strings.TrimPrefix(r.URL.Path, "/task/")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		DefaultResponse{Type: "Error", Message: "Bad option. You should use the ID number that needs to be deleted."}.Response(w, http.StatusUnprocessableEntity)
		return
	}

	if v, exist := storage.MapByID[id]; exist {
		v.MapDelete()
	} else {
		DefaultResponse{Type: "Error", Message: fmt.Sprintf("Task with ID %d isn't found. Try another.", id)}.Response(w, http.StatusBadRequest)
		return
	}

	DefaultResponse{Type: "Message", Message: fmt.Sprintf("Task with ID %d was removed", id)}.Response(w, http.StatusOK)
}
