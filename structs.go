package chooseyourownadventure

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

// Story simply to designate a type for the whole story
type Story map[string]Chapter
