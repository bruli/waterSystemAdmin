package controller

import (
	"net/http"

	pongo3 "github.com/flosch/pongo2/v6"
	"github.com/rs/zerolog"

	"github.com/bruli/waterSystemAdmin/internal/domain/programs"
)

type Section struct {
	ID    string
	Title string
	Data  []programs.Program
}

type Weekly struct {
	WeekDay  string
	Programs []programs.Program
}

func Programs(set *pongo3.TemplateSet, svc *programs.FindAllPrograms, log zerolog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl, err := set.FromFile("programs.html")
		if err != nil {
			log.Error().Err(err).Msgf("error parsing template. Error: %s", err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		context := map[string]interface{}{
			"page": "programs",
		}
		progrms, err := svc.Find(r.Context())
		switch {
		case err != nil:
			log.Error().Err(err).Msgf("error finding programs. Error: %s", err.Error())
			context["error_msg"] = err.Error()
		default:
			sections := []Section{
				{ID: "collapseOne", Title: "Daily", Data: progrms.Daily},
				{ID: "collapseTwo", Title: "Odd", Data: progrms.Odd},
				{ID: "collapseThree", Title: "Even", Data: progrms.Even},
			}
			context["sections"] = sections
			context["weekly"] = buildWeekly(progrms.Weekly)
			context["temperature"] = progrms.Temperature
		}
		if err = tpl.ExecuteWriter(context, w); err != nil {
			log.Error().Err(err).Msgf("error executing template. Error: %s", err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func buildWeekly(progrms []programs.Weekly) []Weekly {
	weekly := make([]Weekly, len(progrms))
	for i, week := range progrms {
		weekly[i] = Weekly{
			WeekDay:  week.WeekDay.String(),
			Programs: week.Programs,
		}
	}
	return weekly
}
