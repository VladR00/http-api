package handlers

import (
	"encoding/json"
	"fmt"
	storage "httpapi/internal/storage"
	"net/http"
)

type DefaultResponse struct {
	Type    string `json:"type"`    // Error | Data | Message
	Message string `json:"message"` // Message
}
type TaskOutputResopnse struct {
	Task []storage.TaskOutput `json:"tasks"`
}

type CustomResponseWriter struct {
	http.ResponseWriter
}

func (response DefaultResponse) Response(w http.ResponseWriter, header int) {
	w.Header().Set("Content-Type", "application/json")
	if header == 0 {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(DefaultResponse{Type: "Error", Message: fmt.Sprintf("Only %s method allowed", response.Message)})
		return
	}
	w.WriteHeader(header)
	json.NewEncoder(w).Encode(response)
}

func (response TaskOutputResopnse) Response(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
