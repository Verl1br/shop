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

type Basket interface {
	AddToBasket(userId, itemId int) (int, error)
	DeleteBasketItem(itemId int) error
	GetBasketItems(userId int) ([]shop.Item, error)
}

type Brand interface {
	AddBrand(item shop.Brand) (int, error)
	DeleteBrand(id int) error
}

type ItemInfo interface {
	Create(item shop.ItemInfo) (int, error)
	GetInfo(id int) (shop.ItemInfo, error)
	DeleteInfo(id int) error
	UpdateInfo(input shop.ItemInfoUpdateInput, id int) error
}

type Service struct {
	Authorization
	Item
	Basket
	Brand
	ItemInfo
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo),
		Item:          NewItemService(repo),
		Basket:        NewBasketService(repo),
		Brand:         NewBrandService(repo),
		ItemInfo:      NewItemInfoService(repo),
	}
}
