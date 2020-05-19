package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type person struct {
	Name string
	Age  int
}

func init() {
	tpl = template.Must(template.ParseFiles("template/template.gohtml"))
}

func (p person) Dblage() int {
	return p.Age * 2
}

func (p person) Takesargs(x int) int {
	return x * 2
}

func main() {
	p1 := person{
		Name: "Micheal Scott",
		Age:  45,
	}

	err := tpl.Execute(os.Stdout, p1)
	if err != nil {
		log.Fatal(err)
	}
}
