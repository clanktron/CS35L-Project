package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	//_  "github.com/cockroachdb/cockroach-go/crdb"
	_ "github.com/lib/pq"
)

func main() {
	dbuser := os.Getenv("DBUSER")
	if dbuser == "" {
		dbuser = "root"
	}
	dbpass := os.Getenv("DBPASS")
	if dbpass == "" {
		dbpass = ""
	}
	dburl := os.Getenv("DBURL")
	if dburl == "" {
		dburl = "0.0.0.0:26257"
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "4000"
	}

	db, err := databaseInit(dbuser, dbpass, dburl)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Database Connected!")
	defer db.Close()

	fmt.Printf("Starting Server...\n")

	a := &App{
		LoginHandler: &LoginHandler{db},
		UserHandler:  &UserHandler{db},
		ListHandler:  &ListHandler{&NoteHandler{db}, db},
	}
	http.ListenAndServe("0.0.0.0:"+port, a)
	fmt.Printf("Now listening on port " + port + ".")
}
