package db

import (
	"go-session-storage-auth-test/models"
)

func SelectUserCountByEmail(email string) (int, error) {
	db, err := conn()
	if err != nil {
		return 0, err
	}
	defer db.Close()

	var count int

	err = db.QueryRow("SELECT COUNT(*) FROM users WHERE email = $1;", email).Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func InsertUser(user *models.User) (userId int64, error error) {
	db, err := conn()
	if err != nil {
		return 0, err
	}
	defer db.Close()

	insertUserSQL := "INSERT INTO users (email, password) VALUES ($1, $2);"
	statement, err := db.Prepare(insertUserSQL)
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(user.Email, user.Password)
	if err != nil {
		return 0, err
	}

	userId, err = result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return userId, nil
}
