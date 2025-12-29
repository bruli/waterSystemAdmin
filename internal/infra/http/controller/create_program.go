package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/bruli/waterSystemAdmin/internal/domain/status"
	"github.com/flosch/pongo2/v6"
	"github.com/rs/zerolog"

	"github.com/bruli/waterSystemAdmin/internal/domain/programs"
	"github.com/bruli/waterSystemAdmin/internal/domain/zones"
)

func CreateProgram(tplSet *pongo2.TemplateSet, zonesSvc *zones.FindZones, createSvc *programs.Create, stSvc *status.FindStatus, log zerolog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl, err := tplSet.FromFile("create_program.html")
		if err != nil {
			log.Error().Err(err).Msgf("error parsing template. Error: %s", err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		programType := r.URL.Query().Get("type")
		context := map[string]interface{}{
			"page": "programs",
			"type": programType,
		}
		st, err := stSvc.Find(r.Context())
		switch {
		case err != nil:
			log.Error().Err(err).Msgf("error finding status. Error: %s", err.Error())
			context["error_msg"] = err.Error()
		default:
			context["status"] = st
		}
		zones, err := zonesSvc.Find(r.Context())
		if err != nil {
			log.Error().Err(err).Msgf("error finding zones. Error: %s", err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		switch {
		case len(zones) == 0:
			context["error_msg"] = "No zones found"
		default:
			context["zones"] = zones
		}

		if r.Method == http.MethodPost {
			processCreateProgramForm(r, context, createSvc, programType)
		}
		if err = tpl.ExecuteWriter(context, w); err != nil {
			log.Error().Err(err).Msgf("error executing template. Error: %s", err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func processCreateProgramForm(r *http.Request, context map[string]interface{}, svc *programs.Create, programType string) {
	if err := r.ParseForm(); err != nil {
		context["error_msg"] = "Error parsing form"
		return
	}

	hour, err := programs.ParseHour(r.FormValue("hour"))
	if err != nil {
		context["error_msg"] = "Invalid hour"
		return
	}
	seconds := r.Form["seconds[]"]

	executions := make([]programs.Execution, 0)

	for i, sec := range seconds {
		zonesKey := fmt.Sprintf("zones_%d[]", i)
		zones := r.Form[zonesKey]

		if len(zones) == 0 {
			context["error_msg"] = "Invalid zones"
			return
		}

		s, err := strconv.Atoi(sec)
		if err != nil {
			context["error_msg"] = "Invalid seconds"
			return
		}

		executions = append(executions, programs.Execution{
			Seconds: s,
			Zones:   zones,
		})
	}
	pr := programs.Program{
		Executions: executions,
		Hour:       hour,
	}
	progType, err := programs.ParseProgramType(programType)
	if err != nil {
		context["error_msg"] = "Invalid program type"
		return
	}
	if err = svc.Create(r.Context(), &pr, progType); err != nil {
		context["error_msg"] = "Error creating program"
		return
	}
	context["success"] = true
}
