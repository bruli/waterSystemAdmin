package controller

import (
	"net/http"

	"github.com/bruli/waterSystemAdmin/internal/domain/status"
	pongo3 "github.com/flosch/pongo2/v6"
	"github.com/rs/zerolog"

	"github.com/bruli/waterSystemAdmin/internal/domain/zones"
)

func Zones(set *pongo3.TemplateSet, svc *zones.FindZones, stSvc *status.FindStatus, log zerolog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl, err := set.FromFile("zones.html")
		if err != nil {
			log.Error().Err(err).Msgf("error parsing template. Error: %s", err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		tplCtx, err := buildStatusInTemplateController(r.Context(), stSvc)
		if err != nil {
			log.Error().Err(err).Msgf("error building status in template controller. Error: %s", err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		tplCtx.Add("page", "zones")
		list, err := svc.Find(r.Context())
		switch {
		case err != nil:
			log.Error().Err(err).Msgf("error finding zones. Error: %s", err.Error())
			tplCtx.AddError(err.Error())
		default:
			tplCtx.Add("zones", list)
		}
		if err = tpl.ExecuteWriter(tplCtx.toPongoContext(), w); err != nil {
			log.Error().Err(err).Msgf("error executing template. Error: %s", err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
