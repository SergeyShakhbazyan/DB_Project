package service

import (
	"ProjectDB/internal"
	"ProjectDB/internal/repository"
<<<<<<< HEAD
	"fmt"
=======
>>>>>>> main
)

type MaterialService interface {
	CreateMaterial(material internal.Material) error
<<<<<<< HEAD
	GetFilteredMaterials(name, materialType string) ([]internal.Material, error)
	UpdateMaterialPrice(percentage float64) error
	GetMaterialCountByType() (map[string]int, error)
=======
	GetMaterialByID(id int) (*internal.Material, error)
	GetAllMaterials() ([]internal.Material, error)
	UpdateMaterial(id int, material internal.Material) error
	DeleteMaterial(id int) error
>>>>>>> main
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

<<<<<<< HEAD
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
=======
func (s *materialService) GetMaterialByID(id int) (*internal.Material, error) {
	return s.repo.GetMaterialByID(id)
}

func (s *materialService) GetAllMaterials() ([]internal.Material, error) {
	return s.repo.GetAllMaterials()
}

func (s *materialService) UpdateMaterial(id int, material internal.Material) error {
	return s.repo.UpdateMaterial(id, material)
}

func (s *materialService) DeleteMaterial(id int) error {
	return s.repo.DeleteMaterial(id)
>>>>>>> main
}
