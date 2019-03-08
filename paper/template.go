package paper

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
