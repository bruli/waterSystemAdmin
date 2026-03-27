package controller

import (
	"net/http"

	"github.com/bruli/waterSystemAdmin/internal/domain/status"
	"github.com/flosch/pongo2/v6"
	"github.com/rs/zerolog"

	"github.com/bruli/waterSystemAdmin/internal/domain/logs"
)

func FindLogs(tplset *pongo2.TemplateSet, svc *logs.FindLogs, stSvc *status.FindStatus, log zerolog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl, err := tplset.FromFile("logs.html")
		if err != nil {
			log.Error().Err(err).Msgf("error parsing template. Error: %s", err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		tplCtx, failed := buildStatusInTemplateController(r.Context(), stSvc)
		if failed {
			log.Error().Msg("error building status in template controller")
		}
		tplCtx.Add("page", "logs")
		result, err := svc.Find(r.Context())
		switch {
		case err != nil:
			log.Error().Err(err).Msgf("error finding logs. Error: %s", err.Error())
			tplCtx.AddError(err.Error())
		default:
			tplCtx.Add("logs", result)
		}
		if err = tpl.ExecuteWriter(tplCtx.toPongoContext(), w); err != nil {
			log.Error().Err(err).Msgf("error executing template. Error: %s", err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
