package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func responseJson(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	dat, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error marshalliing JSON: %s", err)
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(code)
	w.Write(dat)
}

func responseError(w http.ResponseWriter, code int, msg string, err error) {
	if err != nil {
		log.Println(err)
	}
	if code > 499 {
		log.Printf("500 Error: %s", msg)
	}
	type errorResponse struct {
		Error string `json:"error"`
	}
	responseJson(w, code, errorResponse{
		Error: msg,
	})
}
