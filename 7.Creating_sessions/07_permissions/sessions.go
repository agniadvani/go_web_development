package main

import "net/http"

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
