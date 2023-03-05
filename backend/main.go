package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

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

	// REST Schema: (prefix: https://<backendurl>/api/v0/)
	mux := http.NewServeMux()

	mux.Handle("/login", &LoginHandler{db})
	// POST: should accept username and password
	//   - returns authentication cookie/token

	mux.Handle("/user", &UserHandler{db})
	// POST: should accept username and password
	//   - adds user to database
	//   - returns authentication cookie/token
	// DELETE: accept password in body?
	//   - checks password matches jwt user password in database
	//   - deletes user from database
	//   - revokes jwt token

	mux.Handle("/lists", &ListHandler{db})
	// GET: return all lists for user
	// POST: create new list

	// mux.Handle("/api/v0/list/{listid}", &ListHandler{db, testlist})
	// GET: return list metadata
	// PUT: update list metadata
	// DELETE: delete list

	mux.Handle("/notes", &NoteHandler{db})
	// GET: return all notes for list
	// POST: create new note

	// mux.Handle("/api/v0/list/{listid}/note/{noteid}", &NoteHandler{db, testnote})
	// GET: returns note data
	// PUT: update note data
	// DELETE: delete note

	mux.HandleFunc("/", greet)
	fmt.Printf("Starting Server...\n")
	http.ListenAndServe("0.0.0.0:"+port, mux)
	fmt.Printf("Now listening on port " + port + ".")
}

func greet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Fprintf(w, "Hello World! %s", time.Now())
}

//	mux.HandleFunc("/test", qwerty)
//	mux.HandleFunc("/rad-140", jsonMayhaps)
//	mux.HandleFunc("/favicon.ico", favicon)

// func favicon(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Access-Control-Allow-Origin", "*")
// 	http.ServeFile(w, r, "./favicon_io/favicon.ico")
// }

// func jsonMayhaps(w http.ResponseWriter, r *http.Request) {
// 	rawData := []Note{
// 		{Id: 0, Userid: 1, Listid: 1, Content: "Do your math 61 hwk."},
// 		{Id: 1, Userid: 1, Listid: 1, Content: "This is another test."},
// 	}
//
// 	data, err := json.Marshal(rawData)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
//
// 	w.Header().Set("Content-Type", "application/json")
// 	w.Header().Set("Access-Control-Allow-Origin", "*")
// 	w.Write(data)
// }

// func qwerty(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Access-Control-Allow-Origin", "*")
// 	w.Write([]byte("This is a string.\n"))
// 	w.Write([]byte(html.EscapeString(r.URL.Path)))
// 	w.Write([]byte("Another test"))

// 	fmt.Printf("accessed\n")

// }
