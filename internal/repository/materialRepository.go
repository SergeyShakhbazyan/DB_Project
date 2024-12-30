package repository

import (
	"ProjectDB/config"
	"ProjectDB/internal"
	"fmt"
)

type MaterialRepository interface {
	CreateMaterial(material internal.Material) error
	GetFilteredMaterials(name, materialType string) ([]internal.Material, error)
	UpdateMaterialPrice(percentage float64) error
	GetMaterialCountByType() (map[string]int, error)
	GetPaginatedMaterials(page, pageSize int) ([]internal.Material, error)
}

type materialRepository struct {
	session *config.Connection
}

func NewMaterialRepository(session *config.Connection) MaterialRepository {
	return &materialRepository{session: session}
}

func (r *materialRepository) CreateMaterial(material internal.Material) error {
	query := `INSERT INTO "MyDatabase".public.material (name, type, unit_price, unit_of_measurement, alternative) VALUES ($1, $2, $3, $4, $5)`
	_, err := r.session.DB.Exec(query, material.Name, material.Type, material.UnitPrice, material.UnitOfMeasurement, material.Alternative)
	if err != nil {
		return fmt.Errorf("failed to insert material: %w", err)
	}
	return nil
}

func (r *materialRepository) GetFilteredMaterials(name, materialType string) ([]internal.Material, error) {
	query := `SELECT * FROM "MyDatabase".public.material WHERE name LIKE $1 AND type = $2`
	rows, err := r.session.DB.Query(query, name, materialType)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var materials []internal.Material
	for rows.Next() {
		var material internal.Material
		if err := rows.Scan(&material.MaterialID, &material.Name, &material.Type, &material.UnitPrice, &material.UnitOfMeasurement, &material.Alternative); err != nil {
			return nil, err
		}
		materials = append(materials, material)
	}
	return materials, nil
}

func (r *materialRepository) UpdateMaterialPrice(percentage float64) error {
	query := `UPDATE "MyDatabase".public.material SET unit_price = unit_price + (unit_price * $1 / 100)`
	_, err := r.session.DB.Exec(query, percentage)
	return err
}

func (r *materialRepository) GetMaterialCountByType() (map[string]int, error) {
	query := `SELECT type, COUNT(*) AS material_count FROM "MyDatabase".public.material GROUP BY type`
	rows, err := r.session.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to get material count by type: %w", err)
	}
	defer rows.Close()

	materialCounts := make(map[string]int)
	for rows.Next() {
		var materialType string
		var count int
		if err := rows.Scan(&materialType, &count); err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		materialCounts[materialType] = count
	}
	return materialCounts, nil
}

func (r *materialRepository) GetPaginatedMaterials(page, pageSize int) ([]internal.Material, error) {
	query := `SELECT * FROM "MyDatabase".public.material ORDER BY material_id LIMIT $1 OFFSET $2`
	offset := (page - 1) * pageSize

	rows, err := r.session.DB.Query(query, pageSize, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to execute pagination query: %w", err)
	}
	defer rows.Close()

	var materials []internal.Material
	for rows.Next() {
		var material internal.Material
		if err := rows.Scan(&material.MaterialID, &material.Name, &material.Type, &material.UnitPrice, &material.UnitOfMeasurement, &material.Alternative); err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		materials = append(materials, material)
	}
	return materials, nil
}
