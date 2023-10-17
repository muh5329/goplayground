package main

import (
	"fmt"
	"net/http"

	"github.com/muh5329/goplayground/rssagg/internal/auth"
	"github.com/muh5329/goplayground/rssagg/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (cfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apikey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			respondWithError(w, 403, fmt.Sprintf("Auth error: %v", err))
			return
		}

		user, err := cfg.DB.GetUserByAPIKey(r.Context(), apikey)
		if err != nil {
			respondWithError(w, 400, fmt.Sprintf("Could not get user: %v", err))
			return
		}
		handler(w, r, user)
	}
}
