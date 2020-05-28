package main

import (
	"html/template"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

var tpl *template.Template

type user struct {
	Username string
	First    string
	Last     string
}

var dbSessions = make(map[string]string)
var dbUsers = make(map[string]user)

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))

}

func main() {
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/", index)
	http.HandleFunc("/homepage", homepage)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	//Requesting a cookie
	c, err := req.Cookie("session")
	if err == http.ErrNoCookie {
		sID, _ := uuid.NewV4()
		c = &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		http.SetCookie(w, c)
	}
	var u user

	if req.Method == http.MethodPost {
		un := req.FormValue("username")
		f := req.FormValue("firstname")
		l := req.FormValue("lastname")

		u = user{un, f, l}
		dbSessions[c.Value] = un
		dbUsers[un] = u

	}
	// if user already exists
	if un, ok := dbSessions[c.Value]; ok {
		u = dbUsers[un]
	}

	err = tpl.ExecuteTemplate(w, "index.html", u)
}

func homepage(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("session")
	if err != nil {
		http.Redirect(w, req, "/", http.StatusSeeOther)
	}
	un, ok := dbSessions[c.Value]
	if !ok {
		http.Redirect(w, req, "/", http.StatusSeeOther)
	}
	u := dbUsers[un]

	err = tpl.ExecuteTemplate(w, "homepage.html", u)
}
