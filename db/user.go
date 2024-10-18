package db

import "go-session-storage-auth-test/models"

func SelectUserCountByEmail(email string, count *int) error {
	db, err := conn()
	if err != nil {
		return err
	}
	defer db.Close()

	err = db.QueryRow("SELECT COUNT(*) FROM users WHERE email = $1;", email).Scan(&count)
	if err != nil {
		return err
	}
	return nil
}

func InsertUser(user *models.User) error {
	db, err := conn()
	if err != nil {
		return err
	}
	defer db.Close()

	insertUserSQL := "INSERT INTO users (email, password) VALUES ($1, $2);"
	statement, err := db.Prepare(insertUserSQL)
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(user.Email, user.Password)
	if err != nil {
		return err
	}
	return nil
}
