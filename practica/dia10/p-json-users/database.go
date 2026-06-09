package main

import _ "github.com/mattn/go-sqlite3"

func (u *User) GetUsers() ([]User, error) {
	db := GetConnection()

	q := `SELECT id, name, email, pin, status FROM users`

	rows, err := db.Query(q)
	if err != nil {
		return []User{}, err
	}
	defer rows.Close()

	users := []User{}

	for rows.Next() {
		rows.Scan(
			&u.ID,
			&u.Name,
			&u.Email,
			&u.Pin,
			&u.Status,
		)
		users = append(users, *u)
	}
	return users, nil
}
