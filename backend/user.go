package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id       int64
	Username string
	Password string
}

type UserHandler struct {
	db *sql.DB
}

type LoginHandler struct {
	db *sql.DB
}

// accepts POST requests with new user payloads - responds with jwt or error
func (h *UserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	enableCors(w)

	// _, err := verifyJWT(r)
	// if err != nil {
	// 	log.Print(err)
	// }

	if r.Method == http.MethodPost {

		var newuser User
		var err error

		// read json payload into new user object
		if err := parseJSON(w, r, &newuser); err != nil {
			log.Print(err)
		}

		// check contents of payload
		fmt.Print(newuser, "\n")

		// add user to database
		addUser(h.db, newuser)

		// reassign stored values in user object
		newuser, err = getUser(h.db, newuser.Username)
		if err != nil {
			log.Print(err)
		}

		// respond with jwt
		token, err := generateJWT(newuser)
		if err != nil {
			log.Print(err)
		}

		fmt.Print(token, "\n")

		ck := http.Cookie{
			Name:  "token",
			Value: token,
		}

		http.SetCookie(w, &ck)

		return
	} else if r.Method == http.MethodOptions {
		return
	} else {
		// http.Error(w, fmt.Sprintf("Expected method POST, got %v", r.Method), http.StatusMethodNotAllowed)
		fmt.Printf("Expected method POST, got %v", r.Method)
		return
	}

	// else if r.Method == http.MethodDelete {
	// 	parse username from jwt token
	// 	remove row with corresponding username
	// 	return

	// }

}

// accepts POST requests with user credentials - responds with a jwt or error
func (h *LoginHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	enableCors(w)

	// _, err := verifyJWT(r)
	// if err != nil {
	// 	log.Print(err)
	// }

	if r.Method == http.MethodPost {

		var login User

		if err := parseJSON(w, r, &login); err != nil {
			log.Print(err)
		}

		if err := verifyUser(h.db, login); err != nil {
			log.Print(err)
		}

		token, err := generateJWT(login)
		if err != nil {
			log.Print(err)
		}

		http.SetCookie(w, &http.Cookie{
			Name:  "token",
			Value: token,
		})

	} else if r.Method == http.MethodOptions {
		return
	} else {
		http.Error(w, fmt.Sprintf("Expected method POST, got %v", r.Method), http.StatusMethodNotAllowed)
		return
	}
}

// adds user object to the database - returns err
func addUser(db *sql.DB, newuser User) error {
	// hash user password
	if bytes, err := bcrypt.GenerateFromPassword([]byte(newuser.Password), 4); err != nil {
		log.Print(err)
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

// returns user object corresponding to given username
func getUser(db *sql.DB, username string) (User, error) {

	var user User

	// query database for user with matching username
	row, err := db.Query(`SELECT * FROM users WHERE id = ?`, username)
	if err != nil {
		return user, err
	}
	defer row.Close()

	// scan row entry into user object
	if err := row.Scan(user.Id, user.Username, user.Password); err != nil {
		return user, err
	}
	if err = row.Err(); err != nil {
		return user, err
	}

	return user, err
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
