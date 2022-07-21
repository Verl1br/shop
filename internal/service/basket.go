package service

import (
	"github.com/dhevve/shop"
	"github.com/dhevve/shop/internal/repository"
)

type BasketService struct {
	repo repository.Basket
}

func NewBasketService(repo repository.Basket) *BasketService {
	return &BasketService{repo: repo}
}

func (s *BasketService) AddToBasket(userId, itemId int) (int, error) {
	return s.repo.AddToBasket(userId, itemId)
}

func (s *BasketService) DeleteBasketItem(itemId int) error {
	return s.repo.DeleteBasketItem(itemId)
}

func (s *BasketService) GetBasketItems(userId int) ([]shop.Item, error) {
	return s.repo.GetBasketItems(userId)
}
