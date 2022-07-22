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

type Basket interface {
	AddToBasket(userId, itemId int) (int, error)
	DeleteBasketItem(itemId int) error
	GetBasketItems(userId int) ([]shop.Item, error)
}

type Brand interface {
	AddBrand(item shop.Brand) (int, error)
	DeleteBrand(id int) error
}

type Repository struct {
	Authorization
	Item
	Basket
	Brand
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Item:          NewItemPostgres(db),
		Basket:        NewBasketPostgres(db),
		Brand:         NewBrandPostgres(db),
	}
}
