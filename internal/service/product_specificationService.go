package service

import (
	"ProjectDB/internal"
	"ProjectDB/internal/repository"
)

type ProductSpecificationService interface {
	CreateProductSpecification(spec internal.ProductSpecification) error
	GetSpecificationsByEquipment(equipmentID int) ([]internal.ProductSpecification, error)
	UpdateProductionDuration(id int, newDuration int) error
	GetProductionCountByEquipment() (map[int]int, error)
}

type productSpecificationService struct {
	repo repository.ProductSpecificationRepository
}

func NewProductSpecificationService(repo repository.ProductSpecificationRepository) ProductSpecificationService {
	return &productSpecificationService{repo: repo}
}

func (s *productSpecificationService) CreateProductSpecification(spec internal.ProductSpecification) error {
	return s.repo.CreateProductSpecification(spec)
}

func (s *productSpecificationService) GetSpecificationsByEquipment(equipmentID int) ([]internal.ProductSpecification, error) {
	return s.repo.GetSpecificationsByEquipment(equipmentID)
}

func (s *productSpecificationService) UpdateProductionDuration(id int, newDuration int) error {
	return s.repo.UpdateProductionDuration(id, newDuration)
}

func (s *productSpecificationService) GetProductionCountByEquipment() (map[int]int, error) {
	return s.repo.GetProductionCountByEquipment()
}
