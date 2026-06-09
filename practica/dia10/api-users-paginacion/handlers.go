package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

func (u User) create() error {
	db := GetConnection()
	q := `INSERT INTO users (name, email, pin, status) VALUES(?,?,?,?)`

	stmt, err := db.Prepare(q)
	if err != nil {
		return err
	}
	defer stmt.Close()

	r, err := stmt.Exec(u.Name, u.Email, u.Pin, u.Status)
	if err != nil {
		return err
	}

	if i, err := r.RowsAffected(); err != nil || i != 1 {
		return errors.New("Error: se esperaba una fila afectada")
	}

	return nil
}

func (u *User) GetUsers(page, limit int) ([]User, error, int) {
	db := GetConnection()

	offset := (page - 1) * limit
	var totalRows int

	qTotalRows := `SELECT COUNT(*) FROM users`
	q := `SELECT id, name, email, pin, status FROM users LIMIT ? OFFSET ?`

	if err := db.QueryRow(qTotalRows).Scan(&totalRows); err != nil {
		return []User{}, err, 0
	}

	rows, err := db.Query(q, limit, offset)
	if err != nil {
		return []User{}, err, 0
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
	return users, nil, totalRows
}

func (u User) Update() error {
	db := GetConnection()
	q := `UPDATE users set name=?, email=?, pin=?, status=? where id=?`
	stmt, err := db.Prepare(q)
	if err != nil {
		return err
	}
	defer stmt.Close()

	r, err := stmt.Exec(u.Name, u.Email, u.Pin, u.Status, u.ID)
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
	q := `DELETE FROM users WHERE id=?`
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

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	u := new(User)
	query := r.URL.Query()
	pageStr := query.Get("page")
	limitStr := query.Get("limit")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 {
		limit = 10
	}

	users, err, totalRows := u.GetUsers(page, limit)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	//Construir la respuesta con meta datos
	totalPages := (totalRows + limit - 1) / limit
	response := PaginatedResponse{
		Data:       users,
		TotalRows:  totalRows,
		TotalPages: totalPages,
		Page:       page,
		Limit:      limit,
	}
	w.Header().Set("Content-Type", "applitation/json")
	json.NewEncoder(w).Encode(response)

	/*
		j, err := json.Marshal(users)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write(j)
	*/
}

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := user.create(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := user.Update(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
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

func UsersHandlers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		GetUsersHandler(w, r)
	case http.MethodPost:
		CreateUserHandler(w, r)
	case http.MethodPut:
		UpdateUserHandler(w, r)
	case http.MethodDelete:
		DeleteUserHandler(w, r)
	default:
		http.Error(w, "Metodo no permitido", http.StatusBadRequest)
		return
	}
}
