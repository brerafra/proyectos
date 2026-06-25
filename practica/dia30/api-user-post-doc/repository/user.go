package repository

import (
	"context"
	"fmt"
	"main/domain"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository interface {
	Create(u *domain.User) error
	GetById(id int64) (*domain.User, error)
	GetAll(page, limit int) ([]domain.User, int, error)
	Update(s *domain.User) error
	Delete(id int64) error
}

type sqlUserRepository struct {
	db  *pgxpool.Pool
	ctx context.Context
}

func NewSQLUserRepository(db *pgxpool.Pool, ctx context.Context) UserRepository {
	return &sqlUserRepository{db: db, ctx: ctx}
}

func (r *sqlUserRepository) Create(u *domain.User) error {
	q := `INSERT INTO users(name, card, email, is_active, is_admin) VALUES($1, $2, $3, $4, $5);`
	if _, err := r.db.Exec(r.ctx, q, u.Name, u.Card, u.Email, u.IsActive, u.IsAdmin); err != nil {
		return err
	}
	return nil
}

func (r *sqlUserRepository) GetById(id int64) (*domain.User, error) {
	q := `SELECT user_id, name, card, email, is_active, is_admin FROM users WHERE user_id=$1`
	row, err := r.db.Query(r.ctx, q, id)
	if err != nil {
		return nil, err
	}
	defer row.Close()

	user, err := pgx.CollectOneRow(row, pgx.RowToStructByName[domain.User])
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return &user, nil

}

func (r *sqlUserRepository) GetAll(page, limit int) ([]domain.User, int, error) {
	offset := (page - 1) * limit
	var totalRows int
	qTotalRows := `SELECT COUNT(*) FROM users`
	if err := r.db.QueryRow(r.ctx, qTotalRows).Scan(&totalRows); err != nil {
		return nil, 0, err
	}

	q := `SELECT user_id, name, card, email, is_active, is_admin FROM users LIMIT $1 OFFSET $2`

	rows, err := r.db.Query(r.ctx, q, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	users, err := pgx.CollectRows(rows, pgx.RowToStructByName[domain.User])
	if err != nil {
		fmt.Println(err.Error())
		return nil, 0, err
	}

	return users, totalRows, nil
}

func (r *sqlUserRepository) Update(s *domain.User) error {
	q := `UPDATE users set name=$1, card=$2, email=$3, is_active=$4, is_admin=$5 WHERE user_id=$6`
	if _, err := r.db.Exec(r.ctx, q, s.Name, s.Card, s.Email, s.IsActive, s.IsAdmin, s.UserId); err != nil {
		return err
	}
	return nil
}

func (r *sqlUserRepository) Delete(id int64) error {
	q := `DELETE FROM users WHERE user_id=$1`
	if _, err := r.db.Exec(r.ctx, q, id); err != nil {
		return err
	}
	return nil
}
