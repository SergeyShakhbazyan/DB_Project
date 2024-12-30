package service

import (
	"ProjectDB/internal"
	"ProjectDB/internal/repository"
)

type EquipmentService interface {
	CreateEquipment(equipment internal.Equipment) error
	DeleteEquipment(id int) error
	GetFilteredEquipment(name, manufacturer string) ([]internal.Equipment, error)
	GetEquipmentWithMaterials() ([]map[string]interface{}, error)
	UpdateEquipmentLifeTime(startDate string) error
}

type equipmentService struct {
	repo repository.EquipmentRepository
}

func NewEquipmentService(repo repository.EquipmentRepository) EquipmentService {
	return &equipmentService{repo: repo}
}

func (s *equipmentService) CreateEquipment(equipment internal.Equipment) error {
	return s.repo.CreateEquipment(equipment)
}

func (s *equipmentService) DeleteEquipment(id int) error {
	return s.repo.DeleteEquipment(id)
}

func (s *equipmentService) GetFilteredEquipment(name, manufacturer string) ([]internal.Equipment, error) {
	return s.repo.GetFilteredEquipment(name, manufacturer)
}

func (s *equipmentService) GetEquipmentWithMaterials() ([]map[string]interface{}, error) {
	return s.repo.GetEquipmentWithMaterials()
}

func (s *equipmentService) UpdateEquipmentLifeTime(startDate string) error {
	return s.repo.UpdateEquipmentLifeTime(startDate)
}
