package service

import (
	"ProjectDB/internal"
	"ProjectDB/internal/repository"
)

type ProductSpecificationService interface {
	CreateProductSpecification(spec internal.ProductSpecification) error
<<<<<<< HEAD
	GetSpecificationsByEquipment(equipmentID int) ([]internal.ProductSpecification, error)
	UpdateProductionDuration(id int, newDuration int) error
	GetProductionCountByEquipment() (map[int]int, error)
=======
	GetProductSpecificationByID(id int) (*internal.ProductSpecification, error)
	GetAllProductSpecifications() ([]internal.ProductSpecification, error)
	UpdateProductSpecification(id int, spec internal.ProductSpecification) error
	DeleteProductSpecification(id int) error
>>>>>>> main
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

<<<<<<< HEAD
func (s *productSpecificationService) GetSpecificationsByEquipment(equipmentID int) ([]internal.ProductSpecification, error) {
	return s.repo.GetSpecificationsByEquipment(equipmentID)
}

func (s *productSpecificationService) UpdateProductionDuration(id int, newDuration int) error {
	return s.repo.UpdateProductionDuration(id, newDuration)
}

func (s *productSpecificationService) GetProductionCountByEquipment() (map[int]int, error) {
	return s.repo.GetProductionCountByEquipment()
=======
func (s *productSpecificationService) GetProductSpecificationByID(id int) (*internal.ProductSpecification, error) {
	return s.repo.GetProductSpecificationByID(id)
}

func (s *productSpecificationService) GetAllProductSpecifications() ([]internal.ProductSpecification, error) {
	return s.repo.GetAllProductSpecifications()
}

func (s *productSpecificationService) UpdateProductSpecification(id int, spec internal.ProductSpecification) error {
	return s.repo.UpdateProductSpecification(id, spec)
}

func (s *productSpecificationService) DeleteProductSpecification(id int) error {
	return s.repo.DeleteProductSpecification(id)
>>>>>>> main
}
