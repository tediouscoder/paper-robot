package model

import (
	"bufio"
	"context"
	"strconv"
	"strings"

	"github.com/tediouscoder/paper-robot/constants"
	ig "github.com/tediouscoder/paper-robot/internal/github"
	"github.com/tediouscoder/paper-robot/internal/http"
	"github.com/tediouscoder/paper-robot/internal/log"
)

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

// ParsePaper will parse a paper form current issue's content.
func ParsePaper(ctx context.Context, titleOnly bool) (p *Paper, err error) {
	// Parse issue
	issueContent, err := ig.GetIssueContent(ctx)
	if err != nil {
		return
	}

	p, err = parsePaperFromIssueContent(issueContent, titleOnly)
	if err != nil {
		return
	}
	return
}

// parsePaperFromIssueContent will parse content a paper.
func parsePaperFromIssueContent(content string, titleOnly bool) (id *Paper, err error) {
	s := bufio.NewScanner(strings.NewReader(content))

	m := make(map[string]string)
	for s.Scan() {
		text := s.Text()

		if strings.HasPrefix(text, ">") {
			continue
		}

		line := strings.SplitN(text, ":", 2)
		if len(line) != 2 {
			continue
		}

		m[strings.ToLower(line[0])] = strings.TrimSpace(line[1])
	}

	id = &Paper{}

	id.Title = m["title"]
	if id.Title == "" {
		return nil, constants.ErrRequiredFiledMissing
	}

	// Return directly because we only need title.
	if titleOnly {
		return
	}

	id.URL = m["url"]
	if id.URL == "" {
		return nil, constants.ErrRequiredFiledMissing
	}
	resp, err := http.Head(id.URL)
	if err != nil {
		log.Error("URL failed to head", "url", id.URL, "error", err)
		return nil, constants.ErrInvalidPaperURL
	}
	if resp.StatusCode != 200 || resp.Header.Get("content-type") != "application/pdf" {
		log.Error("URL is not a valid pdf url", "url", id.URL)
		return nil, constants.ErrInvalidPaperURL
	}

	id.Source = m["source"]

	terms := strings.Split(strings.ToLower(m["terms"]), ",")
	for _, v := range terms {
		id.Terms = append(id.Terms, strings.TrimSpace(v))
	}

	if m["year"] != "" {
		n, err := strconv.ParseInt(m["year"], 10, 64)
		if err != nil {
			return nil, err
		}
		id.Year = int(n)
	}
	return
}
