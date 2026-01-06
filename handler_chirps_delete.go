package main

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/kenzierivan/chirpy/internal/auth"
	"github.com/kenzierivan/chirpy/internal/database"
)

func (cfg *apiConfig) handlerChirpsDelete(w http.ResponseWriter, req *http.Request) {
	chirpID, err := uuid.Parse(req.PathValue("chirpID"))
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't parse chirp_id", err)
		return
	}	
	accessToken, err := auth.GetBearerToken(req.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Couldn't find access token", err)
		return
	}
	userID, err := auth.ValidateJWT(accessToken, cfg.tokenSecret)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Couldn't validate access token", err)
		return 
	}

	chirp, err := cfg.db.GetChirp(req.Context(), chirpID)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Couldn't find chirp", err)
		return
	}

	if chirp.UserID != userID {
		respondWithError(w, http.StatusForbidden, "Forbidden", nil)
		return
	}

	err = cfg.db.DeleteChirp(req.Context(), database.DeleteChirpParams{
		ID: chirpID,
		UserID: userID,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't delete chirp", err)
	}
	w.WriteHeader(http.StatusNoContent)
}