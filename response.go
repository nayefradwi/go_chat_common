package gochatcommon

import (
	"encoding/json"
	"net/http"
)

func WriteErrorResponse(w http.ResponseWriter, err *BaseError) {
	response := err.GenerateResponse()
	w.WriteHeader(err.Status)
	w.Write(response)
}

func WriteEmptyCreatedResponse(w http.ResponseWriter, m string) {
	w.WriteHeader(http.StatusCreated)
	body := make(map[string]string)
	body["status"] = "OK"
	body["message"] = m
	json.NewEncoder(w).Encode(body)
}

func WriteEmptySuccessResponse(w http.ResponseWriter, m string) {
	body := make(map[string]string)
	body["status"] = "OK"
	body["message"] = m
	json.NewEncoder(w).Encode(body)
}
