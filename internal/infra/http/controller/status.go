package controller

import (
	"net/http"

	pongo3 "github.com/flosch/pongo2/v6"
	"github.com/rs/zerolog"

	"github.com/bruli/waterSystemAdmin/internal/domain/status"
)

func FindStatus(tplSet *pongo3.TemplateSet, svc *status.FindStatus, log zerolog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl, err := tplSet.FromFile("status.html")
		if err != nil {
			log.Error().Err(err).Msgf("error parsing template. Error: %s", err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		context := map[string]interface{}{
			"page": "status",
		}
		st, err := svc.Find(r.Context())
		switch {
		case err != nil:
			log.Error().Err(err).Msgf("error finding status. Error: %s", err.Error())
			context["error_msg"] = err.Error()
		default:
			context["status"] = st
		}
		if err = tpl.ExecuteWriter(context, w); err != nil {
			http.Error(w, "Error executing template", http.StatusInternalServerError)
		}
	}
}
