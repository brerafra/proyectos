package repository

import (
	"api-students-courses/domain"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type StudentRepository interface {
	Create(s *domain.Student) error
	GetById(id int64) (*domain.Student, error)
	GetAll(page, limit int) ([]domain.Student, int, error)
	Update(s *domain.Student) error
	Delete(id int64) error
}

type sqlRepository struct {
	db  *pgxpool.Pool
	ctx context.Context
}

func NewSQLRepository(db *pgxpool.Pool, ctx context.Context) StudentRepository {
	return &sqlRepository{db: db, ctx: ctx}
}

func (r *sqlRepository) Create(s *domain.Student) error {
	q := `INSERT INTO students(name) VALUES($1);`
	if _, err := r.db.Exec(r.ctx, q, s.Name); err != nil {
		return err
	}
	return nil
}

func (r *sqlRepository) GetById(id int64) (*domain.Student, error) {
	q := `SELECT student_id, name FROM students WHERE student_id=$1`
	row, err := r.db.Query(r.ctx, q, id)
	if err != nil {
		return nil, err
	}
	defer row.Close()

	student, err := pgx.CollectOneRow(row, pgx.RowToStructByName[domain.Student])
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return &student, nil
}

func (r *sqlRepository) GetAll(page, limit int) ([]domain.Student, int, error) {
	offset := (page - 1) * limit
	var totalRows int
	qTotalRows := `SELECT COUNT(*) FROM students`
	if err := r.db.QueryRow(r.ctx, qTotalRows).Scan(&totalRows); err != nil {
		return nil, 0, err
	}

	q := `SELECT student_id, name FROM students LIMIT $1 OFFSET $2`

	rows, err := r.db.Query(r.ctx, q, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	students, err := pgx.CollectRows(rows, pgx.RowToStructByName[domain.Student])
	if err != nil {
		fmt.Println(err.Error())
		return nil, 0, err
	}

	return students, totalRows, nil
}

func (r *sqlRepository) Update(s *domain.Student) error {
	q := `UPDATE students set name=$1 WHERE student_id=$2`
	if _, err := r.db.Exec(r.ctx, q, s.Name, s.StudentId); err != nil {
		return err
	}
	return nil
}

func (r *sqlRepository) Delete(id int64) error {
	q := `DELETE FROM students WHERE student_id=$1`
	if _, err := r.db.Exec(r.ctx, q, id); err != nil {
		return err
	}
	return nil
}
