package utils

import (
	"bufio"
	"github.com/tediouscoder/paper-robot/constants"
	"github.com/tediouscoder/paper-robot/internal/log"
	"github.com/tediouscoder/paper-robot/model"
	"net/http"
	"strconv"
	"strings"
)

// ParsePaper will parse a paper.
func ParsePaper(content string) (id *model.Paper, err error) {
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

	id = &model.Paper{}

	id.Title = m["title"]
	if id.Title == "" {
		return nil, constants.ErrRequiredFiledMissing
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
	id.Terms = strings.Split(strings.ToLower(m["terms"]), ",")

	if m["year"] != "" {
		n, err := strconv.ParseInt(m["year"], 10, 64)
		if err != nil {
			return nil, err
		}
		id.Year = int(n)
	}

	if m["month"] != "" {
		n, err := strconv.ParseInt(m["month"], 10, 64)
		if err != nil {
			return nil, err
		}
		id.Month = int(n)
	}
	return
}
