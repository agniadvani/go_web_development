package main

import (
	"log"
	"os"
	"text/template"
	"time"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("template/template.gohtml"))
}

var fm = template.FuncMap{
	"dmy": daymonthyear,
}

func daymonthyear(t time.Time) string {
	return t.Format("02-01-2006")
}
func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "template.gohtml", time.Now())
	if err != nil {
		log.Fatal(err)
	}
}
