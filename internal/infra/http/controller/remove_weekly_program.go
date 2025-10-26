package controller

import (
	"fmt"
	"net/http"

	pongo3 "github.com/flosch/pongo2/v6"
	"github.com/rs/zerolog"

	"github.com/bruli/waterSystemAdmin/internal/domain/programs"
)

func RemoveWeeklyProgram(set *pongo3.TemplateSet, removeSvc *programs.RemoveWeekly, log zerolog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			log.Error().Msgf("Method not allowed. Method: %s", r.Method)
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		tpl, err := set.FromFile("redirect_message.html")
		if err != nil {
			log.Error().Err(err).Msgf("error parsing template. Error: %s", err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		context := map[string]interface{}{
			"page": "programs",
		}

		err = removingWeeklyProgram(r, removeSvc)
		switch {
		case err != nil:
			log.Error().Err(err).Msgf("error removing program. Error: %s", err.Error())
			context["error"] = err.Error()
		default:
			context["success"] = true
		}

		if err = tpl.ExecuteWriter(context, w); err != nil {
			log.Error().Err(err).Msgf("error executing template. Error: %s", err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func removingWeeklyProgram(r *http.Request, svc *programs.RemoveWeekly) error {
	day, err := programs.ParseWeekDay(r.PathValue("weekday"))
	if err != nil {
		return fmt.Errorf("invalid day: %w", err)
	}
	if err = svc.Remove(r.Context(), &day); err != nil {
		return fmt.Errorf("faild remove program: %w", err)
	}
	return nil
}
