package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/bruli/waterSystemAdmin/internal/domain/status"
	pongo3 "github.com/flosch/pongo2/v6"
	"github.com/rs/zerolog"

	"github.com/bruli/waterSystemAdmin/internal/domain/execution"
)

func Execution(set *pongo3.TemplateSet, svc *execution.ExecuteZone, stSvc *status.FindStatus, log zerolog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		tpl, err := set.FromFile("execution.html")
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
		tplCtx.Add("page", "execution")
		tplCtx.Add("id", id)
		if r.Method == http.MethodPost {
			seconds, err := strconv.Atoi(r.FormValue("seconds"))
			if err != nil {
				log.Error().Err(err).Msgf("error parsing seconds. Error: %s", err.Error())
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			if err = svc.Execute(r.Context(), &execution.Execution{
				ID:      id,
				Seconds: seconds,
			}); err != nil {
				tplCtx.AddError(fmt.Sprintf("Error executing zone: %s", err.Error()))
				if err = tpl.ExecuteWriter(tplCtx.toPongoContext(), w); err != nil {
					log.Error().Err(err).Msgf("error executing template. Error: %s", err.Error())
					http.Error(w, err.Error(), http.StatusInternalServerError)
				}
				return
			}
			tplCtx.Add("success", true)
			if err = tpl.ExecuteWriter(tplCtx.toPongoContext(), w); err != nil {
				log.Error().Err(err).Msgf("error executing template. Error: %s", err.Error())
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}
		if err = tpl.ExecuteWriter(tplCtx.toPongoContext(), w); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
