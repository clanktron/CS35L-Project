package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
)

type List struct {
	Id     int64
	Userid int64
	Name   string
}

type ListHandler struct {
	db *sql.DB
}

func (h *ListHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	var list List
	// var err error

	enableCors(w)

	// list.Userid, err = verifyJWT(r)
	// if err != nil {
	// 	log.Print(err)
	// 	return
	// }

	if r.Method == http.MethodGet {

		list.Userid = 0

		lists, err := getLists(h.db, list.Userid)
		if err != nil {
			log.Print("Error getting lists from database")
			log.Print(err)
			return
		}
		if err := renderJSON(w, lists); err != nil {
			log.Print(err)
			return
		}

		return
	} else if r.Method == http.MethodPost {

		if err := parseJSON(w, r, &list); err != nil {
			log.Print(err)
			return
		}
		if err := addList(h.db, list); err != nil {
			log.Print(err)
			return
		}
		log.Print("Added list!\n")

		return
	} else if r.Method == http.MethodPut {
		// TODO: extract listname from url path
		oldlistname := "swagger eg"

		if err := parseJSON(w, r, &list); err != nil {
			log.Print("Error parsing json payload\n")
			log.Print(err)
			return
		}
		if err := updateList(h.db, list, oldlistname); err != nil {
			log.Printf("Error updating list %s\n", oldlistname)
			log.Print(err)
			return
		}
		log.Print("Updated List!\n")

		return
	} else if r.Method == http.MethodDelete {
		// TODO: extract listname from path
		listname := "swipeage"
		list.Name = listname
		list.Userid = 0

		if err := deleteList(h.db, list); err != nil {
			log.Print(err)
			log.Print("Failed to delete list\n")
		}
		log.Printf("Deleted list %s\n", listname)

		return
	} else if r.Method == http.MethodOptions {
		return
	} else {
		http.Error(w, fmt.Sprintf("Expected method GET, POST, PUT, OPTIONS, or DELETE, got %v", r.Method), http.StatusMethodNotAllowed)
		return
	}
}

// Add row in lists table with corresponding userid and name
func addList(db *sql.DB, newlist List) error {
	if _, err := db.Exec(
		`INSERT INTO lists (userid, name)
			VALUES ($1, $2);`, newlist.Userid, newlist.Name); err != nil {
		return err
	}
	return nil
}

// query database for all lists corresponding to the given user - returns a slice of the lists and error
func getLists(db *sql.DB, userid int64) ([]List, error) {
	// An user slice to hold data from returned rows.
	var lists []List

	rows, err := db.Query(`SELECT * FROM lists WHERE userid = ?`, userid)
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
		`UPDATE lists SET name = $1 WHERE name = $2 AND userid = $3;`, newlist.Name, oldlistname, newlist.Id); err != nil {
		return err
	}
	return nil
}

// delete corresponding row in lists table
func deleteList(db *sql.DB, list List) error {
	if _, err := db.Exec(
		`DELETE FROM lists WHERE name = $1 AND userid = $2;`, list.Name, list.Userid); err != nil {
		return err
	}
	return nil
}
