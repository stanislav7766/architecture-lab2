package tools

import (
	"encoding/json"
	"log"
	"net/http"
)

type errorObject struct {
	Message string `json:"message"`
}

func WriteJsonBadRequest(rw http.ResponseWriter, message string) {
	writeJson(rw, http.StatusBadRequest, &errorObject{Message: message})
}

func WriteJsonCustomError(rw http.ResponseWriter, err error) {
	writeJson(rw, http.StatusBadRequest, &errorObject{Message: err.Error()})
}


func writeJson(rw http.ResponseWriter, status int, res interface{}) {
	rw.Header().Set("content-type", "application/json")
	rw.WriteHeader(status)
	err := json.NewEncoder(rw).Encode(res)
	if err != nil {
		log.Printf("Error writing response: %s", err)
	}
}
