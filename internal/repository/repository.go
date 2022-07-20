package repository

import (
	"github.com/dhevve/shop"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user shop.User) (int, error)
	GetUser(username, password string) (shop.User, error)
}

type Item interface {
	CreateItem(input shop.Item) (int, error)
	GetAllItems() ([]shop.Item, error)
	GetById(id int) (shop.Item, error)
	DeleteItem(id int) error
	UpdateItem(input shop.ItemUpdateInput, id int) error
}

type Repository struct {
	Authorization
	Item
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Item:          NewItemPostgres(db),
	}
}
