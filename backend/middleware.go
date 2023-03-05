package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func enableCors(w http.ResponseWriter) {
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Headers", "*")
	(w).Header().Set("Access-Control-Allow-Methods", "*")
	(w).Header().Set("Vary", "Origin")
	(w).Header().Set("Vary", "Access-Control-Request-Method")
	(w).Header().Set("Vary", "Access-Control-Request-Headers")
}

func renderJSON(w http.ResponseWriter, v interface{}) error {
	js, err := json.Marshal(v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return err
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
	return nil
}

func parseJSON(w http.ResponseWriter, r *http.Request, v interface{}) error {
	err := json.NewDecoder(r.Body).Decode(v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}
	return nil
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
	claims["userid"] = user.Id
	claims["exp"] = time.Now().Add(time.Minute * 2).Unix()

	// stringify token
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// checks if http request is authorized/logged in - returns error and username string; empty if err
func verifyJWT(r *http.Request) (int64, error) {

	// pull secret from environment
	secret := os.Getenv("JWTSECRET")
	if secret == "" {
		secret = "12043$521p8ijz4"
	}

	// verify token header exists
	if r.Header["Token"] == nil {
		log.Print("Token not found in header\n")
		return 0, errors.New("Missing auth token")
	}

	// parse and check token validity
	token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("There was an error in parsing")
		}
		return secret, nil
	})
	if err != nil || token == nil {
		log.Print("Invalid JWT Token\n")
		return 0, err
	}

	// parse claims from token
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		log.Print("Failed to parse JWT claims\n")
		return 0, errors.New("Token error")
	}

	// check if token is expired
	exp := claims["exp"].(float64)
	if int64(exp) < time.Now().Local().Unix() {
		log.Print("JWT is expired\n")
		return 0, errors.New("Token Expired")
	}

	userid := claims["userid"].(int64)
	log.Print("Valid request by user with id %i", userid)

	return userid, nil
}

func terminateJWT() {
	// replace jwt with another that expires immediately
}
