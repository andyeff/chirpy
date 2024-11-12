package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

func validateChirp(w http.ResponseWriter, r *http.Request) {
	type chirps struct {
		Body string `json:"body"`
	}
	type returnBody struct {
		Error string `json:"error,omitempty"`
		// Body        string `json:"body,omitempty"`
		CleanedBody string `json:"cleaned_body,omitempty"`
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

	chirpFiltered := profanityFilter(chirp.Body)

	responseJson(w, http.StatusOK, returnBody{
		CleanedBody: chirpFiltered,
	})
}

func profanityFilter(chirp string) string {
	profaneWords := []string{"kerfuffle", "sharbert", "fornax"}
	// chirpLower := strings.ToLower(chirp)
	chirpSplit := strings.Split(chirp, " ")
	chirpNew := []string{}
	for _, word := range chirpSplit {
		for _, profane := range profaneWords {
			if strings.EqualFold(word, profane) {
				word = "****"
			}
		}
		chirpNew = append(chirpNew, word)
	}
	chirpFiltered := strings.Join(chirpNew, " ")
	return chirpFiltered
}
