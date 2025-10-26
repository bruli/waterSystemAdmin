package controller

import (
	"net/http"

	"github.com/rs/zerolog"

	"github.com/bruli/waterSystemAdmin/internal/domain/status"
)

func Deactivate(svc *status.ActivateDeactivate, log zerolog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			log.Error().Msgf("Method not allowed. Method: %s", r.Method)
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		if err := svc.DeActivate(r.Context()); err != nil {
			log.Error().Err(err).Msgf("error deactivating. Error: %s", err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
