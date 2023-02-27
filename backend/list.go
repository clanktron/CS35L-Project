package main

import (
	"database/sql"
)

type List struct {
	Id     int32
	Userid int32
	Name   string
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
