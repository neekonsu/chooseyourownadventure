package chooseyourownadventure

import (
	"log"
	"net/http"
	"text/template"
)

// MapHandler -> take Story and handle requests to #/(key) by returning templated HTML for Chapter(key)
func MapHandler(storyline Story, templateFilename string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// given r.URL.path, write templateHTML(map[r.URL.path] Chapter)
		if chapter, match := storyline[r.URL.Path[1:]]; match {
			TemplateHTML(chapter, templateFilename).Execute(w, chapter)
		} else {
			fallback.ServeHTTP(w, r)
		}
	}
}

// TemplateHTML -> return HTML template from single Chapter struct
func TemplateHTML(chapter Chapter, templateFilename string) *template.Template {
	tmpl := template.Must(template.ParseFiles(templateFilename))
	return tmpl
}

// Check -> generic error handler
func Check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
