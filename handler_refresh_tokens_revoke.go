package main

import (
	"context"
	"net/http"

	"github.com/kenzierivan/chirpy/internal/auth"
)

func (cfg *apiConfig) handlerRevoke(w http.ResponseWriter, req *http.Request) {
	refreshToken, err := auth.GetBearerToken(req.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Couldn't find token", err)
		return
	}
	_, err = cfg.db.RevokeToken(context.Background(), refreshToken)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Couldn't revoke session", err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}