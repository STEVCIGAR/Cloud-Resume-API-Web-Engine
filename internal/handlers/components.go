package handlers

import (
	"net/http"

	"github.com/cloud-resume/api/internal/data"
	"github.com/cloud-resume/api/internal/views/components"
)

// SkillsComponentHandler serves the partial HTML fragment of filtered skills for HTMX swaps.
func SkillsComponentHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cat := r.URL.Query().Get("cat")
		skills := data.DefaultResume.GetSkillsByCategory(cat)

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		component := components.SkillsGrid(skills)
		if err := component.Render(r.Context(), w); err != nil {
			http.Error(w, "Failed to render skills component", http.StatusInternalServerError)
		}
	}
}
