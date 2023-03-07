package main

import (
	"database/sql"
)

type Note struct {
	Id      int64
	Userid  int64
	Listid  int64
	Content string
}

// delete row in notes table corresponding to noteid
func deleteNote(db *sql.DB, note Note) error {
	if _, err := db.Exec(
		`DELETE	FROM notes WHERE id = $1;`, note.Id); err != nil {
		return err
	}
	return nil
}

// query database for full note object given noteid - returns error
func getNote(db *sql.DB, noteid int64) (Note, error) {
	var note Note
	err := db.QueryRow(`SELECT * FROM notes WHERE id = $1`, noteid).Scan(&note.Id, &note.Userid, &note.Listid, &note.Content)
	if err != nil {
		return note, err
	}
	return note, nil
}

// queries database for all notes corresponding to user - returns a Notes slice and error
func getNotes(db *sql.DB, note Note) ([]Note, error) {
	// An user slice to hold data from returned rows.
	var notes []Note

	// query database for all notes with matching username
	rows, err := db.Query(`SELECT * FROM notes WHERE userid = $1 AND listid = $2`, note.Userid, note.Listid)
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

// add row in notes table corresponding to given userid, listid, and content
func addNote(db *sql.DB, newnote Note) error {
	if _, err := db.Exec(
		`INSERT INTO notes (userid, listid, content)
		 VALUES ($1, $2, $3);`, newnote.Userid, newnote.Listid, newnote.Content); err != nil {
		return err
	}
	return nil
}
