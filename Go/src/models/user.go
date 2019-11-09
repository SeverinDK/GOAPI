package models

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
)

type User struct {
	ID       int64  `json:id`
	Username string `json:username`
}

func mapRowToUser(row *sql.Row) *User {
	user := User{}
	err := row.Scan(&user.ID, &user.Username)

	if err != nil {
		return nil
	}

	return &user
}

func mapRowsToUsers(rows *sql.Rows) []*User {
	defer rows.Close()

	users := make([]*User, 0)

	for rows.Next() {
		user := User{}
		err := rows.Scan(&user.ID, &user.Username)

		if err != nil {
			fmt.Println(err)
			continue
		}

		users = append(users, &user)
	}

	return users
}

// GetByKey returns a user matching provided identification key
func GetByKey(db *sql.DB, id int64) *User {
	row := db.QueryRow("SELECT * FROM users WHERE id = ?;", id)

	return mapRowToUser(row)
}

// GetAll gets all users in the database
func GetAll(db *sql.DB) []*User {
	rows, err := db.Query("SELECT * FROM users;")

	if err != nil {
		log.Fatal(err)
	}

	return mapRowsToUsers(rows)
}

// Create creates a user
func Create(db *sql.DB, username string) (*User, error) {
	exists := Exists(db, username)

	if exists == true {
		return nil, errors.New("Username is already in use")
	}

	result, err := db.Exec("INSERT INTO users(username) values(?)", username)

	if err != nil {
		return nil, err
	}

	id, _ := result.LastInsertId()

	return GetByKey(db, id), nil
}

// Exists checks if a username already exists in the database
func Exists(db *sql.DB, username string) bool {
	err := db.QueryRow("SELECT TRUE AS exists FROM users WHERE username = ?", username).Scan(&username)

	if err != nil {
		return false
	}

	return true
}
