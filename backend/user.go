package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id       int32
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

	err := verifyJWT(w, r)
	if err != nil {
		log.Fatal(err)
	}

	if r.Method == http.MethodPost {

		var newuser User

		// read json payload into new user object
		if err := parseJSON(w, r, &newuser); err != nil {
			log.Fatal(err)
		}

		// hash user password
		if bytes, err := bcrypt.GenerateFromPassword([]byte(newuser.Password), 4); err != nil {
			log.Fatal(err)
		} else {
			newuser.Password = string(bytes)
		}

		// respond with jwt
		token, err := generateJWT(newuser)
		if err != nil {
			log.Fatal(err)
		}
		http.SetCookie(w, &http.Cookie{
			Name:  "token",
			Value: token,
		})

		return
	} else {

		http.Error(w, fmt.Sprintf("Expected method POST, got %v", r.Method), http.StatusMethodNotAllowed)

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

	verifyJWT(w, r)

	if r.Method == http.MethodPost {

		var login User

		if err := parseJSON(w, r, &login); err != nil {
			log.Fatal(err)
		}

		if err := verifyUser(h.db, login); err != nil {
			log.Fatal(err)
		}

		token, err := generateJWT(login)
		if err != nil {
			log.Fatal(err)
		}

		http.SetCookie(w, &http.Cookie{
			Name:  "token",
			Value: token,
		})

	} else {
		http.Error(w, fmt.Sprintf("Expected method POST, got %v", r.Method), http.StatusMethodNotAllowed)
		return
	}
}

// generate JWT from given user - returns err and token
func generateJWT(user User) (string, error) {

	// pull secret from environment
	secret := os.Getenv("JWTSECRET")
	if secret == "" {
		secret = "12043$521p8ijz4"
	}

	// generate new jwt
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	// add json payload
	claims["username"] = user.Username
	claims["exp"] = time.Now().Add(time.Minute * 2).Unix()

	// stringify token
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// checks if http request is authorized/logged in
func verifyJWT(w http.ResponseWriter, r *http.Request) error {

	// pull secret from environment
	secret := os.Getenv("JWTSECRET")
	if secret == "" {
		secret = "sampletoken"
	}

	// verify token header exists
	if r.Header["Token"] == nil {
		fmt.Fprintf(w, "Token not found in header")
		return errors.New("Missing auth token")
	}

	// parse and check token validity
	token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("There was an error in parsing")
		}
		return secret, nil
	})
	if err != nil || token == nil {
		fmt.Fprintf(w, "invalid token")
		return err
	}

	// parse claims from token
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		fmt.Fprintf(w, "couldn't parse claims")
		return errors.New("Token error")
	}

	// check if token is expired
	exp := claims["exp"].(float64)
	if int64(exp) < time.Now().Local().Unix() {
		fmt.Fprintf(w, "token expired")
		return errors.New("Token Expired")
	}

	return nil
}

func terminateJWT() {
	// replace jwt with another that expires immediately
}

// adds user object to the database - returns err
func newUser(db *sql.DB, newuser User) error {
	// insert new user object in database
	if _, err := db.Exec(
		`INSERT INTO users (username, password)
			VALUES ('` + newuser.Username + `', '` + newuser.Password + `');`); err != nil {
		return err
	}
	return nil
}

// checks if user credentials are valid - returns nil on success; err otherwise
func verifyUser(db *sql.DB, user User) error {

	var verified User

	// query database for user with matching username
	row, err := db.Query(`SELECT * FROM users WHERE id = ?`, user.Username)
	if err != nil {
		return err
	}
	defer row.Close()

	// scan row entry into user object
	if err := row.Scan(verified.Id, verified.Username, verified.Password); err != nil {
		return err
	}
	if err = row.Err(); err != nil {
		return err
	}

	// compare the user's stored password with the one provided
	if err := bcrypt.CompareHashAndPassword([]byte(verified.Password), []byte(user.Password)); err != nil {
		return err
	}

	return nil
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
