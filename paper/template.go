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
