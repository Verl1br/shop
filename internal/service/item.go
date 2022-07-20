package service

import (
	"github.com/dhevve/shop"
	"github.com/dhevve/shop/internal/repository"
)

type ItemService struct {
	repo repository.Item
}

func NewItemService(repo repository.Item) *ItemService {
	return &ItemService{repo: repo}
}

func (s *ItemService) CreateItem(input shop.Item) (int, error) {
	return s.repo.CreateItem(input)
}

func (s *ItemService) GetAllItems() ([]shop.Item, error) {
	return s.repo.GetAllItems()
}

func (s *ItemService) GetById(id int) (shop.Item, error) {
	return s.repo.GetById(id)
}

func (s *ItemService) DeleteItem(id int) error {
	return s.repo.DeleteItem(id)
}

func (s *ItemService) UpdateItem(input shop.ItemUpdateInput, id int) error {
	return s.repo.UpdateItem(input, id)
}
