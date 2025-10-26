package controller

import (
	"fmt"
	"net/http"
	"strconv"

	pongo3 "github.com/flosch/pongo2/v6"
	"github.com/rs/zerolog"

	"github.com/bruli/waterSystemAdmin/internal/domain/programs"
)

func RemoveTemperatureProgram(set *pongo3.TemplateSet, removeSvc *programs.RemoveTemperature, log zerolog.Logger) http.HandlerFunc {
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

		err = removingTemperatureProgram(r, removeSvc)
		switch {
		case err != nil:
			log.Error().Err(err).Msgf("error removing temperature program. Error: %s", err.Error())
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

func removingTemperatureProgram(r *http.Request, svc *programs.RemoveTemperature) error {
	temp, err := strconv.Atoi(r.PathValue("temperature"))
	if err != nil {
		return fmt.Errorf("invalid temperature: %w", err)
	}
	if err = svc.Remove(r.Context(), temp); err != nil {
		return fmt.Errorf("faild remove temperature program: %w", err)
	}
	return nil
}
