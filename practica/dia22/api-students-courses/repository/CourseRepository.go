package repository

import (
	"api-students-courses/domain"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type CourseRepository interface {
	Create(c *domain.Course) error
	GetById(id int64) (*domain.Course, error)
	GetAll(page, limit int) ([]domain.Course, int, error)
	Update(s *domain.Course) error
	Delete(id int64) error
}

type sqlCourseRepository struct {
	db  *pgxpool.Pool
	ctx context.Context
}

func NewSQLCourseRepository(db *pgxpool.Pool, ctx context.Context) CourseRepository {
	return &sqlCourseRepository{db: db, ctx: ctx}
}

func (r *sqlCourseRepository) Create(s *domain.Course) error {
	q := `INSERT INTO courses(name) VALUES($1);`
	if _, err := r.db.Exec(r.ctx, q, s.Name); err != nil {
		return err
	}
	return nil
}

func (r *sqlCourseRepository) GetById(id int64) (*domain.Course, error) {
	q := `SELECT course_id, name FROM courses WHERE course_id=$1`
	row, err := r.db.Query(r.ctx, q, id)
	if err != nil {
		return nil, err
	}
	defer row.Close()

	course, err := pgx.CollectOneRow(row, pgx.RowToStructByName[domain.Course])
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return &course, nil
}

func (r *sqlCourseRepository) GetAll(page, limit int) ([]domain.Course, int, error) {
	offset := (page - 1) * limit
	var totalRows int
	qTotalRows := `SELECT COUNT(*) FROM courses`
	if err := r.db.QueryRow(r.ctx, qTotalRows).Scan(&totalRows); err != nil {
		return nil, 0, err
	}

	q := `SELECT course_id, name FROM courses LIMIT $1 OFFSET $2`

	rows, err := r.db.Query(r.ctx, q, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	courses, err := pgx.CollectRows(rows, pgx.RowToStructByName[domain.Course])
	if err != nil {
		fmt.Println(err.Error())
		return nil, 0, err
	}

	return courses, totalRows, nil
}

func (r *sqlCourseRepository) Update(s *domain.Course) error {
	q := `UPDATE courses set name=$1 WHERE course_id=$2`
	if _, err := r.db.Exec(r.ctx, q, s.Name, s.CourseId); err != nil {
		return err
	}
	return nil
}

func (r *sqlCourseRepository) Delete(id int64) error {
	q := `DELETE FROM courses WHERE course_id=$1`
	if _, err := r.db.Exec(r.ctx, q, id); err != nil {
		return err
	}
	return nil
}
