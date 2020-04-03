package chooseyourownadventure

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

// MakeStory -> parse JSON file and return map[string]Chapter of ChapterName(key) -> ChapterContents(value)
func MakeStory(r *os.File) Story {
	output := make(Story)
	decoder := json.NewDecoder(r)
	if err := decoder.Decode(&output); err != nil {
		log.Fatal(err)
	}
	return output
}

// PrintMap -> takes Story and returns pretty-printed version of that map
func PrintMap(newMap Story) {
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
