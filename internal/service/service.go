package service

import (
	"github.com/dhevve/shop"
	"github.com/dhevve/shop/internal/repository"
)

type Authorization interface {
	CreateUser(user shop.User) (int, error)
	ParseToken(accessToken string) (int, error)
	GenerateToken(username, password string) (string, error)
}

type Item interface {
	CreateItem(input shop.Item) (int, error)
	GetAllItems() ([]shop.Item, error)
	GetById(id int) (shop.Item, error)
	DeleteItem(id int) error
	UpdateItem(input shop.ItemUpdateInput, id int) error
}

type Service struct {
	Authorization
	Item
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo),
		Item:          NewItemService(repo),
	}
}
