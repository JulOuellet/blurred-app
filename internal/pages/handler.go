package pages

import (
	"html/template"
	"net/http"
	"path/filepath"
)

var templates = template.Must(
	template.ParseGlob(filepath.Join("views", "*.html")),
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "index", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
