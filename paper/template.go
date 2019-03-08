package paper

const readme = `## Papers

| Title | Source | Terms |
|:--------|:--------|:--------|
{{- range $k, $v := .Papers}}
|[{{$k}}]({{$v.URL}})|{{$v.Source}}|
{{- range $_, $term := $v.Terms -}}
{{printf "#%s " $term | title}}
{{- end -}}
|
{{- end}}
`
