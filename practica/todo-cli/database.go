package main

import (
	"errors"
	"time"
)

func (t Todo) create() error {
	db := GetConnection()

	q := `INSERT INTO todos (title, description, kind, updated_at) VALUES(?,?,?,?)`

	stmt, err := db.Prepare(q)
	if err != nil {
		return err
	}
	defer stmt.Close()

	r, err := stmt.Exec(t.Title, t.Description, t.Kind, time.Now())
	if err != nil {
		return err
	}

	if i, err := r.RowsAffected(); err != nil || i != 1 {
		return errors.New("Error: se esperaba una fila afectada")
	}

	return nil
}

func (t *Todo) GetTodos() ([]Todo, error) {
	db := GetConnection()

	q := `SELECT id, title, description, kind, created_at, updated_at FROM todos`

	rows, err := db.Query(q)

	if err != nil {
		return []Todo{}, err
	}
	defer rows.Close()

	var todos []Todo

	for rows.Next() {
		rows.Scan(
			&t.ID,
			&t.Title,
			&t.Description,
			&t.Kind,
			&t.CreatedAt,
			&t.UpdatedAt,
		)
		todos = append(todos, *t)
	}
	return todos, nil
}

func delete(id int64) error {
	db := GetConnection()
	q := `DELETE FROM todos WHERE ID=?`
	stmt, err := db.Prepare(q)
	if err != nil {
		return err
	}
	defer stmt.Close()

	r, err := stmt.Exec(id)
	if err != nil {
		return err
	}

	if i, err := r.RowsAffected(); err != nil || i != 1 {
		return errors.New("Error: se esperaba una fila afectada")
	}
	return nil
}
