package main

import (
	"fmt"
	"net/http"

	//	"context"
	//	"database/sql"
	//    "log"
	//    "os"
	"github.com/gorilla/mux"
	//	"github.com/cockroachdb/cockroach-go/crdb"
	// _ "github.com/lib/pq"
)

func main() {
	resp, err := http.Get("http://google.com/")
	if err != nil {
		// handle error
	}
	fmt.Println(resp)
	defer resp.Body.Close()
	//	body, err := io.ReadAll(resp.Body)
	router := mux.NewRouter()

	router.HandleFunc("/login", login).Methods("GET")

	fmt.Println("hello world")
}
