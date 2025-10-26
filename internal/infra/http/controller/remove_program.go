package controller

import (
	"fmt"
	"net/http"

	pongo3 "github.com/flosch/pongo2/v6"
	"github.com/rs/zerolog"

	"github.com/bruli/waterSystemAdmin/internal/domain/programs"
)

func RemoveProgram(set *pongo3.TemplateSet, svc *programs.Remove, log zerolog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			log.Error().Msg("Method not allowed. Method: " + r.Method + "")
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

		err = removingProgram(r, svc)
		switch {
		case err != nil:
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

func removingProgram(r *http.Request, svc *programs.Remove) error {
	hour, err := programs.ParseHour(r.PathValue("hour"))
	if err != nil {
		return fmt.Errorf("invalid hour: %w", err)
	}
	programType, err := programs.ParseProgramType(r.PathValue("type"))
	if err != nil {
		return fmt.Errorf("invalid type: %w", err)
	}

	err = svc.Remove(r.Context(), &hour, programType)
	if err != nil {
		return fmt.Errorf("faild remove program: %w", err)
	}
	return nil
}
