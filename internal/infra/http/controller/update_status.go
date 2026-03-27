package controller

import (
	"log/slog"
	"net/http"

	"github.com/bruli/waterSystemAdmin/internal/domain/status"
)

func UpdateStatus(svc *status.UpdateStatus, log *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := svc.Update(r.Context()); err != nil {
			log.ErrorContext(r.Context(), "error updating status", slog.String("error", err.Error()))
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
