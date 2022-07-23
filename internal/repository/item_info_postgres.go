package repository

import (
	"fmt"
	"strings"

	"github.com/dhevve/shop"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type ItemInfoPostgres struct {
	db *sqlx.DB
}

func NewItemInfoPostgres(db *sqlx.DB) *ItemInfoPostgres {
	return &ItemInfoPostgres{db: db}
}

func (r *ItemInfoPostgres) Create(item shop.ItemInfo) (int, error) {
	var id int

	query := fmt.Sprintf("INSERT INTO %s (title, description, item_id) VALUES ($1, $2, $3) RETURNING id", "item_info")
	row := r.db.QueryRow(query, item.Title, item.Description, item.ItemId)

	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *ItemInfoPostgres) GetInfo(id int) (shop.ItemInfo, error) {
	var item shop.ItemInfo

	query := fmt.Sprintf("SELECT id, title, description FROM %s WHERE id = $1", "item_info")
	err := r.db.Get(&item, query, id)
	return item, err
}

func (r *ItemInfoPostgres) DeleteInfo(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", "item_info")
	_, err := r.db.Exec(query, id)
	return err
}

func (r *ItemInfoPostgres) UpdateInfo(input shop.ItemInfoUpdateInput, id int) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, *input.Title)
		argId++
	}

	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, *input.Description)
		argId++
	}

	if input.ItemId != nil {
		setValues = append(setValues, fmt.Sprintf("item_id=$%d", argId))
		args = append(args, *input.ItemId)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE item_info SET %s WHERE id = $%d",
		setQuery, argId)
	args = append(args, id)

	logrus.Debugf("updateQuery: %s", query)
	logrus.Debugf("args: %s", args)

	_, err := r.db.Exec(query, args...)
	return err
}
