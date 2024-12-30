package repository

import (
	"ProjectDB/config"
	"ProjectDB/internal"
	"fmt"
)

type EquipmentRepository interface {
	CreateEquipment(equipment internal.Equipment) error
	DeleteEquipment(id int) error
	GetFilteredEquipment(name, manufacturer string) ([]internal.Equipment, error)
	GetEquipmentWithMaterials() ([]map[string]interface{}, error)
	UpdateEquipmentLifeTime(startDate string) error
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

func (r *equipmentRepository) GetFilteredEquipment(name, manufacturer string) ([]internal.Equipment, error) {
	query := `SELECT * FROM "MyDatabase".public.equipment WHERE name LIKE $1 AND manufacturer = $2`
	rows, err := r.session.DB.Query(query, name, manufacturer)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var equipments []internal.Equipment
	for rows.Next() {
		var equipment internal.Equipment
		if err := rows.Scan(&equipment.InventoryNumber, &equipment.Name, &equipment.Manufacturer, &equipment.StartDate, &equipment.LifeTime); err != nil {
			return nil, err
		}
		equipments = append(equipments, equipment)
	}
	return equipments, nil
}

func (r *equipmentRepository) GetEquipmentWithMaterials() ([]map[string]interface{}, error) {
	query := `
		SELECT ps.name AS product_specification, e.name AS equipment, m.name AS material
		FROM "MyDatabase".public.product_specification ps
		JOIN "MyDatabase".public.equipment e ON ps.equipment_id = e.inventory_number
		JOIN "MyDatabase".public.material m ON ps.material_id = m.material_id
	`
	rows, err := r.session.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []map[string]interface{}
	for rows.Next() {
		var productSpec, equipment, material string
		if err := rows.Scan(&productSpec, &equipment, &material); err != nil {
			return nil, err
		}
		results = append(results, map[string]interface{}{
			"product_specification": productSpec,
			"equipment":             equipment,
			"material":              material,
		})
	}
	return results, nil
}

func (r *equipmentRepository) UpdateEquipmentLifeTime(startDate string) error {
	query := `UPDATE "MyDatabase".public.equipment SET lifeTime = lifeTime + 5 WHERE start_date < $1`
	_, err := r.session.DB.Exec(query, startDate)
	return err
}

func (r *equipmentRepository) DeleteEquipment(id int) error {
	query := `DELETE FROM "MyDatabase".public.equipment WHERE inventory_number = $1`
	_, err := r.session.DB.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete equipment: %w", err)
	}
	return nil
}
