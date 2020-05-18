package main

import (
	"log"
	"os"
	"text/template"
)

type rockstar struct {
	Name  string
	Song  string
	Alive bool
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("template/template.gohtml"))
}

func main() {
	r1 := rockstar{
		"David Gilmour",
		"High Hopes",
		true,
	}
	r2 := rockstar{
		"Freddie Mercury",
		"Bohemian Rhapsody",
		false,
	}
	r3 := rockstar{
		"",
		"Baby",
		true,
	}

	xr := []rockstar{r1, r2, r3}

	err := tpl.ExecuteTemplate(os.Stdout, "template.gohtml", xr)
	if err != nil {
		log.Fatal(err)
	}
}
