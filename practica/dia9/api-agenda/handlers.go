package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

func (a Agenda) create() error {
	db := GetConnection()
	q := `INSERT INTO agenda (nombre, numero, direccion, edad) VALUES(?,?,?,?)`

	stmt, err := db.Prepare(q)
	if err != nil {
		return err
	}
	defer stmt.Close()

	r, err := stmt.Exec(a.Nombre, a.Numero, a.Direccion, a.Edad)
	if err != nil {
		return err
	}

	if i, err := r.RowsAffected(); err != nil || i != 1 {
		return errors.New("Error: se esperaba una fila afectada")
	}

	return nil
}

func (a *Agenda) GetAll() ([]Agenda, error) {
	db := GetConnection()
	q := `SELECT id, nombre, numero, direccion, edad FROM agenda`

	rows, err := db.Query(q)
	if err != nil {
		return []Agenda{}, err
	}
	defer rows.Close()

	agendas := []Agenda{}

	for rows.Next() {
		rows.Scan(
			&a.ID,
			&a.Nombre,
			&a.Numero,
			&a.Direccion,
			&a.Edad,
		)
		agendas = append(agendas, *a)
	}
	return agendas, nil
}

func (a Agenda) Update() error {
	db := GetConnection()
	q := `UPDATE agenda set nombre=?, numero=?, direccion=?, edad=? where id=?`
	stmt, err := db.Prepare(q)
	if err != nil {
		return err
	}
	defer stmt.Close()

	r, err := stmt.Exec(a.Nombre, a.Numero, a.Direccion, a.Edad, a.ID)
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
	q := `DELETE FROM agenda WHERE id=?`
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

func GetAgendaHandler(w http.ResponseWriter, r *http.Request) {
	a := new(Agenda)

	agendas, err := a.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	j, err := json.Marshal(agendas)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}

func CreateAgendaHandler(w http.ResponseWriter, r *http.Request) {
	var agenda Agenda

	if err := json.NewDecoder(r.Body).Decode(&agenda); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := agenda.create(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func UpdateAgendaHandler(w http.ResponseWriter, r *http.Request) {
	var agenda Agenda

	if err := json.NewDecoder(r.Body).Decode(&agenda); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := agenda.Update(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func DeleteAgendaHandler(w http.ResponseWriter, r *http.Request) {
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

func AgendasHandlers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		GetAgendaHandler(w, r)
	case http.MethodPost:
		CreateAgendaHandler(w, r)
	case http.MethodPut:
		UpdateAgendaHandler(w, r)
	case http.MethodDelete:
		DeleteAgendaHandler(w, r)
	default:
		http.Error(w, "Metodo no permitido", http.StatusBadRequest)
		return
	}
}
