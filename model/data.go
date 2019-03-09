package model

import (
	"context"
	"encoding/json"

	"github.com/tediouscoder/paper-robot/constants"
	ig "github.com/tediouscoder/paper-robot/internal/github"
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

// GetData will get data from data file.
func GetData(ctx context.Context) (sha string, s Storer, err error) {
	sha, content, err := ig.GetFileContent(ctx, constants.DataFilePath)
	if err != nil {
		return
	}

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
		return "", nil, constants.ErrNotImplemented
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

// PutData will put data.json.
func PutData(ctx context.Context, sha string, d Storer) (err error) {
	bs, err := json.MarshalIndent(d, "", "\t")
	if err != nil {
		log.Error("JSON marshal failed", "data", d, "error", err)
		return err
	}

	err = ig.UpdateFileContent(ctx, constants.DataFilePath, sha, string(bs))
	if err != nil {
		return
	}
	return
}

// UpdateData will update data.
func UpdateData(ctx context.Context, fn func(d Storer) (err error)) (err error) {
	sha, data, err := GetData(ctx)
	if err != nil {
		return
	}

	err = fn(data)
	if err != nil {
		return
	}

	err = PutData(ctx, sha, data)
	if err != nil {
		return
	}

	return
}
