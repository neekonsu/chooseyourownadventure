package chooseyourownadventure

import (
	"encoding/json"
	"fmt"
	"log"
)

// JSON schema to work with:
/*
{
  // Each story arc will have a unique key that represents
  // the name of that particular arc.
  "story-arc": {
    "title": "A title for that story arc. Think of it like a chapter title.",
    "story": [
      "A series of paragraphs, each represented as a string in a slice.",
      "This is a new paragraph in this particular story arc."
    ],
    // Options will be empty if it is the end of that
    // particular story arc. Otherwise it will have one or
    // more JSON objects that represent an "option" that the
    // reader has at the end of a story arc.
    "options": [
      {
        "text": "the text to render for this option. eg 'venture down the dark passage'",
        "arc": "the name of the story arc to navigate to. This will match the story-arc key at the very root of the JSON document"
      }
    ]
  },
  ...
}
*/

// Option represents a single element of the "options" array to be expected from the JSON file
type Option struct {
	Text  string `json:"text"`
	ArcID string `json:"arc"`
}

// Arc represents a single story Arc/Chapter
type Arc struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []Option `json:"options"`
}

// MakeMap -> parse JSON file and return map[string]Arc of ArcName(key) -> ArcContents(value)
func MakeMap(data []byte) map[string]Arc {
	output := make(map[string]Arc)
	if err := json.Unmarshal(data, &output); err != nil {
		log.Fatal(err)
	}
	return output
}

// PrintMap -> takes map[string]Arc and returns pretty-printed version of that map
func PrintMap(newMap map[string]Arc) {
	for k, v := range newMap {
		fmt.Println(k, ": ")
		for _, line := range ListArc(v) {
			fmt.Println("\t", line)
		}
	}
}

// ListArc -> pretty-print individual arc
func ListArc(arc Arc) []string {
	output := []string{
		"Title: " + arc.Title,
		"Story: ",
	}
	for _, e := range arc.Story {
		output = append(output, "\t"+e)
	}
	output = append(output, "Options: ")
	for _, e := range arc.Options {
		output = append(output, "\t"+e.Text)
		output = append(output, "\t"+e.ArcID)
	}
	return output
}
