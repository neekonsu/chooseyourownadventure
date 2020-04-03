package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/neekonsu/chooseyourownadventure"
)

func main() {
	filenameJSON := flag.String("story", "./res/story.json", "filename for the story json file")
	filenameHTML := flag.String("template", "./res/template.html", "filename for the template html file")
	flag.Parse()
	JSONreader, err := os.Open(*filenameJSON)
	chooseyourownadventure.Check(err)
	mux := defaultMux()
	story := chooseyourownadventure.MakeStory(JSONreader)
	chooseyourownadventure.PrintMap(story)
	mapHandler := chooseyourownadventure.MapHandler(story, *filenameHTML, mux)
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", mapHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/intro", http.StatusFound)
	})
	return mux
}
