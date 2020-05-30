package main

import (
	"log"
	"net/http"
	"text/template"

	"golang.org/x/crypto/bcrypt"

	uuid "github.com/satori/go.uuid"
)

type user struct {
	First    string
	Last     string
	Username string
	Password []byte
	role     string
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
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
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
	if u.role != "admin" {
		http.Error(w, "Only Admins have access to this page", http.StatusForbidden)
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
		r := req.FormValue("role")

		if _, ok := dbUser[un]; ok {
			http.Error(w, "Username Exists", http.StatusForbidden)
			return
		}

		sID, _ := uuid.NewV4()
		c := &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		http.SetCookie(w, c)
		//Encrypting Password
		bs, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.MinCost)
		if err != nil {
			log.Fatal(err)
		}
		dbSession[c.Value] = un
		u := user{f, l, un, bs, r}
		dbUser[un] = u

		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	err := tpl.ExecuteTemplate(w, "signup.html", nil)
	if err != nil {
		log.Fatal(err)
	}
}
func login(w http.ResponseWriter, req *http.Request) {
	if AlreadyLoggedIn(req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
	}
	if req.Method == http.MethodPost {
		un := req.FormValue("Username")
		p := req.FormValue("Password")

		u, ok := dbUser[un]
		if !ok {
			http.Error(w, "Username and/or password is incorrect", http.StatusForbidden)
			return
		}
		sID, _ := uuid.NewV4()
		c := &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		http.SetCookie(w, c)

		err := bcrypt.CompareHashAndPassword(u.Password, []byte(p))
		if err != nil {
			http.Error(w, "Username and/or password is incorrect.", http.StatusForbidden)
			return
		}
		dbSession[c.Value] = un

		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	tpl.ExecuteTemplate(w, "login.html", nil)
}

func logout(w http.ResponseWriter, req *http.Request) {
	if !AlreadyLoggedIn(req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
	}
	c, _ := req.Cookie("session")
	delete(dbSession, c.Value)
	c = &http.Cookie{
		Name:   "session",
		Value:  "",
		MaxAge: -1,
	}
	http.SetCookie(w, c)
	http.Redirect(w, req, "/login", http.StatusSeeOther)
}
