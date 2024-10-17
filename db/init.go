package db

import (
	"fmt"
)

func Init() {
	db, err := conn()
	if err != nil {
		fmt.Println("Error initializing database")
		return
	}
	defer db.Close()
	_, err = db.Exec(`
		DROP TABLE IF EXISTS users;

		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			email TEXT NOT NULL,
			password TEXT NOT NULL,
			is_member BOOLEAN NOT NULL DEFAULT FALSE,
		);
	`)

	if err != nil {
		fmt.Println("Error creating table", err)
		return
	}

	fmt.Println("Successfully initialized database")
}
