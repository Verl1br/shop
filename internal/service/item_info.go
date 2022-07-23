package service

import (
	"github.com/dhevve/shop"
	"github.com/dhevve/shop/internal/repository"
)

type ItemInfoService struct {
	repo repository.ItemInfo
}

func NewItemInfoService(repo repository.ItemInfo) *ItemInfoService {
	return &ItemInfoService{repo: repo}
}

func (s *ItemInfoService) Create(item shop.ItemInfo) (int, error) {
	return s.repo.Create(item)
}

func (s *ItemInfoService) GetInfo(id int) (shop.ItemInfo, error) {
	return s.repo.GetInfo(id)
}

func (s *ItemInfoService) DeleteInfo(id int) error {
	return s.repo.DeleteInfo(id)
}

func (s *ItemInfoService) UpdateInfo(input shop.ItemInfoUpdateInput, id int) error {
	return s.repo.UpdateInfo(input, id)
}
