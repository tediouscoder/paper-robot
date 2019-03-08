package model

// Paper represents a paper.
type Paper struct {
	// Required fields
	Title string `json:"title,omitempty"`
	URL   string `json:"url"`

	// Optional fields
	Year   int      `json:"year"`
	Terms  []string `json:"terms"`
	Source string   `json:"source"`
}
