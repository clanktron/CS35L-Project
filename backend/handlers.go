package main

import (
	"database/sql"
	"fmt"
	"html"
	"log"
	"net/http"
	"time"
)

type UserHandler struct {
	db *sql.DB
}

type NoteHandler struct {
	db *sql.DB
}

type ListHandler struct {
	db *sql.DB
}

type LoginHandler struct {
	db *sql.DB
}

func (h *LoginHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// check auth function
	if r.Method == http.MethodGet {
		fmt.Printf("oogly boogly")
	} else if r.Method == http.MethodPost {
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
			log.Fatal(err)
		}
		// fmt.Print(newuser)
		// hash user password
		newuser.Password, err = hashPassword(newuser.Password)
		return
	} else {
		http.Error(w, fmt.Sprintf("Expected method POST, got %v", r.Method), http.StatusMethodNotAllowed)
		return
	}

	// else if r.Method == http.MethodDelete {
	// 	var newuser User
	// 	parseJSON(w, r, &newuser)
	// 	// insert new user object in database

	// 	return
	// }

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
		AddList(h.db, newlist)
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
		parseJSON(w, r, &newnote)
		// store name and userid
		addNote(h.db, newnote)
		return
	} else if r.Method == http.MethodPut {
		// create new note object
		var newnote Note
		parseJSON(w, r, &newnote)
		// update name and userid
		updateNote(h.db, newnote)
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
