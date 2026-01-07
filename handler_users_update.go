package main

import (
	"encoding/json"
	"net/http"

	"github.com/kenzierivan/chirpy/internal/auth"
	"github.com/kenzierivan/chirpy/internal/database"
)

func (cfg *apiConfig) handlerUsersUpdate(w http.ResponseWriter, req *http.Request) {
	type parameters struct {
		Password string `json:"password"`
		Email string `json:"email"`
	}

	type response struct {
		User
	}

	accessToken, err := auth.GetBearerToken(req.Header)
	uuid, err := auth.ValidateJWT(accessToken, cfg.tokenSecret)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Couldn't validate access token", err)
		return
	}

	params := parameters{}
	decoder := json.NewDecoder(req.Body)
	err = decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
		return
	}

	user, err := cfg.db.GetUserByID(req.Context(), uuid)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Couldn't get user by id", err)
		return
	}

	hashed, err := auth.HashPassword(params.Password)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't hash password", err)
		return
	}

	err = cfg.db.UpdateUserEmailPassword(req.Context(), database.UpdateUserEmailPasswordParams{
		HashedPassword: hashed,
		Email: params.Email,
		ID: user.ID,
	})

	respondWithJSON(w, http.StatusOK, response{
		User: User{
			ID: user.ID,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
			Email: params.Email,
			IsChirpyRed: user.IsChirpyRed,
		},
	})
}