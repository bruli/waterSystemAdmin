package controller

import (
	"net/http"

	"github.com/bruli/waterSystemAdmin/internal/domain/password"

	"github.com/flosch/pongo2/v6"
	"github.com/gorilla/sessions"
	"github.com/rs/zerolog"
)

func Login(tplSet *pongo2.TemplateSet, store *sessions.CookieStore, check *password.Check, log zerolog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl, err := tplSet.FromFile("login.html")
		if err != nil {
			log.Error().Err(err).Msgf("error parsing template. Error: %s", err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		context := map[string]interface{}{}
		if r.Method == http.MethodPost {
			valid, err := check.Check(r.Context(), r.FormValue("password"))
			if err != nil {
				log.Error().Err(err).Msgf("error checking password. Error: %s", err.Error())
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			switch {
			case valid:
				session, _ := store.Get(r, "session")
				session.Values["authenticated"] = true
				if err = session.Save(r, w); err != nil {
					log.Error().Err(err).Msgf("error saving session. Error: %s", err.Error())
					http.Error(w, err.Error(), http.StatusInternalServerError)
				}
				http.Redirect(w, r, "/status", http.StatusSeeOther)
				return
			default:
				context["error_msg"] = "Invalid password"
			}
		}

		if err = tpl.ExecuteWriter(context, w); err != nil {
			log.Error().Err(err).Msgf("error executing template. Error: %s", err.Error())
			http.Error(w, "Error executant login template ", http.StatusInternalServerError)
		}
	}
}
