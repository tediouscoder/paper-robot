package paper

import (
	"bytes"
	"strings"
	"testing"
	"text/template"

	"github.com/tediouscoder/paper-robot/model"
)

var content = `{
   "version":1,
   "papers":{
      "Large-scale cluster management at Google with Borg":{
         "url":"https://pdos.csail.mit.edu/6.824/papers/borg.pdf",
         "year":2015,
         "terms":["cluster", "container"],
         "source":"EuroSys"
      }
   }
}`

func TestTemplate(t *testing.T) {
	data, err := model.ParseData(content)
	if err != nil {
		t.Fatal(err)
	}

	funcMap := template.FuncMap{
		// The name "title" is what the function will be called in the template text.
		"title": strings.Title,
	}

	var b bytes.Buffer
	tmpl := template.Must(template.New("readme").Funcs(funcMap).Parse(readme))
	err = tmpl.Execute(&b, data)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(b.String())
}
