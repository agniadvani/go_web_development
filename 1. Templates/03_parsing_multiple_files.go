package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	tpl, err := template.ParseFiles("boilet.gohtml")
	if err != nil {
		log.Fatal(err)
	}

	tpl, err = tpl.ParseFiles("two.gohtml", "three.gohtml")
	if err != nil {
		log.Fatal(err)
	}
	tpl.ExecuteTemplate(os.Stdout, "boilet.gohtml", nil)
	tpl.ExecuteTemplate(os.Stdout, "two.gohtml", nil)
	tpl.ExecuteTemplate(os.Stdout, "three.gohtml", nil)
}
