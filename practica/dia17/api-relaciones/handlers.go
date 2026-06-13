package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/jackc/pgx/v5"
)

func (t Task) create() error {
	db, ctx := GetConnection()
	q := `INSERT INTO tasks(user_id, title, content, kind) VALUES($1, $2, $3, $4)`
	if _, err := db.Exec(ctx, q, t.UserId, t.Title, t.Content, t.Kind); err != nil {
		return err
	}
	return nil
}

func getUserWithTasks(userId int64) (*User, error) {
	db, ctx := GetConnection()
	q := `SELECT u.id, u.name, u.email, u.is_active, u.is_admin, t.id ,t.user_id, t.title, t.content FROM users u LEFT JOIN tasks t ON u.id = t.user_id WHERE u.id = $1`
	rows, err := db.Query(ctx, q, userId)
	if err != nil {
		fmt.Println("error en query")
		return nil, err
	}
	defer rows.Close()

	var user *User

	for rows.Next() {
		var u User
		//var t sql.NullInt32
		var taskID, taskUserID sql.NullInt64
		var postTitle, postContent sql.NullString

		if err = rows.Scan(&u.ID, &u.Name, &u.Email, &u.IsActive, &u.IsAdmin, &taskID, &taskUserID, &postTitle, &postContent); err != nil {
			fmt.Println("error asignado datos a variables")
			return nil, err
		}

		if user == nil {
			user = &u
		}

		if taskUserID.Valid {
			task := Task{
				ID:      int64(taskID.Int64),
				UserId:  int64(taskUserID.Int64),
				Title:   postTitle.String,
				Content: postContent.String,
			}
			user.Tasks = append(user.Tasks, task)
		}
	}
	return user, nil
}

func (t *Task) getAll(page, limit int) ([]Task, int, error) {
	db, ctx := GetConnection()
	offset := (page - 1) * limit
	var totalRows int
	qTotalRows := `SELECT COUNT(*) FROM tasks`
	if err := db.QueryRow(ctx, qTotalRows).Scan(&totalRows); err != nil {
		return []Task{}, 0, err
	}
	q := `SELECT id,user_id, title, content, kind, created_at FROM tasks LIMIT $1 OFFSET $2`

	rows, err := db.Query(ctx, q, limit, offset)
	if err != nil {
		return []Task{}, 0, err
	}
	defer rows.Close()

	tasks, err := pgx.CollectRows(rows, pgx.RowToStructByName[Task])
	if err != nil {
		return []Task{}, 0, err
	}

	return tasks, totalRows, nil
}

func getUserWithTasksHandler(w http.ResponseWriter, r *http.Request) {
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
	fmt.Printf("ID: %d", id)

	user, err := getUserWithTasks(id)
	if err != nil {
		http.Error(w, "No se encontro usuario", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "applitation/json")
	json.NewEncoder(w).Encode(user)
}

func GetTasksHandler(w http.ResponseWriter, r *http.Request) {
	t := new(Task)
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

	tasks, totalRows, err := t.getAll(page, limit)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	totalPages := (totalRows + limit - 1) / limit
	response := PaginatedTaskResponse{
		Data:       tasks,
		TotalRows:  totalRows,
		TotalPages: totalPages,
		Page:       page,
		Limit:      limit,
	}
	w.Header().Set("Content-Type", "applitation/json")
	json.NewEncoder(w).Encode(response)
}

func CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task Task

	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := task.create(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (u User) create() error {
	db, ctx := GetConnection()
	q := `INSERT INTO users(name, email) VALUES($1, $2)`
	if _, err := db.Exec(ctx, q, u.Name, u.Email); err != nil {
		return err
	}
	return nil
}

func getUserById(id int64) (User, error) {
	db, ctx := GetConnection()
	q := `SELECT id, name, email, is_active, is_admin FROM users where id=$1`
	row, err := db.Query(ctx, q, id)
	if err != nil {
		return User{}, nil
	}
	defer row.Close()

	user, err := pgx.CollectOneRow(row, pgx.RowToStructByName[User])
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func (u *User) getAll(page int, limit int) ([]User, int, error) {
	db, ctx := GetConnection()
	offset := (page - 1) * limit
	var totalRows int
	qTotalRows := `SELECT COUNT(*) FROM users`

	if err := db.QueryRow(ctx, qTotalRows).Scan(&totalRows); err != nil {
		return []User{}, 0, err
	}

	q := `SELECT id, name, email, is_active, is_admin FROM users LIMIT $1 OFFSET $2`

	rows, err := db.Query(ctx, q, limit, offset)
	if err != nil {
		return []User{}, 0, err
	}
	defer rows.Close()

	users, err := pgx.CollectRows(rows, pgx.RowToStructByName[User])
	if err != nil {
		return []User{}, 0, err
	}

	return users, totalRows, nil
}

func (u User) Update() error {
	db, ctx := GetConnection()
	q := `UPDATE users set name=$1, email=$2, is_active=$3, is_admin=$4 where id=$5`

	if _, err := db.Exec(ctx, q, u.Name, u.Email, u.IsActive, u.IsAdmin, u.ID); err != nil {
		return err
	}
	return nil
}

func delete(id int64) error {
	db, ctx := GetConnection()
	q := `DELETE FROM users where ID=$1`
	if _, err := db.Exec(ctx, q, id); err != nil {
		return err
	}
	return nil
}

func GetUsersByIdHandler(w http.ResponseWriter, r *http.Request) {
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

	user, err := getUserById(id)
	if err != nil {
		http.Error(w, "No se encontro usuario", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "applitation/json")
	json.NewEncoder(w).Encode(user)
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

	users, totalRows, err := u.getAll(page, limit)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	totalPages := (totalRows + limit - 1) / limit
	response := PaginatedUserResponse{
		Data:       users,
		TotalRows:  totalRows,
		TotalPages: totalPages,
		Page:       page,
		Limit:      limit,
	}
	w.Header().Set("Content-Type", "applitation/json")
	json.NewEncoder(w).Encode(response)
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

func UpdateUsersHandler(w http.ResponseWriter, r *http.Request) {
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

func DeleteUsersHandler(w http.ResponseWriter, r *http.Request) {
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

func userTaskHandlers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getUserWithTasksHandler(w, r)
	default:
		http.Error(w, "Metodo no permitido", http.StatusBadRequest)
		return
	}
}

func tasksHandlers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		GetTasksHandler(w, r)
	case http.MethodPost:
		CreateTaskHandler(w, r)
	default:
		http.Error(w, "Metodo no permitido", http.StatusBadRequest)
		return
	}
}

func usersHandlers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		idStr := r.URL.Query().Get("id")
		if idStr != "" {
			GetUsersByIdHandler(w, r)
		} else {
			GetUsersHandler(w, r)
		}
	case http.MethodPost:
		CreateUserHandler(w, r)
	case http.MethodPut:
		UpdateUsersHandler(w, r)
	case http.MethodDelete:
		DeleteUsersHandler(w, r)
	default:
		http.Error(w, "Metodo no permitido", http.StatusBadRequest)
		return
	}
}
