package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
)

type List struct {
	Id     int32
	Userid int32
	Name   string
}

type ListHandler struct {
	db *sql.DB
}

func (h *ListHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	var list List
	var err error

	list.Userid, err = verifyJWT(w, r)
	if err != nil {
		log.Print(err)
	}

	// respond with all lists corresponding to userid
	if r.Method == http.MethodGet {

		lists, err := getLists(h.db, list.Userid)
		if err != nil {
			log.Print("Error getting lists from database")
		}
		renderJSON(w, lists)

		return
	} else if r.Method == http.MethodPost {

		parseJSON(w, r, &list)
		addList(h.db, list)

		return
	} else if r.Method == http.MethodPut {
		// TODO: extract listname from url path
		var oldlistname string

		parseJSON(w, r, &list)
		updateList(h.db, list, oldlistname)

		return
	} else if r.Method == http.MethodDelete {
		// TODO: extract listname from path
		var listname string
		list.Name = listname

		deleteList(h.db, list)

		return
	} else {
		http.Error(w, fmt.Sprintf("Expected method GET, POST, PUT, or DELETE, got %v", r.Method), http.StatusMethodNotAllowed)
		return
	}
}

// Add row in lists table with corresponding userid and name
func addList(db *sql.DB, newlist List) error {
	if _, err := db.Exec(
		`INSERT INTO lists (userid, name)
		 VALUES (` + string(newlist.Userid) + `, ` + string(newlist.Name) + `);`); err != nil {
		return err
	}
	return nil
}

// query database for all lists corresponding to the given user - returns a slice of the lists and error
func getLists(db *sql.DB, userid int32) ([]List, error) {
	// An user slice to hold data from returned rows.
	var lists []List

	rows, err := db.Query(`SELECT * FROM lists WHERE userid = ?`, string(userid))
	if err != nil {
		return lists, err
	}
	defer rows.Close()

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

// update row in lists table with corresponding listname and username - returns error
func updateList(db *sql.DB, newlist List, oldlistname string) error {
	if _, err := db.Exec(
		// TODO: need to match userid as well
		`UPDATE lists SET name = '` + newlist.Name + `' WHERE name = '` + oldlistname + `';`); err != nil {
		log.Print(err)
		return err
	}
	return nil
}

// delete corresponding row in lists table
func deleteList(db *sql.DB, list List) error {
	if _, err := db.Exec(
		// TODO: need to match userid as well
		`DELETE FROM lists WHERE name = '` + list.Name + `';`); err != nil {
		log.Print(err)
		return err
	}
	return nil
}
