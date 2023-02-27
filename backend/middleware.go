package main

import (
	"encoding/json"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

var sampleSecretKey = []byte("SecretYouShouldHide")

func generateJWT() (string, error) {
	// maybe put this in main()?
	// secret := os.Getenv("JWTSECRET")
	// if secret == "" {
	// 	secret = "12043$521p8ijz4"
	// }
	// generate jwt from userid, secret, and
	return "thisisadummytoken", nil
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

func jsonMayhaps(w http.ResponseWriter, r *http.Request) {
	rawData := []Note{
		{Id: 0, Userid: 1, Listid: 1, Content: "Do your math 61 hwk."},
		{Id: 1, Userid: 1, Listid: 1, Content: "This is another test."},
	}

	data, err := json.Marshal(rawData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(data)
}
