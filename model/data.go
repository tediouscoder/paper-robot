package model

import (
	"encoding/json"

	"github.com/tediouscoder/paper-robot/constants"
	"github.com/tediouscoder/paper-robot/internal/log"
)

const currentVersion = 1

// Storer will handle all data operations.
type Storer interface {
	AddPaper(*Paper) error
	UpdatePaper(*Paper) error
	RemovePaper(*Paper) error

	Migrate() (Storer, error)
}

// version is the version of current storer.
type version struct {
	Version int `json:"version"`
}

// ParseData will parse content into data.
func ParseData(content string) (s Storer, err error) {
	ver := &version{}
	err = json.Unmarshal([]byte(content), ver)
	if err != nil {
		log.Error("JSON unmarshal failed", "version", ver, "error", err)
		return
	}

	switch ver.Version {
	case 0:
		s, err = newV0Data(content)
	case 1:
		s, err = newV1Data(content)
	default:
		return nil, constants.ErrNotImplemented
	}
	if err != nil {
		return
	}

	s, err = s.Migrate()
	if err != nil {
		return
	}
	return
}

// FormatData will format data into json.
func FormatData(d Storer) (content string, err error) {
	bs, err := json.MarshalIndent(d, "", "\t")
	if err != nil {
		log.Error("JSON marshal failed", "data", d, "error", err)
		return "", err
	}
	return string(bs), nil
}
