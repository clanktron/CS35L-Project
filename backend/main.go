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
)

type User struct {
	Id       int32
	Username string
	Password string
}

type List struct {
	Id     int32
	Userid int32
	List   string
}

type Note struct {
	Id      int32
	Userid  int32
	Listid  int32
	Content string
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "4000"
	}
	dbuser := os.Getenv("DBUSER")
	if dbuser == "" {
		dbuser = "root"
	}
	//	dbpass := os.Getenv("DBPASS")
	//	if dbpass == "" {
	//		dbpass = ""
	//	}
	dburl := os.Getenv("DBURL")
	if dburl == "" {
		dburl = "0.0.0.0:26257"
	}

	// Initialize database connection
	connStr := "postgresql://" + dbuser + "@" + dburl + "/defaultdb?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	} else if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	fmt.Println("Database Connected!")

	// Create the "users" table.
	if _, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS account (
			Id SERIAL PRIMARY KEY, 
			username CHAR(20) NOT NULL, 
			password CHAR(20) NOT NULL
		)`); err != nil {
		log.Fatal(err)
	}

	if _, err := db.Exec(
		`CREATE TABLE IF NOT EXISTS list (
	           Id SERIAL PRIMARY KEY,
	           Userid INT REFERENCES account(Id) NOT NULL,
	           name CHAR(20) NOT NULL
	           )`); err != nil {
		log.Fatal(err)
	}

	if _, err := db.Exec(
		`CREATE TABLE IF NOT EXISTS note (
	           Id INT PRIMARY KEY,
	           Userid INT REFERENCES account(Id) NOT NULL,
	           Listid INT REFERENCES list(Id) NOT NULL,
	           content VARCHAR(280)
	           )`); err != nil {
		log.Fatal(err)
	}

	// Ensure admin user is in "account" table.
	// if _, err := db.Exec(
	// 	`INSERT INTO account (
	// 		Id,
	// 		username,
	// 		password
	// 	) VALUES (
	// 		1,
	// 		admin,
	// 		admin
	// 	)`); err != nil {
	// 	log.Fatal(err)
	// }

	// resp, err := http.Get("http://google.com/")
	// if err != nil {
	// 	// handle error
	// }
	// fmt.Println(resp)
	// defer resp.Body.Close()
	//	body, err := io.ReadAll(resp.Body)
	//
	http.HandleFunc("/test", qwerty)
	http.HandleFunc("/", greet)
	http.HandleFunc("/rad-140", jsonMayhaps)
	http.HandleFunc("/favicon.ico", favicon)
	fmt.Printf("Starting Server...\n")
	http.ListenAndServe("0.0.0.0:"+port, nil)
	fmt.Printf("Now listening on port " + port + ".")
}

// func databaseTest(w http.ResponseWriter, r *http.Request, db *sql.DB) {
// 	var res string
// 	var todos []string
// 	rows, err := db.Query("SELECT * FROM todos")
// 	defer rows.Close()
// 	if err != nil {
// 		log.Fatalln(err)
// 		w.Write([]byte("An error occurred.\n"))
// 		fmt.Printf("A database error occurred.")
// 	}
// 	for rows.Next() {
// 		rows.Scan(&res)
// 		todos = append(todos, res)
// 	}
// }

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
