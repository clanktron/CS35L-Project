package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/crypto/bcrypt"
)

type LoginHandler struct {
	db *sql.DB
}

func (h *LoginHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// check auth function
	// verifyJWT()
	if r.Method == http.MethodPost {
		var login User
		parseJSON(w, r, &login)
		// TODO: query database for user with matching username and password
		// rows, err := h.db.Query(`SELECT * FROM users WHERE username  = ?`, login.Username)
		//
		// TODO: respond with jwt token derived from username (include userid?)
	} else {
		http.Error(w, fmt.Sprintf("Expected method POST, got %v", r.Method), http.StatusMethodNotAllowed)
		return
	}
}

type User struct {
	Id       int32
	Username string
	Password string
}

type UserHandler struct {
	db *sql.DB
}

var sampleSecretKey = []byte("SecretYouShouldHide")

type Claims struct {
}

func generateJWT() (string, error) {
	// maybe put this in main()?
	secret := os.Getenv("JWTSECRET")
	if secret == "" {
		secret = "12043$521p8ijz4"
	}
	// generate jwt from header, payload, and secret
	return "thisisadummytoken", nil
}

func verifyJWT(jwt string) {
	// take jwt and
	return
}

func terminateJWT() {
	// replace jwt with another that expires immediately
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 4)
	return string(bytes), err
}

func checkPasswordHash(password string, hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err
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

func (h *UserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// create new user object and store username and password
		var newuser User
		err := parseJSON(w, r, &newuser)
		if err != nil {
			log.Fatal(err)
		}
		// fmt.Print(newuser)
		// hash user password
		newuser.Password, err = hashPassword(newuser.Password)
		return
	} else {
		http.Error(w, fmt.Sprintf("Expected method POST, got %v", r.Method), http.StatusMethodNotAllowed)
		return
	}

	// else if r.Method == http.MethodDelete {
	// 	var newuser User
	// 	parseJSON(w, r, &newuser)
	// 	// insert new user object in database

	// 	return
	// }

}
