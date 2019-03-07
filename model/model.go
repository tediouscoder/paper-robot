package model

import (
	"encoding/json"

	"github.com/tediouscoder/paper-robot/internal/log"
)

// Data is the data for paper robot.
type Data struct {
	UpdatedAt int64    `json:"updated_at"`
	Papers    []*Paper `json:"papers"`
}

// ParseData will parse content into data.
func ParseData(content string) (d *Data, err error) {
	d = &Data{}
	log.Debug("Current content", "content", content)
	err = json.Unmarshal([]byte(content), d)
	if err != nil {
		log.Error("JSON unmarshal failed", "data", d, "error", err)
		return
	}
	return
}

// FormatData will format data into json.
func FormatData(d *Data) (content string, err error) {
	bs, err := json.Marshal(d)
	if err != nil {
		log.Error("JSON marshal failed", "data", d, "error", err)
		return "", err
	}
	return string(bs), nil
}

// Paper represents a paper.
type Paper struct {
	// Required fields
	Title string `json:"title"`
	URL   string `json:"url"`

	// Optional fields
	Year   int      `json:"year"`
	Month  int      `json:"month"`
	Terms  []string `json:"terms"`
	Source string   `json:"source"`
}
