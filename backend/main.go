package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"html"
	"log"
	"net/http"
	"os"
	"time"

	//_  "github.com/cockroachdb/cockroach-go/crdb"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "4000"
	}

	db := databaseInit()

	// - REST Schema: (prefix: https://\<backendurl\>/api/v0/)
	//    - /login
	//        - POST: should accept username and password
	//            - returns authentication cookie/token (stored in browser by frontend)
	//    - /lists
	//        - GET: return all lists for user
	//        - POST: create new list
	//    - /lists/{list.name}
	//        - GET: return list metadata
	//        - PUT: update list metadata
	//        - DELETE: delete list
	//    - /lists/{list.name}/notes
	//        - GET: return all notes for list
	//        - POST: create new note
	//    - /lists/{list.name}/notes/{noteid}
	//        - GET: returns note data
	//        - PUT: update note data
	//        - DELETE: delete note

	clayton := User{Username: "clayton", Password: "test123"}
	// testnote := Note{Id: 0, Userid: 3, Listid: 0, Content: "This is a reminder to eat ur veggies."}
	testlist := List{Id: 0, Userid: 3, Name: "Important"}

	mux := http.NewServeMux()
	mux.Handle("/login", &LoginHandler{db, clayton})
	mux.Handle("/api/v0/user", &UserHandler{db, clayton})
	mux.Handle("/api/v0/list", &ListHandler{db, testlist})
	// mux.Handle("/api/v0/list/{listid}", &ListHandler{db, testlist})
	// mux.Handle("/api/v0/list/{listid}/note/", &NoteHandler{db, testnote})
	// mux.Handle("/api/v0/list/{listid}/note/{noteid}", &NoteHandler{db, testnote})

	mux.HandleFunc("/test", qwerty)
	mux.HandleFunc("/", greet)
	mux.HandleFunc("/rad-140", jsonMayhaps)
	mux.HandleFunc("/favicon.ico", favicon)

	fmt.Printf("Starting Server...\n")
	http.ListenAndServe("0.0.0.0:"+port, mux)
	fmt.Printf("Now listening on port " + port + ".")
}

type User struct {
	Id       int32
	Username string
	Password string
}

type List struct {
	Id     int32
	Userid int32
	Name   string
}

type Note struct {
	Id      int32
	Userid  int32
	Listid  int32
	Content string
}

type LoginHandler struct {
	db  *sql.DB
	usr User
}

type UserHandler struct {
	db  *sql.DB
	usr User
}

type ListHandler struct {
	db   *sql.DB
	list List
}

type NoteHandler struct {
	db   *sql.DB
	note Note
}

func (h *LoginHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// if POST request (all other types fail)
	// handle post request containing username and password in plaintext
	//
	// if user exists, check if password is valid, else return error
	//
	// return jwt token derived from username (include userid?)
}

func (h *UserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// accepts
	hashed, _ := bcrypt.GenerateFromPassword([]byte(h.usr.Password), 4)
	h.usr.Password = string(hashed)
	if _, err := h.db.Exec(
		`INSERT INTO users (username, password)
		 VALUES ('` + h.usr.Username + `', '` + h.usr.Password + `');`); err != nil {
		log.Fatal(err)
	}
}

func (h *ListHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// extract userid from jwt token
	//
	// if GET request return all lists
	//
	// if POST request then,
	// add row in notes table with corresponding userid, listid, and content
	// if _, err := h.db.Exec(
	// 	`INSERT INTO note (userid, listid, content)
	// 	 VALUES (` + string(h.note.Userid) + `, ` + string(h.note.Listid) + `, '` + h.note.Content + `');`); err != nil {
	// 	log.Fatal(err)
	// }
	//
	// If PUT request
	// extract noteid and note content from POST content
	// update row in notes table corresponding to noteid
	// If DELETE request
	// extract noteid
	// delete corresponding row in notes table
}

func (h *NoteHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// extract userid from jwt token
	//
	// if GET request return all notes
	//
	// if POST request then,
	// extract listid and note content from POST content
	// add row in notes table with corresponding userid, listid, and content
	// if _, err := h.db.Exec(
	// 	`INSERT INTO note (userid, listid, content)
	// 	 VALUES (` + string(h.note.Userid) + `, ` + string(h.note.Listid) + `, '` + h.note.Content + `');`); err != nil {
	// 	log.Fatal(err)
	// }
	//
	// If PUT request
	// extract noteid and note content from POST content
	// update row in notes table corresponding to noteid
	// If DELETE request
	// extract noteid
	// delete corresponding row in notes table
}

func databaseInit() *sql.DB {
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

	// Initialize database connection
	var connStr string
	if dbpass == "" {
		connStr = "postgresql://" + dbuser + "@" + dburl + "/defaultdb?sslmode=disable"
	} else {
		connStr = "postgresql://" + dbuser + ":" + dbpass + "@" + dburl + "/defaultdb?sslmode=disable"
	}

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	} else if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
	// defer db.Close()
	fmt.Println("Database Connected!")

	// Create the "users" table.
	if _, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY, 
			username STRING NOT NULL UNIQUE, 
			password STRING NOT NULL 
		)`); err != nil {
		log.Fatal(err)
	}

	// Create the "list" table.
	if _, err := db.Exec(
		`CREATE TABLE IF NOT EXISTS list (
	           id SERIAL PRIMARY KEY,
	           userid INT NOT NULL REFERENCES users(Id) ON UPDATE CASCADE ON DELETE CASCADE,
	           name STRING NOT NULL
		   )`); err != nil {
		log.Fatal(err)
	}

	// Create the "note" table.
	if _, err := db.Exec(
		`CREATE TABLE IF NOT EXISTS note (
	           id SERIAL PRIMARY KEY,
	           userid INT NOT NULL REFERENCES users(Id) ON UPDATE CASCADE ON DELETE CASCADE,
	           listid INT NOT NULL REFERENCES list(Id) ON UPDATE CASCADE ON DELETE CASCADE,
	           content VARCHAR(280)
	           )`); err != nil {
		log.Fatal(err)
	}

	// Ensure admin user is in "users" table.
	if _, err := db.Exec(
		`INSERT INTO users (id, username, password)
		 VALUES (0, 'admin', 'admin');`); err != nil {
		log.Fatal(err)
	}
	return db
}

func allUsers(db *sql.DB) ([]User, error) {
	rows, err := db.Query(`SELECT id FROM users WHERE username = ?`, "admin")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// An user slice to hold data from returned rows.
	var users []User

	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var usr User
		if err := rows.Scan(&usr.Id, &usr.Username, &usr.Password); err != nil {
			return users, err
		}
		users = append(users, usr)
	}
	if err = rows.Err(); err != nil {
		return users, err
	}
	return users, nil
}

func jsonMayhaps(w http.ResponseWriter, r *http.Request) {
	rawData := []Note{
		{Id: 0, Userid: 1, Listid: 1, Content: "Do your math 61 hwk."},
		{Id: 1, Userid: 1, Listid: 1, Content: "This is another test."},
	}

	data, err := json.Marshal(rawData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(data)
}

func favicon(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	http.ServeFile(w, r, "./favicon_io/favicon.ico")
}

func qwerty(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write([]byte("This is a string.\n"))
	w.Write([]byte(html.EscapeString(r.URL.Path)))
	w.Write([]byte("Another test"))

	fmt.Printf("accessed\n")
}

func greet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Fprintf(w, "Hello World! %s", time.Now())
}
