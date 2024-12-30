package service

import (
	"ProjectDB/internal"
	"ProjectDB/internal/repository"
)

type MaterialService interface {
	CreateMaterial(material internal.Material) error
	GetMaterialByID(id int) (*internal.Material, error)
	GetAllMaterials() ([]internal.Material, error)
	UpdateMaterial(id int, material internal.Material) error
	DeleteMaterial(id int) error
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
}
