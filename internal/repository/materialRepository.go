package repository

import (
	"ProjectDB/config"
	"ProjectDB/internal"
	"database/sql"
	"fmt"
)

type MaterialRepository interface {
	CreateMaterial(material internal.Material) error
	GetMaterialByID(id int) (*internal.Material, error)
	GetAllMaterials() ([]internal.Material, error)
	UpdateMaterial(id int, material internal.Material) error
	DeleteMaterial(id int) error
}

type materialRepository struct {
	session *config.Connection
}

func NewMaterialRepository(session *config.Connection) MaterialRepository {
	return &materialRepository{session: session}
}

func (r *materialRepository) CreateMaterial(material internal.Material) error {
	query := `INSERT INTO "MyDatabase".public.material (material_id, name, type, unit_price, unit_of_measurement, alternative) VALUES ($1, $2, $3, $4, $5)`
	_, err := r.session.DB.Exec(query, material.MaterialID, material.Name, material.Type, material.UnitPrice, material.UnitOfMeasurement, material.Alternative)
	if err != nil {
		return fmt.Errorf("failed to insert material: %w", err)
	}
	return nil
}

func (r *materialRepository) GetMaterialByID(id int) (*internal.Material, error) {
	query := `SELECT material_id, name, type, unit_price, unit_of_measurement, alternative FROM "MyDatabase".public.material WHERE material_id = $1`
	row := r.session.DB.QueryRow(query, id)

	var material internal.Material
	err := row.Scan(&material.MaterialID, &material.Name, &material.Type, &material.UnitPrice, &material.UnitOfMeasurement, &material.Alternative)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to get material by ID: %w", err)
	}
	return &material, nil
}

func (r *materialRepository) GetAllMaterials() ([]internal.Material, error) {
	query := `SELECT material_id, name, type, unit_price, unit_of_measurement, alternative FROM "MyDatabase".public.material`
	rows, err := r.session.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch all materials: %w", err)
	}
	defer rows.Close()

	var materials []internal.Material
	for rows.Next() {
		var material internal.Material
		err := rows.Scan(&material.MaterialID, &material.Name, &material.Type, &material.UnitPrice, &material.UnitOfMeasurement, &material.Alternative)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		materials = append(materials, material)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %w", err)
	}
	return materials, nil
}

func (r *materialRepository) UpdateMaterial(id int, material internal.Material) error {
	query := `UPDATE "MyDatabase".public.material SET name = $1, type = $2, unit_price = $3, unit_of_measurement = $4, alternative = $5 WHERE material_id = $6`
	_, err := r.session.DB.Exec(query, material.Name, material.Type, material.UnitPrice, material.UnitOfMeasurement, material.Alternative, id)
	if err != nil {
		return fmt.Errorf("failed to update material: %w", err)
	}
	return nil
}

func (r *materialRepository) DeleteMaterial(id int) error {
	query := `DELETE FROM "MyDatabase".public.material WHERE material_id = $1`
	_, err := r.session.DB.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete material: %w", err)
	}
	return nil
}
