package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"time"

	_ "github.com/mattn/go-sqlite3"
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

func (t Todo) Update() error {
	db := GetConnection()
	q := `UPDATE todos set title=?, description=?, kind=?, updated_at=? WHERE id=?`

	stmt, err := db.Prepare(q)
	if err != nil {
		return err
	}
	defer stmt.Close()

	r, err := stmt.Exec(t.Title, t.Description, t.Kind, time.Now(), t.ID)
	if err != nil {
		return err
	}

	if i, err := r.RowsAffected(); err != nil || i != 1 {
		return errors.New("Error: se esperaba una fila afectada")
	}
	return nil
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

func GetTodosHandler(w http.ResponseWriter, r *http.Request) {
	t := new(Todo)

	todos, err := t.GetTodos()
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	j, err := json.Marshal(todos)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(j)

}

func CreateTodoHandler(w http.ResponseWriter, r *http.Request) {
	var todo Todo

	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = todo.create()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func UpdateTodoHandler(w http.ResponseWriter, r *http.Request) {
	var todo Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := todo.Update(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func DeleteTodoHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "Query id es requerido", http.StatusBadRequest)
		return
	}

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Query id debe ser un numero", http.StatusBadRequest)
		return
	}
	if err = delete(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func TodosHandlers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		GetTodosHandler(w, r)
	case http.MethodPost:
		CreateTodoHandler(w, r)
	case http.MethodPut:
		UpdateTodoHandler(w, r)
	case http.MethodDelete:
		DeleteTodoHandler(w, r)
	default:
		http.Error(w, "Metodo no permitido", http.StatusBadRequest)
		return
	}
}
