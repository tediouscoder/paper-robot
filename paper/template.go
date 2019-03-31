package paper

import (
	"bytes"
	"strings"
	"text/template"

	"github.com/tediouscoder/paper-robot/model"
)

var funcMap = template.FuncMap{
	// The name "title" is what the function will be called in the template text.
	"title": strings.Title,
}

const readme = `## Papers

| Title | Source | Terms |
|:--------|:--------|:--------|
{{- range $k, $v := .Papers}}
|[{{$k}}]({{$v.URL}})|{{$v.Source}}|
{{- range $_, $term := $v.Terms -}}
{{- if ne $term "" }}
{{- printf "#%s " $term | title}}
{{- end }}
{{- end -}}
|
{{- end}}
`

const filteredContent = `## Papers

| Title | Source | Terms |
|:--------|:--------|:--------|
{{- range $k, $v := .Papers}}
|[{{$k}}]({{$v.URL}})|{{$v.Source}}|
{{- range $_, $term := $v.Terms -}}
{{- if ne $term "" }}
{{- printf "#%s " $term | title}}
{{- end }}
{{- end -}}
|
{{- end}}
`

// GenerateREADME will generate README.md
func GenerateREADME(data model.Storer) (s string, err error) {
	var b bytes.Buffer
	t := template.Must(template.New("readme").Funcs(funcMap).Parse(readme))
	err = t.Execute(&b, data)
	if err != nil {
		return
	}
	return b.String(), nil
}

// GenerateFilteredContent will generate FilteredContent
func GenerateFilteredContent(data model.Storer, filter model.Filter) (s string, err error) {
	var b bytes.Buffer

	papers, err := data.FilterPaper(filter)
	if err != nil {
		return
	}

	t := template.Must(template.New("readme").Funcs(funcMap).Parse(readme))
	err = t.Execute(&b, struct {
		Papers []*model.Paper
	}{
		Papers: papers,
	})
	if err != nil {
		return
	}
	return b.String(), nil
}
