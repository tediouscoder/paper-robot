package paper

const readme = `## Papers

| Title | Source | Terms | URL |
|:--------|:--------|:--------|:--------|
{{- range $_, $v := .Papers}}
|{{$v.Title}}|{{$v.Source}}|{{$v.Terms}}|{{$v.URL}}|
{{- end}}
`
