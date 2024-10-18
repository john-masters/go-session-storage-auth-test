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
		DROP TABLE IF EXISTS sessions;

		CREATE TABLE users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			email TEXT NOT NULL,
			password TEXT NOT NULL,
			is_member BOOLEAN NOT NULL DEFAULT FALSE,
		);

		CREATE TABLE sessions (
			id TEXT PRIMARY KEY,
			user_id INTEGER NOT NULL,
			FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
		);
	`)

	if err != nil {
		fmt.Println("Error creating table", err)
		return
	}

	fmt.Println("Successfully initialized database")
}
