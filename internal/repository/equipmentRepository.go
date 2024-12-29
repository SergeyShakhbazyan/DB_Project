package repository

import (
	"ProjectDB/config"
	"ProjectDB/internal"
)

type EquipmentRepository interface {
	CreateEquipment(equipment internal.Equipment) error
	GetEquipmentByID(id int) (*internal.Equipment, error)
	GetAllEquipment() ([]internal.Equipment, error)
	UpdateEquipment(id int, equipment internal.Equipment) error
	DeleteEquipment(id int) error
}

type equipmentRepository struct {
	session *config.Connection
}

func NewEquipmentRepository(session *config.Connection) EquipmentRepository {
	return &equipmentRepository{session: session}
}

func (r *equipmentRepository) CreateEquipment(equipment internal.Equipment) error {
	query := `INSERT INTO "MyDatabase".public.equipment (name, manufacturer, start_date, lifeTime) VALUES ($1, $2, $3, $4)`
	_, err := r.session.DB.Exec(query, equipment.Name, equipment.Manufacturer, equipment.StartDate, equipment.LifeTime)
	return err
}

func (r *equipmentRepository) GetEquipmentByID(id int) (*internal.Equipment, error) {
	query := `SELECT inventory_number, name, manufacturer, start_date, lifeTime FROM "MyDatabase".public.equipment WHERE inventory_number = $1`
	row := r.session.DB.QueryRow(query, id)

	var equipment internal.Equipment
	err := row.Scan(&equipment.InventoryNumber, &equipment.Name, &equipment.Manufacturer, &equipment.StartDate, &equipment.LifeTime)
	if err != nil {
		return nil, err
	}
	return &equipment, nil
}

func (r *equipmentRepository) GetAllEquipment() ([]internal.Equipment, error) {
	query := `SELECT inventory_number, name, manufacturer, start_date, lifeTime FROM "MyDatabase".public.equipment`
	rows, err := r.session.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var equipments []internal.Equipment
	for rows.Next() {
		var equipment internal.Equipment
		err := rows.Scan(&equipment.InventoryNumber, &equipment.Name, &equipment.Manufacturer, &equipment.StartDate, &equipment.LifeTime)
		if err != nil {
			return nil, err
		}
		equipments = append(equipments, equipment)
	}
	return equipments, nil
}

func (r *equipmentRepository) UpdateEquipment(id int, equipment internal.Equipment) error {
	query := `UPDATE "MyDatabase".public.equipment SET name = $1, manufacturer = $2, start_date = $3, lifeTime = $4 WHERE inventory_number = $5`
	_, err := r.session.DB.Exec(query, equipment.Name, equipment.Manufacturer, equipment.StartDate, equipment.LifeTime, id)
	return err
}

func (r *equipmentRepository) DeleteEquipment(id int) error {
	query := `DELETE FROM "MyDatabase".public.equipment WHERE inventory_number = $1`
	_, err := r.session.DB.Exec(query, id)
	return err
}
