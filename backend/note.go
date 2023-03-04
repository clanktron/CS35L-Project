package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
)

type Note struct {
	Id      int32
	Userid  int32
	Listid  int32
	Content string
}

type NoteHandler struct {
	db *sql.DB
}

func (h *NoteHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	var note Note
	var err error

	note.Userid, err = verifyJWT(w, r)
	if err != nil {
		return
	}

	if r.Method == http.MethodGet {

		notes, err := getNotes(h.db, note.Userid)
		if err != nil {
			log.Fatal(err)
		}
		renderJSON(w, notes)

		return
	} else if r.Method == http.MethodPost {

		parseJSON(w, r, &note)
		addNote(h.db, note)

		return
	} else if r.Method == http.MethodPut {

		parseJSON(w, r, &note)
		updateNote(h.db, note)

		return
	} else if r.Method == http.MethodDelete {
		// TODO: extract noteid from path
		var noteid int32
		note.Id = noteid

		deleteNote(h.db, note)

		return
	} else {
		http.Error(w, fmt.Sprintf("Expected method GET, POST, PUT, or DELETE, got %v", r.Method), http.StatusMethodNotAllowed)
		return
	}

}

// delete row in notes table corresponding to noteid
func deleteNote(db *sql.DB, note Note) error {
	if _, err := db.Exec(
		`DELETE	FROM notes WHERE id = '` + string(note.Id) + `';`); err != nil {
		log.Print(err)
		return err
	}
	return nil
}

// queries database for all notes corresponding to user - returns a Notes slice and error
func getNotes(db *sql.DB, userid int32) ([]Note, error) {
	// An user slice to hold data from returned rows.
	var notes []Note

	// query database for all notes with matching username
	rows, err := db.Query(`SELECT * FROM notes WHERE userid = ?`, userid)
	if err != nil {
		return notes, err
	}
	defer rows.Close()

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

// update row in notes table corresponding to noteid
func updateNote(db *sql.DB, newnote Note) error {
	if _, err := db.Exec(
		`UPDATE notes SET content = '` + newnote.Content + `' WHERE id = '` + string(newnote.Id) + `';`); err != nil {
		return err
	}
	return nil
}

// add row in notes table
func addNote(db *sql.DB, newnote Note) error {
	if _, err := db.Exec(
		`INSERT INTO note (userid, listid, content)
		 VALUES (` + string(newnote.Userid) + `, ` + string(newnote.Listid) + `, '` + newnote.Content + `');`); err != nil {
		return err
	}
	return nil
}
