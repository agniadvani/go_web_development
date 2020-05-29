package main

import (
	"log"
	"net/http"
	"text/template"

	uuid "github.com/satori/go.uuid"
)

type user struct {
	First    string
	Last     string
	Username string
	Password string
}

var dbSession = make(map[string]string)
var dbUser = make(map[string]user)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/", index)
	http.HandleFunc("/homepage", homepage)
	http.HandleFunc("/signup", signup)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	u := GetUser(req)
	err := tpl.ExecuteTemplate(w, "index.html", u)
	if err != nil {
		log.Fatal(err)
	}
}

func homepage(w http.ResponseWriter, req *http.Request) {
	u := GetUser(req)
	if !AlreadyLoggedIn(req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	tpl.ExecuteTemplate(w, "homepage.html", u)
}

func signup(w http.ResponseWriter, req *http.Request) {
	if AlreadyLoggedIn(req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	if req.Method == http.MethodPost {

		f := req.FormValue("First")
		l := req.FormValue("Last")
		un := req.FormValue("Username")
		p := req.FormValue("Password")

		if _, ok := dbSession[un]; ok {
			http.Error(w, "Username Exists", http.StatusForbidden)
			return
		}

		sID, _ := uuid.NewV4()
		c := &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		http.SetCookie(w, c)

		dbSession[c.Value] = un
		u := user{f, l, un, p}
		dbUser[un] = u

		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	err := tpl.ExecuteTemplate(w, "signup.html", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func GetUser(req *http.Request) user {
	var u user
	c, err := req.Cookie("session")
	if err != nil {
		return u
	}
	if un, ok := dbSession[c.Value]; ok {
		u = dbUser[un]
	}
	return u

}

func AlreadyLoggedIn(req *http.Request) bool {
	c, err := req.Cookie("session")
	if err != nil {
		return false
	}

	un := dbSession[c.Value]
	_, ok := dbUser[un]
	return ok
}
