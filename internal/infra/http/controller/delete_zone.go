package controller

import (
	"net/http"

	"github.com/bruli/waterSystemAdmin/internal/domain/zones"
)

func DeleteZone(svc *zones.Delete) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			id := r.PathValue("id")
			_ = svc.Delete(r.Context(), id)
		}

		http.Redirect(w, r, "/zones", http.StatusSeeOther)
	}
}
