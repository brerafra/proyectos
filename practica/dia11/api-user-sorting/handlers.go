package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

func (p Producto) create() error {
	db := GetConnection()
	q := `INSERT INTO productos (nombre, precio, status) VALUES (?,?,?)`
	stmt, err := db.Prepare(q)
	if err != nil {
		return err
	}
	defer stmt.Close()

	r, err := stmt.Exec(p.Nombre, p.Precio, p.Status)
	if err != nil {
		return err
	}

	if i, err := r.RowsAffected(); err != nil || i != 1 {
		return errors.New("Error: se esperaba una fila afectada")
	}

	return nil
}

func (p *Producto) getProductos(page int, limit int, ordenar string, dir string) ([]Producto, int, error) {
	db := GetConnection()
	offset := (page - 1) * limit
	var totalRows int

	qTotalRows := `SELECT COUNT(*) FROM productos`
	q := `SELECT id, nombre, precio, status FROM productos`
	if ordenar != "" {
		if dir == "asc" || dir == "desc" {
			q = q + " ORDER BY " + ordenar + " " + dir
		}
	}
	q = q + ` LIMIT ? OFFSET ?`
	fmt.Println(q)

	if err := db.QueryRow(qTotalRows).Scan(&totalRows); err != nil {
		return []Producto{}, 0, err
	}

	rows, err := db.Query(q, limit, offset)
	if err != nil {
		return []Producto{}, 0, err
	}
	defer rows.Close()

	productos := []Producto{}
	for rows.Next() {
		rows.Scan(
			&p.ID,
			&p.Nombre,
			&p.Precio,
			&p.Status,
		)
		productos = append(productos, *p)
	}
	return productos, totalRows, nil
}

func (p Producto) Update() error {
	db := GetConnection()
	q := `UPDATE productos set nombre=?, precio=?, status=? where id=?`
	stmt, err := db.Prepare(q)
	if err != nil {
		return err
	}
	defer stmt.Close()

	r, err := stmt.Exec(p.Nombre, p.Precio, p.Status, p.ID)
	if err != nil {
		return err
	}

	if i, err := r.RowsAffected(); err != nil || i != 1 {
		return errors.New("Error se esperaba una fila afectada")
	}
	return nil
}

func delete(id int64) error {
	db := GetConnection()

	q := `DELETE FROM productos WHERE id=?`
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

func GetProductsHandler(w http.ResponseWriter, r *http.Request) {
	p := new(Producto)
	query := r.URL.Query()
	pageStr := query.Get("page")
	limitStr := query.Get("limit")
	ordenar := query.Get("ordenar")
	dir := strings.ToLower(query.Get("dir"))

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 {
		limit = 10
	}

	if dir != "asc" {
		if dir != "desc" {
			dir = ""
		}
	}

	products, totalRows, err := p.getProductos(page, limit, ordenar, dir)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	totalPages := (totalRows + limit - 1) / limit
	response := PaginatedResponse{
		Data:       products,
		TotalRows:  totalRows,
		TotalPages: totalPages,
		Page:       page,
		Limit:      limit,
	}
	w.Header().Set("Content-Type", "applitation/json")
	json.NewEncoder(w).Encode(response)

}

func CreateProductHandler(w http.ResponseWriter, r *http.Request) {
	var product Producto
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := product.create(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func UpdateProductHandler(w http.ResponseWriter, r *http.Request) {
	var product Producto

	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := product.Update(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func DeleteProductHandler(w http.ResponseWriter, r *http.Request) {
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

func ProductsHandlers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		GetProductsHandler(w, r)
	case http.MethodPost:
		CreateProductHandler(w, r)
	case http.MethodPut:
		UpdateProductHandler(w, r)
	case http.MethodDelete:
		DeleteProductHandler(w, r)
	default:
		http.Error(w, "Metodo no permitido", http.StatusBadRequest)
		return
	}
}
