package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstthree,
}
var tpl *template.Template

type rockstar struct {
	Name string
	Band string
}

func main() {
	tpl := template.Must(template.New("").Funcs(fm).ParseGlob("template/template.gohtml"))
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

	rs := []rockstar{r1, r2, r3}

	err := tpl.ExecuteTemplate(os.Stdout, "template.gohtml", rs)
	if err != nil {
		log.Fatal(err)
	}

}
func firstthree(s string) string {
	s = strings.TrimSpace(s)
	s = s[:3]
	return s

}
