package service

import (
	"ProjectDB/internal"
)

type EquipmentService interface {
	CreateEquipment(equipment internal.Equipment) error
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

func (s *equipmentService) DeleteEquipment(id int) error {
	return s.repo.DeleteEquipment(id)
}
