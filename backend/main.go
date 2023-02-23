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
	mux := http.NewServeMux()
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
	db *sql.DB
}

type UserHandler struct {
	db *sql.DB
}

type ListHandler struct {
	db *sql.DB
}

type NoteHandler struct {
	db *sql.DB
}

func generateJWT() {
	// maybe put this in main()?
	// secret := os.Getenv("JWTSECRET")
	// if secret == "" {
	// 	secret = "12043$521p8ijz4"
	// }
	// generate jwt from userid, secret, and
}
func terminateJWT() {
	// replace jwt with another that expires immediately
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 4)
	return string(bytes), err
}

func CheckPasswordHash(password string, hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err
}

func renderJSON(w http.ResponseWriter, v interface{}) error {
	js, err := json.Marshal(v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
	return err
}

func parseJSON(w http.ResponseWriter, r *http.Request, v interface{}) error {
	err := json.NewDecoder(r.Body).Decode(v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Fatal(err)
	}
	return err
}

func (h *LoginHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// check auth function
	if r.Method == http.MethodPost {
		// handle post request containing username and password in plaintext
		var login User
		parseJSON(w, r, &login)
		// TODO: query database for user with matching username and password
		// rows, err := h.db.Query(`SELECT * FROM users WHERE username  = ?`, login.Username)
		//
		// TODO: respond with jwt token derived from username (include userid?)
	} else {
		http.Error(w, fmt.Sprintf("Expected method POST, got %v", r.Method), http.StatusMethodNotAllowed)
		return
	}
}

func (h *UserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// create new user object and store username and password
		var newuser User
		err := parseJSON(w, r, &newuser)
		if err != nil {
		}
		// hash user password
		newuser.Password, err = HashPassword(newuser.Password)
		// insert new user object in database
		if _, err := h.db.Exec(
			`INSERT INTO users (username, password)
			VALUES ('` + newuser.Username + `', '` + newuser.Password + `');`); err != nil {
			log.Fatal(err)
		} else {
			http.Error(w, fmt.Sprintf("Expected method POST, got %v", r.Method), http.StatusMethodNotAllowed)
			return
		}
		return
	} else if r.Method == http.MethodDelete {
		var newuser User
		parseJSON(w, r, &newuser)
		// insert new user object in database
		if _, err := h.db.Exec(
			`INSERT INTO users (username, password)
			VALUES ('` + newuser.Username + `', '` + newuser.Password + `');`); err != nil {
			log.Fatal(err)
		} else {
			http.Error(w, fmt.Sprintf("Expected method POST, got %v", r.Method), http.StatusMethodNotAllowed)
			return
		}
		return
	}
}

func (h *ListHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// TODO: call auth func; if unauthorized redirect to login (will be done in func)
	//
	anvar := List{Id: 0, Userid: 3, Name: "Important"}
	if r.Method == http.MethodGet {
		renderJSON(w, anvar)
		return
		// return all lists
	} else if r.Method == http.MethodPost {
		// create new list object
		var newlist List
		// store name and userid
		parseJSON(w, r, &newlist)
		// add row in lists table with corresponding userid and name
		if _, err := h.db.Exec(
			`INSERT INTO lists (userid, name)
			 VALUES (` + string(newlist.Userid) + `, ` + string(newlist.Name) + `);`); err != nil {
			log.Fatal(err)
		}
		return
	} else if r.Method == http.MethodPut {
		// TODO: extract listname from url path
		// create new list object
		var newlist List
		// store name and userid
		parseJSON(w, r, &newlist)
		// TODO: update row in lists table with corresponding userid and name
		// if _, err := h.db.Exec(
		// 	`UPDATE lists SET name = '` + newlist.Name + `' WHERE id = '` + listname + `';`); err != nil {
		// 	log.Fatal(err)
		// }
		return
	} else if r.Method == http.MethodDelete {
		// TODO: extract list name from path
		// delete corresponding row in lists table
		// if _, err := h.db.Exec(`DELETE FROM lists WHERE name = '` + listname + `';`); err != nil {
		// 	log.Fatal(err)
		// }
		return
	} else {
		http.Error(w, fmt.Sprintf("Expected method GET, POST, PUT, or DELETE, got %v", r.Method), http.StatusMethodNotAllowed)
		return
	}
}

func (h *NoteHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// TODO: call auth func (middleware?)
	//
	if r.Method == http.MethodGet {
		// TODO: get userid from url path
		// get all notes for userid
		// notes, err := getNotes(h.db, userid)
		// if err != nil {
		// 	log.Fatal(err)
		// }
		// // respond with all note objects
		// renderJSON(w, notes)
		return
	} else if r.Method == http.MethodPost {
		// create new note object
		var newnote Note
		// store name and userid
		parseJSON(w, r, &newnote)
		// add row in notes table
		if _, err := h.db.Exec(
			`INSERT INTO note (userid, listid, content)
			 VALUES (` + string(newnote.Userid) + `, ` + string(newnote.Listid) + `, '` + newnote.Content + `');`); err != nil {
			log.Fatal(err)
		}
		return
	} else if r.Method == http.MethodPut {
		// create new note object
		var newnote Note
		// store name and userid
		parseJSON(w, r, &newnote)
		// update row in notes table corresponding to noteid
		if _, err := h.db.Exec(
			`UPDATE notes SET content = '` + newnote.Content + `' WHERE id = '` + string(newnote.Id) + `';`); err != nil {
			log.Fatal(err)
		}
		return
	} else if r.Method == http.MethodDelete {
		// TODO: extract noteid path
		// delete row in notes table corresponding to noteid
		//if _, err := h.db.Exec(
		//	`DELETE	FROM notes WHERE Id = '` + noteid + `';`); err != nil {
		//	log.Fatal(err)
		//}
		return
	} else {
		http.Error(w, fmt.Sprintf("Expected method GET, POST, PUT, or DELETE, got %v", r.Method), http.StatusMethodNotAllowed)
		return
	}

}

func getLists(db *sql.DB, userid int32) ([]List, error) {
	rows, err := db.Query(`SELECT * FROM lists WHERE userid = ?`, userid)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// An user slice to hold data from returned rows.
	var lists []List

	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var lst List
		if err := rows.Scan(&lst.Id, &lst.Userid, &lst.Name); err != nil {
			return lists, err
		}
		lists = append(lists, lst)
	}
	if err = rows.Err(); err != nil {
		return lists, err
	}
	return lists, nil
}

func getNotes(db *sql.DB, userid int32) ([]Note, error) {
	// query database for all notes with matching userid
	rows, err := db.Query(`SELECT * FROM notes WHERE userid = ?`, userid)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// An user slice to hold data from returned rows.
	var notes []Note

	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var nt Note
		if err := rows.Scan(&nt.Id, &nt.Userid, &nt.Listid, &nt.Content); err != nil {
			return notes, err
		}
		notes = append(notes, nt)
	}
	if err = rows.Err(); err != nil {
		return notes, err
	}
	return notes, nil
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
	// Note: Possibly change SERIAL (postgres & cockroach) to UUID (cockroach only)
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
