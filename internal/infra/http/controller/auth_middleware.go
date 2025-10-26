package controller

import (
	"net/http"

	"github.com/bruli/waterSystemAdmin/internal/domain/password"

	"github.com/gorilla/sessions"
	"github.com/rs/zerolog"
)

type MiddlewareFunc func(next http.HandlerFunc) http.HandlerFunc

func AuthMiddleware(store *sessions.CookieStore, exists *password.Exists, log zerolog.Logger) MiddlewareFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			passwdExists, err := exists.Exists(r.Context())
			if err != nil {
				log.Error().Err(err).Msgf("error checking if password exists. Error: %s", err.Error())
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			if !passwdExists {
				http.Redirect(w, r, "/password", http.StatusSeeOther)
				return
			}
			session, err := store.Get(r, "session")
			if err != nil {
				log.Error().Err(err).Msgf("error getting session. Error: %s", err.Error())
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			if session == nil {
				http.Redirect(w, r, "/login", http.StatusSeeOther)
				return
			}
			authenticated, ok := session.Values["authenticated"]
			if !ok {
				http.Redirect(w, r, "/login", http.StatusSeeOther)
				return
			}
			auth, _ := authenticated.(bool)
			if auth {
				next(w, r)
				return
			}
			http.Redirect(w, r, "/login", http.StatusSeeOther)
		}
	}
}
