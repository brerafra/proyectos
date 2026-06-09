package repository

import (
	"database/sql"
	"main/domain"
)

type UserRepository interface {
	Create(user *domain.User) error
	GetById(id int) (*domain.User, error)
}

type sqlRepository struct {
	db *sql.DB
}

func NewSQLRepository(db *sql.DB) UserRepository {
	return &sqlRepository{db: db}
}

func (r *sqlRepository) Create(user *domain.User) error {
	query := `INSERT INTO users (name, email) VALUES (?,?)`
	result, err := r.db.Exec(query, user.Name, user.Email)
	if err != nil {
		return err
	}
	id, _ := result.LastInsertId()
	user.ID = int(id)
	return nil
}

func (r *sqlRepository) GetById(id int) (*domain.User, error) {
	query := `SELECT id, name, email FROM users WHERE id = ?`
	row := r.db.QueryRow(query, id)

	var user domain.User
	if err := row.Scan(&user.ID, &user.Name, &user.Email); err != nil {
		return nil, err
	}
	return &user, nil
}
