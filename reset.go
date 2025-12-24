package main

import (
	"context"
	"net/http"
)

func (cfg *apiConfig) handlerReset(w http.ResponseWriter, req *http.Request) {
	if cfg.platform != "dev" {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("Access denied"))
	}
	cfg.fileserverHits.Store(0)
	cfg.db.Reset(context.Background())
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hits reset to 0"))
}