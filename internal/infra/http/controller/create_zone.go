package controller

import (
	"net/http"
	"strconv"

	pongo3 "github.com/flosch/pongo2/v6"
	"github.com/rs/zerolog"

	"github.com/bruli/waterSystemAdmin/internal/domain/zones"
)

func CreateZone(set *pongo3.TemplateSet, svc *zones.Create, log zerolog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl, err := set.FromFile("create_zone.html")
		if err != nil {
			log.Error().Err(err).Msgf("error parsing template. Error: %s", err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		context := map[string]interface{}{
			"page":   "zones",
			"relays": []int{1, 2, 3, 4},
		}
		if r.Method == http.MethodPost {
			processCreateZoneForm(r, context, svc)
		}

		if err = tpl.ExecuteWriter(context, w); err != nil {
			log.Error().Err(err).Msgf("error executing template. Error: %s", err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func processCreateZoneForm(r *http.Request, context map[string]interface{}, svc *zones.Create) {
	id := r.FormValue("id")
	name := r.FormValue("name")
	relayValues := r.Form["relays"]

	if id == "" || name == "" || len(relayValues) == 0 {
		context["error_msg"] = "All fields are required"
		return
	}
	rel := make([]int, len(relayValues))
	for i, val := range relayValues {
		num, err := strconv.Atoi(val)
		if err != nil {
			context["error_msg"] = "Invalid relay value"
			return
		}
		rel[i] = num
	}
	if err := svc.Create(r.Context(), &zones.Zone{
		ID:     id,
		Name:   name,
		Relays: rel,
	}); err != nil {
		context["error_msg"] = "Create zone failed."
		return
	}

	context["success"] = true
}
