package main

import (
	"database/sql"
)

func databaseInit(dbuser string, dbpass string, dburl string) (*sql.DB, error) {

	// Initialize database connection
	var connStr string
	if dbpass == "" {
		connStr = "postgresql://" + dbuser + "@" + dburl + "/defaultdb?sslmode=disable"
	} else {
		connStr = "postgresql://" + dbuser + ":" + dbpass + "@" + dburl + "/defaultdb?sslmode=disable"
	}

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return db, err
	} else if err = db.Ping(); err != nil {
		return db, err
	}

	// Create the "users" table.
	// Note: Possibly change SERIAL (postgres & cockroach) to UUID (cockroach only)
	if _, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY, 
			username STRING NOT NULL UNIQUE, 
			password STRING NOT NULL 
		)`); err != nil {
		return db, err
	}

	// Create the "list" table.
	if _, err := db.Exec(
		`CREATE TABLE IF NOT EXISTS list (
	           id SERIAL PRIMARY KEY,
	           userid INT NOT NULL REFERENCES users(Id) ON UPDATE CASCADE ON DELETE CASCADE,
	           name STRING NOT NULL
		   )`); err != nil {
		return db, err
	}

	// Create the "note" table.
	if _, err := db.Exec(
		`CREATE TABLE IF NOT EXISTS note (
	           id SERIAL PRIMARY KEY,
	           userid INT NOT NULL REFERENCES users(Id) ON UPDATE CASCADE ON DELETE CASCADE,
	           listid INT NOT NULL REFERENCES list(Id) ON UPDATE CASCADE ON DELETE CASCADE,
	           content VARCHAR(280)
	           )`); err != nil {
		return db, err
	}

	// Ensure admin user is in "users" table.
	// if _, err := db.Exec(
	// 	`INSERT INTO users (id, username, password)
	// 	 VALUES (0, 'admin', 'admin');`); err != nil {
	// 	log.Fatal(err)
	// }
	return db, err
}
