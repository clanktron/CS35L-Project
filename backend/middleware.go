package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func enableCors(w http.ResponseWriter) {
	(w).Header().Set("Access-Control-Allow-Headers", "Accept, Accept-Charset, Accept-Language, Access-Control-Request-Method, Access-Control-Request-Headers, Authorization, Content-Length, Content-Type, Cookie, Date, Forwarded, Origin, User-Agent")
	(w).Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	(w).Header().Set("Access-Control-Allow-Credentials", "true")
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
		secret = "cZydEgWlyBB9HwucTwyExcaYEX1g77q2RsrqlnT1YURvsoNfvWUvLwdS6q58SewMrue1kuMrtl3ahkD6LAvpFlKVBD5oDpariBK2MUcV3HD9f2lfAmThIihvvXi2lho02qgeSEKFuB2Slk5EZkVy17FxEzbJOUywLr0jfmP+aKzLv4TiuHaHbOiVIPefUQHioY1beAnUNPaRFELyv4675uDMWJTpjlXbA6K7+w=="
	}

	// generate new jwt
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	// add claims payload
	claims["exp"] = time.Now().Add(time.Minute * 60).Unix()
	claims["userid"] = fmt.Sprint(user.Id)

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
		secret = "cZydEgWlyBB9HwucTwyExcaYEX1g77q2RsrqlnT1YURvsoNfvWUvLwdS6q58SewMrue1kuMrtl3ahkD6LAvpFlKVBD5oDpariBK2MUcV3HD9f2lfAmThIihvvXi2lho02qgeSEKFuB2Slk5EZkVy17FxEzbJOUywLr0jfmP+aKzLv4TiuHaHbOiVIPefUQHioY1beAnUNPaRFELyv4675uDMWJTpjlXbA6K7+w=="
	}

	// get token from cookie
	c, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			// If the cookie is not set, return an unauthorized status
			return 0, err
		}
		// For any other type of error, return a bad request status
		return 0, err
	}

	// Get the JWT string from the cookie
	tokenstring := c.Value

	// parse and check token validity
	token, err := jwt.Parse(tokenstring, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return "", errors.New("Invalid JWT Token")
		}
		return []byte(secret), nil
	})
	if err != nil {
		return 0, err
	}

	// parse claims from token
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("Failed to parse JWT claims")
	}

	// check if token is expired
	exp := claims["exp"].(float64)
	if int64(exp) < time.Now().Local().Unix() {
		return 0, errors.New("Token Expired")
	}

	s_userid := claims["userid"].(string)
	log.Printf("Valid request by user with id %s", s_userid)
	userid, err := strconv.ParseInt(s_userid, 10, 64)

	return userid, nil
}

func terminateJWT() {
	// replace jwt with another that expires immediately
}
