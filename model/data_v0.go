package model

import (
	"encoding/json"

	"github.com/tediouscoder/paper-robot/constants"
	"github.com/tediouscoder/paper-robot/internal/log"
)

type v0Data struct {
	Version int `json:"version"`

	UpdatedAt int64    `json:"updated_at"`
	Papers    []*Paper `json:"papers"`
}

func newV0Data(c string) (d *v0Data, err error) {
	d = &v0Data{}
	err = json.Unmarshal([]byte(c), d)
	if err != nil {
		log.Error("JSON unmarshal failed", "data", d, "error", err)
		return
	}

	// We should always set v0Data's version to 0.
	d.Version = 0
	return
}

// AddPaper implements Storer.AddPaper
func (d *v0Data) AddPaper(p *Paper) error {
	d.Papers = append(d.Papers, p)

	return nil
}

// UpdatePaper implements Storer.UpdatePaper
func (d *v0Data) UpdatePaper(p *Paper) error {
	return constants.ErrNotImplemented
}

// RemovePaper implements Storer.RemovePaper
func (d *v0Data) RemovePaper(p *Paper) error {
	return constants.ErrNotImplemented
}

// Migrate implements Migrate.Migrate
func (d *v0Data) Migrate() (s Storer, err error) {
	nd, err := newV1Data("{}")
	if err != nil {
		return
	}

	nd.Version = currentVersion
	nd.UpdatedAt = d.UpdatedAt

	for _, v := range d.Papers {
		p := *v
		err = nd.AddPaper(&p)
		if err != nil {
			return
		}
	}
	return nd, nil
}
