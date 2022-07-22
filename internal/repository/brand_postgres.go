package repository

import (
	"fmt"

	"github.com/dhevve/shop"
	"github.com/jmoiron/sqlx"
)

type BrandPostgres struct {
	db *sqlx.DB
}

func NewBrandPostgres(db *sqlx.DB) *BrandPostgres {
	return &BrandPostgres{db: db}
}

func (r *BrandPostgres) AddBrand(item shop.Brand) (int, error) {
	var id int

	query := fmt.Sprintf("INSERT INTO %s (name) values ($1) RETURNING id", "brand")
	row := r.db.QueryRow(query, item.Name)

	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *BrandPostgres) DeleteBrand(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", "brand")
	_, err := r.db.Exec(query, id)
	return err
}
