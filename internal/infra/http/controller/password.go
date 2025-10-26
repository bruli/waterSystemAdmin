package controller

import (
	"net/http"

	"github.com/bruli/waterSystemAdmin/internal/domain/password"
	pongo3 "github.com/flosch/pongo2/v6"
	"github.com/rs/zerolog"
)

func Password(set *pongo3.TemplateSet, create *password.Create, log zerolog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl, err := set.FromFile("password.html")
		if err != nil {
			log.Error().Err(err).Msgf("error parsing template. Error: %s", err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		context := map[string]interface{}{}
		if r.Method == http.MethodPost {
			passw := r.FormValue("password")
			rePassword := r.FormValue("re-password")
			switch {
			case passw != rePassword:
				context["error_msg"] = "Passwords do not match"
			default:
				if err = create.Create(r.Context(), passw); err != nil {
					log.Error().Err(err).Msgf("error creating password. Error: %s", err.Error())
					http.Error(w, err.Error(), http.StatusInternalServerError)
				}
				http.Redirect(w, r, "/", http.StatusSeeOther)
				return
			}

		}

		if err = tpl.ExecuteWriter(context, w); err != nil {
			log.Error().Err(err).Msgf("error executing template. Error: %s", err.Error())
		}
	}
}
