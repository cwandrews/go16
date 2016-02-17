// You have to run this with a go 1.6 compiler as is. Replace '-}}'
// and '{{-' with '}}' and '{{' respectively to get it to compile with
// a 1.5 compiler, but the whitespace will be different.
package main

import (
	"os"
	"text/template"
)

const tmpl = `
{{ range . -}}
{{ . }}
{{- end }}`

func main() {
	t := template.Must(template.New("").Parse(tmpl))
	t.Execute(os.Stdout, []string{"mars", "mercury", "pluto", "neptune"})
}
