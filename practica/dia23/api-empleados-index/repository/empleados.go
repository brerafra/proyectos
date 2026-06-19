package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"main.go/domain"
)

type EmployeeRepository interface {
	Create(s *domain.Empleado) error
	GetById(id int64) (*domain.Empleado, error)
	GetAll(page, limit int) ([]domain.Empleado, int, error)
	Update(s *domain.Empleado) error
	Delete(id int64) error
}

type sqlEmployeeRepository struct {
	db  *pgxpool.Pool
	ctx context.Context
}

func NewSQLEmployeeRepository(db *pgxpool.Pool, ctx context.Context) EmployeeRepository {
	return &sqlEmployeeRepository{db: db, ctx: ctx}
}

func (r *sqlEmployeeRepository) Create(s *domain.Empleado) error {
	q := `INSERT INTO empleados(nombre, departamento, salario) VALUES($1,$2,$3);`
	if _, err := r.db.Exec(r.ctx, q, s.Nombre, s.Departamento, s.Salario); err != nil {
		return err
	}
	return nil
}

func (r *sqlEmployeeRepository) GetById(id int64) (*domain.Empleado, error) {
	q := `SELECT id, nombre, departamento, salario 
			FROM empleados
			WHERE id=$1`
	row, err := r.db.Query(r.ctx, q, id)
	if err != nil {
		return nil, err
	}
	defer row.Close()

	jdata, err := pgx.CollectOneRow(row, pgx.RowToStructByName[domain.Empleado])
	if err != nil {
		return nil, err
	}
	return &jdata, nil
}

func (r *sqlEmployeeRepository) GetAll(page, limit int) ([]domain.Empleado, int, error) {
	offset := (page - 1) * limit
	var tRows int
	q := `SELECT COUNT(*) FROM empleados`
	if err := r.db.QueryRow(r.ctx, q).Scan(&tRows); err != nil {
		return nil, 0, err
	}

	q = `SELECT id, nombre, departamento, salario FROM empleados LIMIT $1 OFFSET $2`

	rows, err := r.db.Query(r.ctx, q, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	jdata, err := pgx.CollectRows(rows, pgx.RowToStructByName[domain.Empleado])
	if err != nil {
		return nil, 0, err
	}
	return jdata, tRows, nil

}

func (r *sqlEmployeeRepository) Update(s *domain.Empleado) error {
	q := `UPDATE empleados set nombre=$1, departamento=$2, salario=$3 WHERE id=$4`
	if _, err := r.db.Exec(r.ctx, q, s.Nombre, s.Departamento, s.Salario, s.Id); err != nil {
		return err
	}
	return nil
}

func (r *sqlEmployeeRepository) Delete(id int64) error {
	q := `DELETE FROM empleados WHERE id=$1`
	if _, err := r.db.Exec(r.ctx, q, id); err != nil {
		return err
	}
	return nil
}
