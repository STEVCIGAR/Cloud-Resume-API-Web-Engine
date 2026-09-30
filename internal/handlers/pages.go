package handlers

import (
	"net/http"

	"github.com/cloud-resume/api/internal/data"
	"github.com/cloud-resume/api/internal/views"
)

// HomeHandler renders the complete dashboard page using Templ.
func HomeHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		component := views.Home(data.DefaultResume, "all")
		if err := component.Render(r.Context(), w); err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	}
}
