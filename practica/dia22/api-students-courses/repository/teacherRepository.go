package repository

import (
	"api-students-courses/domain"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type TeacherRepository interface {
	Create(s *domain.Teacher) error
	GetById(id int64) (*domain.Teacher, error)
	GetAll(page, limit int) ([]domain.Teacher, int, error)
	Update(s *domain.Teacher) error
	Delete(id int64) error
}

type sqlTeacherRepository struct {
	db  *pgxpool.Pool
	ctx context.Context
}

func NewSQLTeacherRepository(db *pgxpool.Pool, ctx context.Context) TeacherRepository {
	return &sqlTeacherRepository{db: db, ctx: ctx}
}

func (r *sqlTeacherRepository) Create(s *domain.Teacher) error {
	q := `INSERT INTO teachers(name,shift) VALUES($1,$2);`
	if _, err := r.db.Exec(r.ctx, q, s.Name, s.Shift); err != nil {
		return err
	}
	return nil
}

func (r *sqlTeacherRepository) GetById(id int64) (*domain.Teacher, error) {
	q := `SELECT teacher_id, name, shift FROM teachers WHERE teacher_id=$1`
	row, err := r.db.Query(r.ctx, q, id)
	if err != nil {
		return nil, err
	}
	defer row.Close()

	data, err := pgx.CollectOneRow(row, pgx.RowToStructByName[domain.Teacher])
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return &data, nil
}

func (r *sqlTeacherRepository) GetAll(page, limit int) ([]domain.Teacher, int, error) {
	offset := (page - 1) * limit
	var totalRows int
	qTotalRows := `SELECT COUNT(*) FROM teachers`
	if err := r.db.QueryRow(r.ctx, qTotalRows).Scan(&totalRows); err != nil {
		return nil, 0, err
	}

	q := `SELECT teacher_id, name, shift FROM teachers LIMIT $1 OFFSET $2`

	rows, err := r.db.Query(r.ctx, q, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	data, err := pgx.CollectRows(rows, pgx.RowToStructByName[domain.Teacher])
	if err != nil {
		fmt.Println(err.Error())
		return nil, 0, err
	}

	return data, totalRows, nil
}

func (r *sqlTeacherRepository) Update(s *domain.Teacher) error {
	q := `UPDATE teachers set name=$1, shift=$2 WHERE student_id=$3`
	if _, err := r.db.Exec(r.ctx, q, s.Name, s.Shift, s.TeacherId); err != nil {
		return err
	}
	return nil
}

func (r *sqlTeacherRepository) Delete(id int64) error {
	q := `DELETE FROM teachers WHERE teacher_id=$1`
	if _, err := r.db.Exec(r.ctx, q, id); err != nil {
		return err
	}
	return nil
}
