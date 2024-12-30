package service

import (
	"ProjectDB/internal"
	"ProjectDB/internal/repository\"
)

type EquipmentService interface {
	CreateEquipment(equipment internal.Equipment) error
	GetEquipmentByID(id int) (*internal.Equipment, error)
	GetAllEquipment() ([]internal.Equipment, error)
	UpdateEquipment(id int, equipment internal.Equipment) error
	DeleteEquipment(id int) error
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

func (s *equipmentService) GetEquipmentByID(id int) (*internal.Equipment, error) {
	return s.repo.GetEquipmentByID(id)
}

func (s *equipmentService) GetAllEquipment() ([]internal.Equipment, error) {
	return s.repo.GetAllEquipment()
}

func (s *equipmentService) UpdateEquipment(id int, equipment internal.Equipment) error {
	return s.repo.UpdateEquipment(id, equipment)
}

func (s *equipmentService) DeleteEquipment(id int) error {
	return s.repo.DeleteEquipment(id)
}
