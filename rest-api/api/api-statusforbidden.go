package api

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func ErrorForbidden(w http.ResponseWriter, r *http.Request) {
	var errorForbidden *DefaultApiResponse = &DefaultApiResponse{
		Status:  strconv.Itoa(http.StatusForbidden),
		Message: "Sorry, you need authorization in order to use talismo api",
	}

	jsonResponse, _ := json.Marshal(errorForbidden)
	w.WriteHeader(http.StatusForbidden)
	w.Write(jsonResponse)
}
