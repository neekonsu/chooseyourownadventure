package main

import (
	"fmt"
	"net/http"

	"github.com/neekonsu/chooseyourownadventure"
)

func main() {
	mux := defaultMux()
	JSON := []byte(`{
	"story-arc": {
		"title": "A title for that story arc. Think of it like a chapter title.",
		"story": [
			"A series of paragraphs, each represented as a string in a slice.",
			"This is a new paragraph in this particular story arc."
		],
		"options": [
			{
				"text": "the text to render for this option. eg 'venture down the dark passage'",
				"arc": "the name of the story arc to navigate to. This will match the story-arc key at the very root of the JSON document"
			},
			{
				"text": "the text to render for this option. eg 'venture down the dark passage'",
				"arc": "the name of the story arc to navigate to. This will match the story-arc key at the very root of the JSON document"
			}
		]
	},
	"story-arc2": {
		"title": "A title 2for that story arc. Think of it like a chapter title.",
		"story": [
			"A seri2es of paragraphs, each represented as a string in a slice.",
			"This is a new p2aragraph in this particular story arc."
		],
		"options": [
			{
				"text": "the text2 to render for this option. eg 'venture down the dark passage'",
				"arc": "the name o2f the story arc to navigate to. This will match the story-arc key at the very root of the JSON document"
			},
			{
				"text": "the text2 to render for this option. eg 'venture down the dark passage'",
				"arc": "the name o2f the story arc to navigate to. This will match the story-arc key at the very root of the JSON document"
			}
		]
	}
}`)
	newMap := chooseyourownadventure.MakeMap(JSON)
	chooseyourownadventure.PrintMap(newMap)
	mapHandler := chooseyourownadventure.MapHandler(newMap, mux)
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", mapHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/intro", http.StatusOK)
	})
	return mux
}
