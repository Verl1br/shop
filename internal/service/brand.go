package service

import (
	"github.com/dhevve/shop"
	"github.com/dhevve/shop/internal/repository"
)

type BrandService struct {
	repo repository.Brand
}

func NewBrandService(repo repository.Brand) *BrandService {
	return &BrandService{repo: repo}
}

func (s *BrandService) AddBrand(item shop.Brand) (int, error) {
	return s.repo.AddBrand(item)
}

func (s *BrandService) DeleteBrand(id int) error {
	return s.repo.DeleteBrand(id)
}
