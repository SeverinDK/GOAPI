package models

import (
	"database/sql"
	"errors"
)

type User struct {
	ID       int64  `json:id`
	Username string `json:username`
}

// GetByKey returns a user matching provided identification key
func GetByKey(db *sql.DB, id int64) (*User, error) {
	row := db.QueryRow("SELECT * FROM users WHERE id = ?;", id)
	user, err := mapRowToUser(row)

	if err != nil {
		return nil, err
	}

	return user, nil
}

// GetAll gets all users in the database
func GetAll(db *sql.DB) ([]*User, error) {
	rows, err := db.Query("SELECT * FROM users;")

	if err != nil {
		return nil, err
	}

	users, err := mapRowsToUsers(rows)

	if err != nil {
		return nil, err
	}

	return users, nil
}

// Create creates a user
func Create(db *sql.DB, username string) (*User, error) {
	exists, err := UsernameExists(db, username)

	if err != nil {
		return nil, err
	}

	if exists == true {
		return nil, errors.New("Username is already in use")
	}

	result, err := db.Exec("INSERT INTO users(username) values(?)", username)

	if err != nil {
		return nil, err
	}

	id, _ := result.LastInsertId()

	user, err := GetByKey(db, id)

	if err != nil {
		return nil, err
	}

	return user, nil
}

// UsernameExists if a given username already exists in the database
func UsernameExists(db *sql.DB, username string) (bool, error) {
	exists := false
	err := db.QueryRow("SELECT EXISTS (SELECT TRUE FROM users WHERE username = ?);", username).Scan(&exists)

	if err != nil {
		return false, err
	}

	return exists, nil
}

func mapRowToUser(row *sql.Row) (*User, error) {
	user := User{}
	err := row.Scan(&user.ID, &user.Username)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func mapRowsToUsers(rows *sql.Rows) ([]*User, error) {
	defer rows.Close()

	users := make([]*User, 0)

	for rows.Next() {
		user := User{}
		err := rows.Scan(&user.ID, &user.Username)

		if err != nil {
			return nil, err
		}

		users = append(users, &user)
	}

	return users, nil
}
