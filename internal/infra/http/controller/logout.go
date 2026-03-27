package controller

import (
	"log/slog"
	"net/http"

	"github.com/gorilla/sessions"
)

func Logout(store *sessions.CookieStore, log *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := store.Get(r, "session")
		session.Values["authenticated"] = false
		if err := session.Save(r, w); err != nil {
			log.ErrorContext(r.Context(), "error saving session",
				slog.String("error", err.Error()))
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
