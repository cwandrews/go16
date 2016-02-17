// This has to be compiled with a go 1.6 compiler
package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

const (
	// Define a root template where a block gets declared called
	// list
	root_tmpl = `Names:{{block "list" .}}{{"\n"}}{{range .}}{{println "-" .}}{{end}}{{end}}`
	// Define an overlay template to that redefines the block
	overlay_tmpl = `{{define "list"}} {{join . ", "}}{{end}}`
)

var (
	funcs   = template.FuncMap{"join": strings.Join}
	planets = []string{"mars", "mercury", "pluto", "neptune"}
)

func main() {
	rootTmpl, err := template.New("root").Funcs(funcs).Parse(root_tmpl)
	if err != nil {
		log.Fatal(err)
	}
	overlayTmpl, err := template.Must(rootTmpl.Clone()).Parse(overlay_tmpl)
	if err != nil {
		log.Fatal(err)
	}
	if err := rootTmpl.Execute(os.Stdout, planets); err != nil {
		log.Fatal(err)
	}
	if err := overlayTmpl.Execute(os.Stdout, planets); err != nil {
		log.Fatal(err)
	}
}
