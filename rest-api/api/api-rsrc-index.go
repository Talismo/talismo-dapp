package api

import (
	"encoding/json"
	"net/http"
)

func ResourceIndex(w http.ResponseWriter, r *http.Request) {
	var errorNotFound *DefaultApiResponse = &DefaultApiResponse{
		Status:  "200",
		Message: "Welcome to Talismo!",
	}

	jsonResponse, _ := json.Marshal(errorNotFound)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}
