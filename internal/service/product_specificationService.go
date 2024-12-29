package service

import (
	"ProjectDB/internal"
	"ProjectDB/internal/repository"
)

type ProductSpecificationService interface {
	CreateProductSpecification(spec internal.ProductSpecification) error
	GetProductSpecificationByID(id int) (*internal.ProductSpecification, error)
	GetAllProductSpecifications() ([]internal.ProductSpecification, error)
	UpdateProductSpecification(id int, spec internal.ProductSpecification) error
	DeleteProductSpecification(id int) error
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
}
