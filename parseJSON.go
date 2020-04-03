package chooseyourownadventure

import (
	"encoding/json"
	"fmt"
	"log"
)

// Option represents a single element of the "options" array to be expected from the JSON file
type Option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

// Chapter represents a single story Arc/Chapter
type Chapter struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []Option `json:"options"`
}

// MakeMap -> parse JSON file and return map[string]Chapter of ChapterName(key) -> ChapterContents(value)
func MakeMap(data []byte) map[string]Chapter {
	output := make(map[string]Chapter)
	if err := json.Unmarshal(data, &output); err != nil {
		log.Fatal(err)
	}
	return output
}

// PrintMap -> takes map[string]Chapter and returns pretty-printed version of that map
func PrintMap(newMap map[string]Chapter) {
	for k, v := range newMap {
		fmt.Println(k, ": ")
		for _, line := range ListChapter(v) {
			fmt.Println("\t", line)
		}
	}
}

// ListChapter -> pretty-print individual Chapter
func ListChapter(chapter Chapter) []string {
	output := []string{
		"Title: " + chapter.Title,
		"Story: ",
	}
	for _, e := range chapter.Story {
		output = append(output, "\t"+e)
	}
	output = append(output, "Options: ")
	for _, e := range chapter.Options {
		output = append(output, "\t"+e.Text)
		output = append(output, "\t"+e.Arc)
	}
	return output
}
