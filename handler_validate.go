package main

import (
	"encoding/json"
	"net/http"
)

func handlerValidate(w http.ResponseWriter, req *http.Request) {
	type parameter struct {
		Body string `json:"body"`
	}

	type returnVals struct {
		Valid bool `json:"valid"`
	}

	decoder := json.NewDecoder(req.Body)
	param := parameter{}
	err := decoder.Decode(&param)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
		return
	}

	const maxChirpLength = 140
	if len(param.Body) > maxChirpLength {
		respondWithError(w, http.StatusBadRequest, "Chirp is too long", nil)
		return
	}
	respondWithJSON(w, http.StatusOK, returnVals{
		Valid: true,
	})
}