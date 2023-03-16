package main

import (
	"database/sql"
)

type List struct {
	Id     int64
	Userid int64
	Name   string
}

// query database for all lists corresponding to the given user - returns a slice of the lists and error
func getLists(db *sql.DB, userid int64) ([]List, error) {
	// An user slice to hold data from returned rows.
	var lists []List

	rows, err := db.Query(`SELECT * FROM lists WHERE userid = $1`, userid)
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
func addList(db *sql.DB, newlist List) error {
	if _, err := db.Exec(
		`INSERT INTO lists (userid, name)
			VALUES ($1, $2);`, newlist.Userid, newlist.Name); err != nil {
		return err
	}
	return nil
}

// query database for list attributes given list name - returns error
func getList(db *sql.DB, listname string, userid int64) (List, error) {
	var list List
	err := db.QueryRow(`SELECT * FROM lists WHERE name = $1 AND userid = $2`, listname, userid).Scan(&list.Id, &list.Userid, &list.Name)
	if err != nil {
		return list, err
	}
	return list, nil
}

// update row in lists table with corresponding listname and username - returns error
func updateList(db *sql.DB, updatedlist List, oldname string, userid int64) error {
	if _, err := db.Exec(
		`UPDATE lists SET name = $1 WHERE name = $2 AND userid = $3;`, updatedlist.Name, oldname, userid); err != nil {
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
