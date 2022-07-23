package repository

import (
	"fmt"
	"strings"

	"github.com/dhevve/shop"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type ItemPostgres struct {
	db *sqlx.DB
}

func NewItemPostgres(db *sqlx.DB) *ItemPostgres {
	return &ItemPostgres{db: db}
}

func (r *ItemPostgres) CreateItem(input shop.Item) (int, error) {
	var id int

	query := fmt.Sprintf("INSERT INTO %s (name, price, brand_id) values ($1, $2, $3) RETURNING id", "items")
	row := r.db.QueryRow(query, input.Name, input.Price, input.BrandId)

	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *ItemPostgres) GetAllItems() ([]shop.Item, error) {
	var items []shop.Item

	query := fmt.Sprintf("SELECT name, price FROM %s", "items")
	err := r.db.Select(&items, query)
	return items, err
}

func (r *ItemPostgres) GetById(id int) (shop.Item, error) {
	var item shop.Item

	query := fmt.Sprintf("SELECT name, price FROM %s WHERE id = $1", "items")

	err := r.db.Get(&item, query, id)
	return item, err
}

func (r *ItemPostgres) DeleteItem(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", "items")
	_, err := r.db.Exec(query, id)
	return err
}

func (r *ItemPostgres) UpdateItem(input shop.ItemUpdateInput, id int) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Name != nil {
		setValues = append(setValues, fmt.Sprintf("name=$%d", argId))
		args = append(args, *input.Name)
		argId++
	}

	if input.Price != nil {
		setValues = append(setValues, fmt.Sprintf("price=$%d", argId))
		args = append(args, *input.Price)
		argId++
	}

	if input.BrandId != nil {
		setValues = append(setValues, fmt.Sprintf("brand_id=$%d", argId))
		args = append(args, *input.BrandId)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE items SET %s WHERE id = $%d",
		setQuery, argId)
	args = append(args, id)

	logrus.Debugf("updateQuery: %s", query)
	logrus.Debugf("args: %s", args)

	_, err := r.db.Exec(query, args...)
	return err
}
