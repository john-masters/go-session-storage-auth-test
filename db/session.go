package db

import (
	"go-session-storage-auth-test/utils"
)

func SelectSessionCountByUserId(userId int64, count *int) error {
	db, err := conn()
	if err != nil {
		return err
	}
	defer db.Close()

	err = db.QueryRow("SELECT COUNT(*) FROM sessions WHERE user_id = $1;", userId).Scan(&count)
	if err != nil {
		return err
	}
	return nil
}

func InsertSession(userId int64) (sessionId string, error error) {
	db, err := conn()
	if err != nil {
		return "", err
	}
	defer db.Close()

	hash, err := utils.GenerateRandomHash()
	if err != nil {
		return "", err
	}

	insertUserSQL := "INSERT INTO sessions (id, user_id) VALUES ($1, $2);"
	statement, err := db.Prepare(insertUserSQL)
	if err != nil {
		return "", err
	}
	defer statement.Close()

	_, err = statement.Exec(hash, userId)
	if err != nil {
		return "", err
	}
	return hash, nil
}

func DeleteSessionByUserId(userId int64) error {
	db, err := conn()
	if err != nil {
		return err
	}
	defer db.Close()

	deleteSessionSQL := "DELETE FROM sessions WHERE user_id = $1;"
	statement, err := db.Prepare(deleteSessionSQL)
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(userId)
	if err != nil {
		return err
	}
	return nil
}
