package controller

import (
	"log/slog"
	"net/http"

	"github.com/bruli/waterSystemAdmin/internal/domain/logs"
	"github.com/bruli/waterSystemAdmin/internal/domain/status"
	"github.com/flosch/pongo2/v6"
)

func FindLogs(tplset *pongo2.TemplateSet, svc *logs.FindLogs, stSvc *status.FindStatus, log *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl, err := tplset.FromFile("logs.html")
		if err != nil {
			log.ErrorContext(r.Context(), "error parsing template",
				slog.String("error", err.Error()))
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		tplCtx, failed := buildStatusInTemplateController(r.Context(), stSvc)
		if failed {
			log.ErrorContext(r.Context(), "error building status in template controller")
		}
		tplCtx.Add("page", "logs")
		result, err := svc.Find(r.Context())
		switch {
		case err != nil:
			log.ErrorContext(r.Context(), "error finding logs", slog.String("error", err.Error()))
			tplCtx.AddError(err.Error())
		default:
			tplCtx.Add("logs", result)
		}
		if err = tpl.ExecuteWriter(tplCtx.toPongoContext(), w); err != nil {
			log.ErrorContext(r.Context(), "error executing template", slog.String("error", err.Error()))
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
