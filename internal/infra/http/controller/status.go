package controller

import (
	"context"
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
		tplCtx, err := buildStatusInTemplateController(r.Context(), svc)
		if err != nil {
			log.Error().Err(err).Msgf("error building status in template controller. Error: %s", err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		tplCtx.Add("page", "status")
		if err = tpl.ExecuteWriter(tplCtx.toPongoContext(), w); err != nil {
			http.Error(w, "Error executing template", http.StatusInternalServerError)
		}
	}
}

func buildStatusInTemplateController(ctx context.Context, svc *status.FindStatus) (TemplateController, error) {
	tplCtx := newTemplateController()
	st, err := svc.Find(ctx)
	if err != nil {
		return nil, err
	}
	tplCtx.Add("status", st)
	return tplCtx, nil
}
