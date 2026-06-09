package main

import (
	"errors"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func (n Note) create() error {
	db := GetConnection()

	q := `INSERT INTO notes (title, description, updated_at) VALUES(?,?,?)`

	//preparamos petición para insertar de manera segura
	stmt, err := db.Prepare(q)
	if err != nil {
		return err
	}
	defer stmt.Close()

	r, err := stmt.Exec(n.Title, n.Despcrition, time.Now())
	if err != nil {
		return err
	}

	//Confirmamos si una fila fue afectada, debido al insert
	//de lo contrario enviamos error
	if i, err := r.RowsAffected(); err != nil || i != 1 {
		return errors.New("Error: se esperaba una fila afectada")
	}
	return nil
}

func (n *Note) GetNotes() ([]Note, error) {
	db := GetConnection()
	q := `SELECT
			id, title, description, created_at, updated_at
			FROM notes`

	rows, err := db.Query(q)
	if err != nil {
		return []Note{}, err
	}
	defer rows.Close()

	notes := []Note{}

	for rows.Next() {
		rows.Scan(
			&n.ID,
			&n.Title,
			&n.Despcrition,
			&n.CreatedAt,
			&n.UpdatedAt,
		)
		notes = append(notes, *n)
	}
	return notes, nil
}

func (n Note) Update() error {
	db := GetConnection()
	q := `UPDATE notes set title=?, description=?, updated_at=? 
			WHERE id=?`

	stmt, err := db.Prepare(q)
	if err != nil {
		return err
	}
	defer stmt.Close()

	r, err := stmt.Exec(n.Title, n.Despcrition, time.Now(), n.ID)
	if err != nil {
		return err
	}

	if i, err := r.RowsAffected(); err != nil || i != 1 {
		return errors.New("Error: se esperaba una fila afectada")
	}
	return nil
}

func (n Note) delete(id int64) error {
	db := GetConnection()
	q := `DELETE FROM notes WHERE id=?`

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
