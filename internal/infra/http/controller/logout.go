package controller

import (
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/rs/zerolog"
)

func Logout(store *sessions.CookieStore, log zerolog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := store.Get(r, "session")
		session.Values["authenticated"] = false
		if err := session.Save(r, w); err != nil {
			log.Error().Err(err).Msgf("error saving session. Error: %s", err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
