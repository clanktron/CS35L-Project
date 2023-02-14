package main

import (
	"fmt"
	"html"
	"net/http"
	"time"
)

func qwerty(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "nuttyyyyy\n")
	fmt.Printf("accessed\n")
	fmt.Fprintf(w, "hello %q", html.EscapeString(r.URL.Path))
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! %s", time.Now())
}

func main() {
	fmt.Printf("Starting Server...\n")
	http.HandleFunc("/test", qwerty)
	http.HandleFunc("/", greet)
	http.ListenAndServe("0.0.0.0:8080", nil)
}

// package main
//
// import (
// 	"database/sql"
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"os"
//
// 	//	"context"
// 	// no gorilla mux
// 	"github.com/gorilla/mux"
// 	//	"github.com/cockroachdb/cockroach-go/crdb"
// 	_ "github.com/lib/pq"
// )
//
// func main() {
// 	// Initialize database connection
// 	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL")+"&application_name=$ Jot")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer db.Close()
//
// 	// Create the "users" table.
// 	if _, err := db.Exec(
// 		`CREATE TABLE IF NOT EXISTS account (
//             userid INT PRIMARY KEY,
//             username CHAR(20),
//             password CHAR(20)
//             `); err != nil {
// 		log.Fatal(err)
// 	}
//
// 	if _, err := db.Exec(
// 		`CREATE TABLE IF NOT EXISTS list (
//             listid INT PRIMARY KEY,
//             userid INT REFERENCES account(userid),
//             list CHAR(20),
//             `); err != nil {
// 		log.Fatal(err)
// 	}
//
// 	if _, err := db.Exec(
// 		`CREATE TABLE IF NOT EXISTS note (
//             noteid INT PRIMARY KEY,
//             userid INT REFERENCES account(userid),
//             listid INT REFERENCES list(listid),
//             note VARCHAR(280)
//             `); err != nil {
// 		log.Fatal(err)
// 	}
//
// 	// Ensure admin user is in "account" table.
// 	if _, err := db.Exec(
// 		"INSERT INTO account (userid, username, password) VALUES (1, admin, adminpass)"); err != nil {
// 		log.Fatal(err)
// 	}
//
// 	resp, err := http.Get("http://google.com/")
// 	if err != nil {
// 		// handle error
// 	}
// 	fmt.Println(resp)
// 	defer resp.Body.Close()
// 	//	body, err := io.ReadAll(resp.Body)
// 	router := mux.NewRouter()
//
// 	router.HandleFunc("/login", login).Methods("GET")
//
// 	fmt.Println("hello world")
// }
