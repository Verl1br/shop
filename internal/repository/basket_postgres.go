package repository

import (
	"fmt"

	"github.com/dhevve/shop"
	"github.com/jmoiron/sqlx"
)

type BasketPostgres struct {
	db *sqlx.DB
}

func NewBasketPostgres(db *sqlx.DB) *BasketPostgres {
	return &BasketPostgres{db: db}
}

func (r *BasketPostgres) AddToBasket(userId, itemId int) (int, error) {
	var id int

	query := fmt.Sprintf("INSERT INTO %s (basket_id, item_id) VALUES ($1, $2) RETURNING id", "basket_item")
	row := r.db.QueryRow(query, userId, itemId)

	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *BasketPostgres) DeleteBasketItem(itemId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE item_id = $1", "basket_item")
	_, err := r.db.Exec(query, itemId)
	return err
}

func (r *BasketPostgres) GetBasketItems(userId int) ([]shop.Item, error) {
	var items []shop.Item

	query := fmt.Sprintf("SELECT items.id, items.name, items.price FROM %s INNER JOIN %s ON items.id = basket_item.item_id WHERE basket_item.basket_id = $1", "items", "basket_item")
	err := r.db.Select(&items, query, userId)
	return items, err
}
