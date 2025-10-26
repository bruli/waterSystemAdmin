package controller

import (
	"net/http"
	"strconv"

	pongo3 "github.com/flosch/pongo2/v6"
	"github.com/rs/zerolog"

	"github.com/bruli/waterSystemAdmin/internal/domain/programs"
	"github.com/bruli/waterSystemAdmin/internal/domain/zones"
)

func CreateTemperatureProgram(set *pongo3.TemplateSet, zonesSvc *zones.FindZones, createSvc *programs.CreateTemperature, log zerolog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl, err := set.FromFile("create_temperature_program.html")
		if err != nil {
			log.Error().Err(err).Msgf("error parsing template. Error: %s", err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		context := map[string]interface{}{
			"page": "programs",
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
			processCreateTemperatureForm(r, context, createSvc)
		}

		if err = tpl.ExecuteWriter(context, w); err != nil {
			log.Error().Err(err).Msgf("error executing template. Error: %s", err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func processCreateTemperatureForm(r *http.Request, context map[string]interface{}, svc *programs.CreateTemperature) {
	if err := r.ParseForm(); err != nil {
		context["error_msg"] = "invalid form"
		return
	}
	temp, err := strconv.Atoi(r.FormValue("temperature"))
	if err != nil {
		context["error_msg"] = err.Error()
		return
	}
	hours := r.Form["hours[]"]
	if len(hours) == 0 {
		context["error_msg"] = "hours is required"
		return
	}
	prgms := make([]programs.Program, len(hours))
	for i, hour := range hours {
		execSeconds := r.Form["executions_"+strconv.Itoa(i)+"_seconds[]"]

		if len(execSeconds) == 0 {
			context["error_msg"] = "hours is required"
			return
		}
		h, err := programs.ParseHour(hour)
		if err != nil {
			context["error_msg"] = err.Error()
			return
		}
		exec := make([]programs.Execution, len(execSeconds))
		for n, execSecond := range execSeconds {
			sec, err := strconv.ParseInt(execSecond, 10, 64)
			if err != nil {
				context["error_msg"] = "invalid seconds"
				return
			}
			zones := r.Form["executions_"+strconv.Itoa(i)+"_zones_"+strconv.Itoa(n)+"[]"]
			if len(zones) == 0 {
				context["error_msg"] = "zone is required"
				return
			}
			exec[n] = programs.Execution{
				Seconds: int(sec),
				Zones:   zones,
			}
		}
		prgms[i] = programs.Program{
			Executions: exec,
			Hour:       h,
		}
	}
	temperature := programs.TemperatureProgram{
		Temperature: temp,
		Programs:    prgms,
	}
	if err = svc.Create(r.Context(), &temperature); err != nil {
		context["error_msg"] = err.Error()
		return
	}
	context["success"] = true
}
