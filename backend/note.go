package main

import (
	"database/sql"
	"fmt"
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

func GetNotes(db *sql.DB, userid int32) ([]Note, error) {
	// An user slice to hold data from returned rows.
	var notes []Note

	// query database for all notes with matching userid
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

func updateNote(db *sql.DB, newnote Note) error {
	// update row in notes table corresponding to noteid
	if _, err := db.Exec(
		`UPDATE notes SET content = '` + newnote.Content + `' WHERE id = '` + string(newnote.Id) + `';`); err != nil {
		return err
	}
	return nil
}

func addNote(db *sql.DB, newnote Note) error {
	// add row in notes table
	if _, err := db.Exec(
		`INSERT INTO note (userid, listid, content)
		 VALUES (` + string(newnote.Userid) + `, ` + string(newnote.Listid) + `, '` + newnote.Content + `');`); err != nil {
		return err
	}
	return nil
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
