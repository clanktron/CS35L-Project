package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
)

type Note struct {
	Id      int64
	Userid  int64
	Listid  int64
	Content string
}

type NoteHandler struct {
	db *sql.DB
}

func (h *NoteHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	var note Note
	// var err error

	enableCors(w)

	// note.Userid, err = verifyJWT(r)
	// if err != nil {
	// 	return
	// }

	if r.Method == http.MethodGet {

		note.Userid = 0

		notes, err := getNotes(h.db, note.Userid)
		if err != nil {
			log.Fatal(err)
			log.Print("Failed to get notes from database\n")
			return
		}
		if err := renderJSON(w, notes); err != nil {
			log.Print(err)
			log.Printf("Json response failed\n")
			return
		}

		return
	} else if r.Method == http.MethodPost {

		if err := parseJSON(w, r, &note); err != nil {
			log.Print(err)
			log.Printf("Failed to parse json payload\n")
			return
		}
		note.Listid = 845202198593536001
		if err := addNote(h.db, note); err != nil {
			log.Print(err)
			log.Print("Failed to add note to database\n")
			return
		}
		log.Printf("Added note to list with id %d", note.Listid)

		return
	} else if r.Method == http.MethodPut {

		if err := parseJSON(w, r, &note); err != nil {
			log.Print(err)
			log.Printf("Failed to parse json payload\n")
			return
		}
		note.Listid = 845202198593536001
		if err := updateNote(h.db, note); err != nil {
			log.Print(err)
			log.Print("Failed to update note in database\n")
			return
		}
		log.Printf("Updated note with content %s", note.Content)

		return
	} else if r.Method == http.MethodDelete {
		// TODO: extract noteid from path
		var noteid int64
		note.Id = noteid

		if err := deleteNote(h.db, note); err != nil {
			log.Print(err)
			log.Print("Failed to delete note from database\n")
			return
		}
		log.Printf("Deleted note with content %s", note.Content)

		return
	} else if r.Method == http.MethodOptions {
		return
	} else {
		http.Error(w, fmt.Sprintf("Expected method GET, POST, PUT, or DELETE, got %v", r.Method), http.StatusMethodNotAllowed)
		return
	}

}

// delete row in notes table corresponding to noteid
func deleteNote(db *sql.DB, note Note) error {
	if _, err := db.Exec(
		`DELETE	FROM notes WHERE id = ?;`, note.Id); err != nil {
		log.Print(err)
		return err
	}
	return nil
}

// queries database for all notes corresponding to user - returns a Notes slice and error
func getNotes(db *sql.DB, userid int64) ([]Note, error) {
	// An user slice to hold data from returned rows.
	var notes []Note

	// query database for all notes with matching username
	rows, err := db.Query(`SELECT * FROM notes WHERE userid = $1`, userid)
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
		`UPDATE notes SET content = $1 WHERE id = $2 AND listid = $3 AND userid = $4;`,
		newnote.Content, newnote.Id, newnote.Listid, newnote.Userid); err != nil {
		return err
	}
	return nil
}

// add row in notes table
func addNote(db *sql.DB, newnote Note) error {
	if _, err := db.Exec(
		`INSERT INTO notes (userid, listid, content)
		 VALUES ($1, $2, $3);`, newnote.Userid, newnote.Listid, newnote.Content); err != nil {
		return err
	}
	return nil
}
