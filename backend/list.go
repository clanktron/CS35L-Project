package main

import (
	"database/sql"
	"fmt"
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

func GetLists(db *sql.DB, userid int32) ([]List, error) {
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

// Add row in lists table with corresponding userid and name
func AddList(db *sql.DB, newlist List) error {
	if _, err := db.Exec(
		`INSERT INTO lists (userid, name)
		 VALUES (` + string(newlist.Userid) + `, ` + string(newlist.Name) + `);`); err != nil {
		return err
	}
	return nil
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
