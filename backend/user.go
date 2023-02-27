package main

import (
	"database/sql"
)

type User struct {
	Id       int32
	Username string
	Password string
}

func newUser(db *sql.DB, newuser User) error {
	// insert new user object in database
	if _, err := db.Exec(
		`INSERT INTO users (username, password)
			VALUES ('` + newuser.Username + `', '` + newuser.Password + `');`); err != nil {
		return err
	}
	return nil
}
