package db

import (
	"go-session-storage-auth-test/utils"
)

func InsertSession(userId int64) error {
	db, err := conn()
	if err != nil {
		return err
	}
	defer db.Close()

	hash, err := utils.GenerateRandomHash()
	if err != nil {
		return err
	}

	insertUserSQL := "INSERT INTO sessions (id, user_id) VALUES ($1, $2);"
	statement, err := db.Prepare(insertUserSQL)
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(hash, userId)
	if err != nil {
		return err
	}
	return nil
}
