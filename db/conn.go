package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func conn() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./test.db")

	if err != nil {
		fmt.Println("Error opening database:", err)
		return db, err
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("Error connecting to the database:", err)
		return db, err
	}

	return db, nil
}
