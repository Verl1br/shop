package repository

import (
	"fmt"

	"github.com/dhevve/shop"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user shop.User) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int

	createQueryUser := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) values ($1, $2, $3) RETURNING id", "users")

	row := tx.QueryRow(createQueryUser, user.Name, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	createQueryBasket := fmt.Sprintf("INSERT INTO %s (user_id) values ($1) RETURNING id", "baskets")
	_, err = tx.Exec(createQueryBasket, id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}

func (r *AuthPostgres) GetUser(username, password string) (shop.User, error) {
	var user shop.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE username=$1 AND password_hash=$2", "users")
	err := r.db.Get(&user, query, username, password)
	return user, err
}
