package main

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func main() {
	db, err = sql.Open("mysql", "user:password@tcp(localhost:5555)/dbname?charset=utf8")
	check(err)
	defer db.Close()
	err = db.Ping()
	check(err)
	http.HandleFunc("/", index)
	http.HandleFunc("/amigos", amigos)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Database successfully connected.")
}

func amigos(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query(`SELECT aName from amigo`)
	check(err)
	var s, name string
	s = "List of Amigos:\n"
	for rows.Next() {
		err = rows.Scan(&name)
		check(err)
		s += name + "\n"
	}
	fmt.Fprintln(w, s)
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
