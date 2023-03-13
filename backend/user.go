package main

import (
	"database/sql"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id       int64
	Username string
	Password string
}

// adds user object to the database - returns err
func initAdmin(db *sql.DB, admin User) error {
	// hash user password
	if bytes, err := bcrypt.GenerateFromPassword([]byte(admin.Password), 4); err != nil {
		log.Print(err)
	} else {
		admin.Password = string(bytes)
	}
	// insert new user object in database
	if _, err := db.Exec(
		`UPSERT INTO users (id, username, password)
			VALUES (0, '` + admin.Username + `', '` + admin.Password + `');`); err != nil {
		return err
	}
	// insert default list for admin in database
	if _, err := db.Exec(
		`UPSERT INTO lists (id, userid, name)
			VALUES (0, 0, $1);`, "my tasks"); err != nil {
		return err
	}
	return nil
}

func addUser(db *sql.DB, newuser User) error {
	// hash user password
	if bytes, err := bcrypt.GenerateFromPassword([]byte(newuser.Password), 4); err != nil {
		return err
	} else {
		newuser.Password = string(bytes)
	}
	// insert new user object in database
	if _, err := db.Exec(
		`INSERT INTO users (username, password)
			VALUES ('` + newuser.Username + `', '` + newuser.Password + `');`); err != nil {
		return err
	}
	return nil
}

// returns error and full user object for given username
func getUser(db *sql.DB, username string) (User, error) {
	var user User
	err := db.QueryRow(`SELECT * FROM users WHERE username = $1`, username).Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {
		return user, err
	}
	return user, nil
}

// queries database for all users - returns err and a slice of all users; empty if err
func getUsers(db *sql.DB) ([]User, error) {

	var users []User

	rows, err := db.Query(`SELECT * FROM users`)
	if err != nil {
		users = nil
		return users, err
	}
	defer rows.Close()

	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var usr User
		if err := rows.Scan(&usr.Id, &usr.Username, &usr.Password); err != nil {
			users = nil
			return users, err
		}
		users = append(users, usr)
	}
	if err = rows.Err(); err != nil {
		users = nil
		return users, err
	}

	return users, nil
}

// checks if user credentials are valid - returns nil on success; err otherwise
func verifyUser(db *sql.DB, user User) error {

	verified, err := getUser(db, user.Username)
	if err != nil {
		return err
	}

	// compare the user's stored password with the one provided
	if err := bcrypt.CompareHashAndPassword([]byte(verified.Password), []byte(user.Password)); err != nil {
		return err
	}

	return nil
}

// delete corresponding row in users table
func deleteUser(db *sql.DB, userid int64) error {
	if _, err := db.Exec(
		`DELETE FROM users WHERE id = $1;`, userid); err != nil {
		return err
	}
	return nil
}
