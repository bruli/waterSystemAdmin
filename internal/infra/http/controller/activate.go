package controller

import (
	"log/slog"
	"net/http"

	"github.com/bruli/waterSystemAdmin/internal/domain/status"
)

func Activate(svc *status.ActivateDeactivate, log *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			log.ErrorContext(r.Context(), "Method not allowed", slog.String("method", r.Method))
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		if err := svc.Activate(r.Context()); err != nil {
			log.ErrorContext(r.Context(), "error activating", slog.String("error", err.Error()))
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
