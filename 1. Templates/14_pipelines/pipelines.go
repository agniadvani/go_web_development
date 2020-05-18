package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

var fm = template.FuncMap{
	"dblr": dblr,
	"sqr":  sqr,
	"sqrt": sqrt,
}

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseGlob("template/template.gohtml"))
}
func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "template.gohtml", 4)
	if err != nil {
		log.Fatal(err)
	}

}

func dblr(x int) int {
	dbl := x + x
	return dbl
}
func sqr(x int) int {
	sq := x * x
	return sq
}
func sqrt(x int) int {
	if x == 0 {
		return 0
	}
	if x == 1 {
		return 1
	}
	for i := 0; i <= x; i++ {
		if i*i == x {
			return i
		}
	}
	return 0
}
