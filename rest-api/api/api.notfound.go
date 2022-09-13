package api

import (
	"encoding/json"
	"net/http"
)

func ErrorNotFound(w http.ResponseWriter, r *http.Request) {
	var errorNotFound *DefaultApiResponse = &DefaultApiResponse{
		Status:  "404",
		Message: "Welcome to Talismo, please double check your endpoint, ths is not it!",
	}

	jsonResponse, _ := json.Marshal(errorNotFound)
	w.WriteHeader(http.StatusNotFound)
	w.Write(jsonResponse)
}
