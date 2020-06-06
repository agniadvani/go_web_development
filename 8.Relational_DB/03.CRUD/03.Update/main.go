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
	db, err = sql.Open("mysql", "user:password@tcp(Loaclhost:5555)/dbname?charset=utf8")
	check(err)
	defer db.Close()
	err = db.Ping()
	check(err)
	http.HandleFunc("/", index)
	http.HandleFunc("/create", create)
	http.HandleFunc("/insert", insert)
	http.HandleFunc("/read", read)
	http.HandleFunc("/update", update)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Database successfully connected.")
}

func create(w http.ResponseWriter, r *http.Request) {
	stmt, err := db.Prepare(`CREATE TABLE customer (name VARCHAR(20));`)
	check(err)

	rows, err := stmt.Exec()
	check(err)

	n, err := rows.RowsAffected()
	check(err)

	fmt.Fprintln(w, "Created Table customer", n)
}

func insert(w http.ResponseWriter, r *http.Request) {
	stmt, err := db.Prepare(`INSERT INTO customer VALUES("James");`)
	check(err)

	rows, err := stmt.Exec()
	check(err)

	n, err := rows.RowsAffected()
	check(err)

	fmt.Fprintln(w, "Inserted Records = ", n)
}

func read(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query(`SELECT * FROM customer`)
	check(err)
	var name string
	for rows.Next() {
		err = rows.Scan(&name)
		check(err)
	}
	fmt.Fprintln(w, "Records:", name)
}

func update(w http.ResponseWriter, r *http.Request) {
	stmt, err := db.Prepare(`UPDATE customer SET name = "Jimmy" WHERE name = "James";`)
	check(err)

	rows, err := stmt.Exec()
	check(err)

	n, err := rows.RowsAffected()
	check(err)

	fmt.Fprintln(w, "Records Updated:", n)
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
