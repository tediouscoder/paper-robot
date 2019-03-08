package model

import (
	"encoding/json"

	"github.com/tediouscoder/paper-robot/internal/log"
)

type v1Data struct {
	Version int `json:"version"`

	UpdatedAt int64             `json:"updated_at"`
	Papers    map[string]*Paper `json:"papers"`
}

func newV1Data(c string) (d *v1Data, err error) {
	d = &v1Data{}
	err = json.Unmarshal([]byte(c), d)
	if err != nil {
		log.Error("JSON unmarshal failed", "data", d, "error", err)
		return
	}

	// We should always set v1Data's version to 1.
	d.Version = 1
	if d.Papers == nil {
		d.Papers = make(map[string]*Paper)
	}
	return
}

// AddPaper implements Storer.AddPaper
func (d *v1Data) AddPaper(p *Paper) error {
	d.Papers[p.Title] = p
	// Set title to "" to reduce data.json's size.
	d.Papers[p.Title].Title = ""

	return nil
}

// UpdatePaper implements Storer.UpdatePaper
func (d *v1Data) UpdatePaper(p *Paper) error {
	return d.AddPaper(p)
}

// RemovePaper implements Storer.RemovePaper
func (d *v1Data) RemovePaper(p *Paper) error {
	delete(d.Papers, p.Title)

	return nil
}

// Migrate implements Migrate.Migrate
func (d *v1Data) Migrate() (Storer, error) {
	return d, nil
}
