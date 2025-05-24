package utils

import (
	"encoding/json"
	"net/http"
)

type SuccessResponse struct {
	Message string		`json:"message"`
	Data 	interface{}	`json:"data,omitempty"`	
}

type ErrorResponse struct {
	Message string	`json:"message"`
}

func RespondWithSuccess (w http.ResponseWriter, code int, message string, data interface{}) {
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(SuccessResponse{message,data})
}

func RespondWithRedirect(w http.ResponseWriter, r *http.Request, url string, code int) {
    http.Redirect(w, r, url, code)
}


func RespondWithError (w http.ResponseWriter, code int, message string) {
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(ErrorResponse{message})
}