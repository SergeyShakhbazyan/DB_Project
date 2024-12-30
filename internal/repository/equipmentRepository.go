package repository

import (
	"ProjectDB/config"
	"ProjectDB/internal"
)

type EquipmentRepository interface {
	CreateEquipment(equipment internal.Equipment) error
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

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var equipments []internal.Equipment
	for rows.Next() {
		var equipment internal.Equipment
			return nil, err
		}
		equipments = append(equipments, equipment)
	}
	return equipments, nil
}

	return err
}

func (r *equipmentRepository) DeleteEquipment(id int) error {
	query := `DELETE FROM "MyDatabase".public.equipment WHERE inventory_number = $1`
	_, err := r.session.DB.Exec(query, id)
}
