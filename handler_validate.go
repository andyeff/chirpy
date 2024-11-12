package main

import (
	"encoding/json"
	"net/http"
)

func validateChirp(w http.ResponseWriter, r *http.Request) {
	type chirps struct {
		Body string `json:"body"`
	}
	type returnBody struct {
		Error string `json:"error,omitempty"`
		Valid bool   `json:"valid,omitempty"`
	}
	decoder := json.NewDecoder(r.Body)
	chirp := chirps{}
	err := decoder.Decode(&chirp)
	if err != nil {
		responseError(w, http.StatusInternalServerError, "couldn't decode parameters", err)
		return
	}

	const maxChirpLength = 140
	if len(chirp.Body) > maxChirpLength {
		responseError(w, http.StatusBadRequest, "Chirp is too long", nil)
		return
	}
	if len(chirp.Body) == 0 {
		responseError(w, http.StatusBadRequest, "Chirp is empty", nil)
		return
	}

	responseJson(w, http.StatusOK, returnBody{
		Valid: true,
	})
}
