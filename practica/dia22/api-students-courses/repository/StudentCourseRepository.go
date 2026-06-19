package repository

import (
	"api-students-courses/domain"
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type StudentCourseRepository interface {
	AddStudentToCourse(student_id, course_id int64) error
	GetStudentCourses(student_id int64) ([]domain.Course, error)
	GetStudentInCourses(course_id int64) ([]domain.Student, error)
}

type sqlStudentCourseRepository struct {
	db  *pgxpool.Pool
	ctx context.Context
}

func NewSQLStudentCourseRepository(db *pgxpool.Pool, ctx context.Context) StudentCourseRepository {
	return &sqlStudentCourseRepository{db: db, ctx: ctx}
}

func (r *sqlStudentCourseRepository) AddStudentToCourse(student_id, course_id int64) error {
	q := `INSERT INTO student_courses(student_id, course_id) VALUES($1, $2)`
	if _, err := r.db.Exec(r.ctx, q, student_id, course_id); err != nil {
		return err
	}
	return nil
}

func (r *sqlStudentCourseRepository) GetStudentCourses(student_id int64) ([]domain.Course, error) {
	q := `SELECT c.course_id, c.name
		FROM courses c
		INNER JOIN student_courses sc ON c.course_id = sc.course_id
		WHERE sc.student_id=$1;`

	rows, err := r.db.Query(r.ctx, q, student_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	stdcourses, err := pgx.CollectRows(rows, pgx.RowToStructByName[domain.Course])
	if err != nil {
		return nil, err
	}
	return stdcourses, nil

}

func (r *sqlStudentCourseRepository) GetStudentInCourses(course_id int64) ([]domain.Student, error) {
	q := `SELECT s.student_id, s.name 
		FROM students s
		INNER JOIN student_courses sc ON s.student_id = sc.student_id
		WHERE sc.course_id=$1`

	rows, err := r.db.Query(r.ctx, q, course_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	stdcourses, err := pgx.CollectRows(rows, pgx.RowToStructByName[domain.Student])
	if err != nil {
		return nil, err
	}
	return stdcourses, nil

}
