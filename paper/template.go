package paper

const readme = `## Papers

| Title | Source | Terms |
|:--------|:--------|:--------|
{{- range $_, $v := .Papers}}
|[{{$v.Title}}]({{$v.URL}})|{{$v.Source}}|{{$v.Terms}}|
{{- end}}
`
