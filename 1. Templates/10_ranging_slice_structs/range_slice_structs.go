package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type rockstar struct {
	Name string
	Band string
}

func init() {
	tpl = template.Must(template.ParseGlob("template/template.gohtml"))

}
func main() {

	r1 := rockstar{
		`David Gilmour`,
		`Pink Floyd`,
	}

	r2 := rockstar{
		`Jimmy Page`,
		`Led Zeppelin`,
	}

	r3 := rockstar{
		`Mick Jagger`,
		`Rolling Stones`,
	}

	r4 := rockstar{
		`Jim Morrison`,
		`The Doors`,
	}

	r5 := rockstar{
		`James Hetfield`,
		`Metallica`,
	}

	rs := []rockstar{r1, r2, r3, r4, r5}

	err := tpl.ExecuteTemplate(os.Stdout, "template.gohtml", rs)
	if err != nil {
		log.Fatal(err)
	}
}
