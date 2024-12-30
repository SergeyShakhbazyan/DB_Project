package service

import (
	"ProjectDB/internal"
	"ProjectDB/internal/repository"
	"fmt"
)

type MaterialService interface {
	CreateMaterial(material internal.Material) error
	GetFilteredMaterials(name, materialType string) ([]internal.Material, error)
	UpdateMaterialPrice(percentage float64) error
	GetMaterialCountByType() (map[string]int, error)
}

type materialService struct {
	repo repository.MaterialRepository
}

func NewMaterialService(repo repository.MaterialRepository) MaterialService {
	return &materialService{repo: repo}
}

func (s *materialService) CreateMaterial(material internal.Material) error {
	return s.repo.CreateMaterial(material)
}

func (s *materialService) GetFilteredMaterials(name, materialType string) ([]internal.Material, error) {
	return s.repo.GetFilteredMaterials(name, materialType)
}

func (s *materialService) UpdateMaterialPrice(percentage float64) error {
	if percentage < 0 {
		return fmt.Errorf("percentage must be non-negative")
	}
	return s.repo.UpdateMaterialPrice(percentage)
}

func (s *materialService) GetMaterialCountByType() (map[string]int, error) {
	return s.repo.GetMaterialCountByType()
}
