package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("template/template.gohtml"))
}

func main() {

	score := struct {
		ScoreA int
		ScoreB int
	}{
		32,
		44,
	}

	err := tpl.ExecuteTemplate(os.Stdout, "template.gohtml", score)
	if err != nil {
		log.Fatal(err)
	}
}
