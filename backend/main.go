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

	// - REST Schema: (prefix: https://\<backendurl\>/api/v0/)
	mux := http.NewServeMux()
	// loginH := LoginHandler{db: db}
	mux.Handle("/login", &LoginHandler{db})
	//    - /login
	//        - POST: should accept username and password
	//            - returns authentication cookie/token
	mux.Handle("/api/v0/user", &UserHandler{db})
	//    - /user
	//        - POST: should accept username and password
	//            - adds user to database
	//            - returns authentication cookie/token
	//        - DELETE: accept password in body?
	//			  - checks password matches jwt user password in database
	//            - deletes user from database
	//            - revokes jwt token
	mux.Handle("/api/v0/lists", &ListHandler{db})
	//    - /lists
	//        - GET: return all lists for user
	//        - POST: create new list
	// mux.Handle("/api/v0/list/{listid}", &ListHandler{db, testlist})
	//    - /lists/{list.name}
	//        - GET: return list metadata
	//        - PUT: update list metadata
	//        - DELETE: delete list
	// mux.Handle("/api/v0/list/{listid}/note/", &NoteHandler{db, testnote})
	//    - /lists/{list.name}/notes
	//        - GET: return all notes for list
	//        - POST: create new note
	// mux.Handle("/api/v0/list/{listid}/note/{noteid}", &NoteHandler{db, testnote})
	//    - /lists/{list.name}/notes/{noteid}
	//        - GET: returns note data
	//        - PUT: update note data
	//        - DELETE: delete note

	mux.HandleFunc("/test", qwerty)
	mux.HandleFunc("/", greet)
	mux.HandleFunc("/rad-140", jsonMayhaps)
	mux.HandleFunc("/favicon.ico", favicon)
	fmt.Printf("Starting Server...\n")
	http.ListenAndServe("0.0.0.0:"+port, mux)
	fmt.Printf("Now listening on port " + port + ".")
}
