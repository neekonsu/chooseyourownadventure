package chooseyourownadventure

import (
	"net/http"
	"text/template"
)

// MapHandler -> take map[string]Arc and handle requests to #/(key) by returning templated HTML for Arc(key)
func MapHandler(storyline map[string]Arc, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// given r.URL.path, write templateHTML(map[r.URL.path] Arc)
		if arc, match := storyline[r.URL.Path[1:]]; match {
			TemplateHTML(arc).Execute(w, arc)
		} else {
			fallback.ServeHTTP(w, r)
		}
	}
}

// TemplateHTML -> return HTML template from single Arc struct
func TemplateHTML(arc Arc) *template.Template {
	tmpl := template.Must(template.ParseFiles("template.html"))
	return tmpl
}
