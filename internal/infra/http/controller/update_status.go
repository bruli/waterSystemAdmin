package controller

import (
	"net/http"

	"github.com/rs/zerolog"

	"github.com/bruli/waterSystemAdmin/internal/domain/status"
)

func UpdateStatus(svc *status.UpdateStatus, log zerolog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := svc.Update(r.Context()); err != nil {
			log.Error().Err(err).Msgf("error updating status. Error: %s", err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
