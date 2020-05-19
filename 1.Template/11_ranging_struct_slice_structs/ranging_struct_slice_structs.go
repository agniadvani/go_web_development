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

type actor struct {
	Name  string
	Movie string
}

type celebrities struct {
	Rocknroll []rockstar
	Roleplay  []actor
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

	a1 := actor{
		`Tom Cruise`,
		`Top Gun`,
	}

	a2 := actor{
		`Daniel Craig`,
		`Quantum Of Solace`,
	}

	rs := []rockstar{r1, r2, r3}
	ac := []actor{a1, a2}

	data := celebrities{
		Rocknroll: rs,
		Roleplay:  ac,
	}

	err := tpl.ExecuteTemplate(os.Stdout, "template.gohtml", data)
	if err != nil {
		log.Fatal(err)
	}
}
