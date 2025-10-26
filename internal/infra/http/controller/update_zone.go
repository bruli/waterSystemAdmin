package controller

import (
	"net/http"
	"strconv"

	pongo3 "github.com/flosch/pongo2/v6"
	"github.com/rs/zerolog"

	"github.com/bruli/waterSystemAdmin/internal/domain/zones"
)

func UpdateZone(set *pongo3.TemplateSet, findSvc *zones.FindZones, update *zones.Update, log zerolog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl, err := set.FromFile("update_zone.html")
		if err != nil {
			log.Error().Err(err).Msgf("error parsing template. Error: %s", err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		id := r.PathValue("id")
		context := map[string]interface{}{
			"page": "zones",
		}
		if r.Method == http.MethodPost {
			processUpdateZoneForm(r, id, context, update)
		}
		list, err := findSvc.Find(r.Context())
		if err != nil {
			log.Error().Err(err).Msgf("error finding zones. Error: %s", err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		var zo *zones.Zone
		for _, zon := range list {
			if zon.ID == id {
				zo = &zon
			}
		}
		context["zone"] = zo
		context["relays"] = []int{1, 2, 3, 4}
		if err = tpl.ExecuteWriter(context, w); err != nil {
			log.Error().Err(err).Msgf("error executing template. Error: %s", err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func processUpdateZoneForm(r *http.Request, id string, context map[string]interface{}, update *zones.Update) {
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
	zo := &zones.Zone{
		ID:     id,
		Name:   name,
		Relays: rel,
	}
	if err := update.Update(r.Context(), zo); err != nil {
		context["error_msg"] = "Update zone failed."
		return
	}
	context["success"] = true
}
