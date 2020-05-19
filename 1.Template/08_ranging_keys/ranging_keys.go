package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func main() {
	tpl = template.Must(template.ParseGlob("template/template.gohtml"))
	rockstars := map[string]string{
		`Pink Floyd`:     `David Gilmour`,
		`Led Zeppelin`:   `Jimmy Page`,
		`Rolling Stones`: `Mick Jagger`,
		`The Doors`:      `Jim Morrison`,
	}
	err := tpl.ExecuteTemplate(os.Stdout, "template.gohtml", rockstars)
	if err != nil {
		log.Fatal(err)
	}
}
