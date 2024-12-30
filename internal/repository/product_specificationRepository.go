package repository

import (
	"ProjectDB/config"
	"ProjectDB/internal"
<<<<<<< HEAD
=======
	"database/sql"
>>>>>>> main
	"fmt"
)

type ProductSpecificationRepository interface {
	CreateProductSpecification(spec internal.ProductSpecification) error
<<<<<<< HEAD
	GetSpecificationsByEquipment(equipmentID int) ([]internal.ProductSpecification, error)
	UpdateProductionDuration(id int, newDuration int) error
	GetProductionCountByEquipment() (map[int]int, error)
=======
	GetProductSpecificationByID(id int) (*internal.ProductSpecification, error)
	GetAllProductSpecifications() ([]internal.ProductSpecification, error)
	UpdateProductSpecification(id int, spec internal.ProductSpecification) error
	DeleteProductSpecification(id int) error
>>>>>>> main
}

type productSpecificationRepository struct {
	session *config.Connection
}

func NewProductSpecificationRepository(session *config.Connection) ProductSpecificationRepository {
	return &productSpecificationRepository{session: session}
}

func (r *productSpecificationRepository) CreateProductSpecification(spec internal.ProductSpecification) error {
	query := `
		INSERT INTO "MyDatabase".public.product_specification (name, production_duration, equipment_id, material_id, quantity)
		VALUES ($1, $2, $3, $4, $5)
	`
	_, err := r.session.DB.Exec(query, spec.Name, spec.ProductionDuration, spec.EquipmentID, spec.MaterialID, spec.Quantity)
	if err != nil {
		return fmt.Errorf("failed to insert product specification: %w", err)
	}
	return nil
}

<<<<<<< HEAD
func (r *productSpecificationRepository) GetSpecificationsByEquipment(equipmentID int) ([]internal.ProductSpecification, error) {
	query := `SELECT * FROM "MyDatabase".public.product_specification WHERE equipment_id = $1`
	rows, err := r.session.DB.Query(query, equipmentID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var specifications []internal.ProductSpecification
	for rows.Next() {
		var spec internal.ProductSpecification
		if err := rows.Scan(&spec.SpecificationID, &spec.Name, &spec.ProductionDuration, &spec.EquipmentID, &spec.MaterialID, &spec.Quantity); err != nil {
			return nil, err
		}
		specifications = append(specifications, spec)
	}
	return specifications, nil
}

func (r *productSpecificationRepository) UpdateProductionDuration(id int, newDuration int) error {
	query := `UPDATE "MyDatabase".public.product_specification SET production_duration = $1 WHERE specification_id = $2`
	_, err := r.session.DB.Exec(query, newDuration, id)
	return err
}

func (r *productSpecificationRepository) GetProductionCountByEquipment() (map[int]int, error) {
	query := `SELECT equipment_id, COUNT(*) AS production_count FROM "MyDatabase".public.product_specification GROUP BY equipment_id`
	rows, err := r.session.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to get production count by equipment: %w", err)
	}
	defer rows.Close()

	productionCounts := make(map[int]int)
	for rows.Next() {
		var equipmentID, count int
		if err := rows.Scan(&equipmentID, &count); err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		productionCounts[equipmentID] = count
	}
	return productionCounts, nil
=======
func (r *productSpecificationRepository) GetProductSpecificationByID(id int) (*internal.ProductSpecification, error) {
	query := `
		SELECT specification_id, name, production_duration, equipment_id, material_id, quantity
		FROM "MyDatabase".public.product_specification
		WHERE specification_id = $1
	`
	row := r.session.DB.QueryRow(query, id)

	var spec internal.ProductSpecification
	err := row.Scan(&spec.SpecificationID, &spec.Name, &spec.ProductionDuration, &spec.EquipmentID, &spec.MaterialID, &spec.Quantity)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Not found
		}
		return nil, fmt.Errorf("failed to get product specification by ID: %w", err)
	}
	return &spec, nil
}

func (r *productSpecificationRepository) GetAllProductSpecifications() ([]internal.ProductSpecification, error) {
	query := `
		SELECT specification_id, name, production_duration, equipment_id, material_id, quantity
		FROM "MyDatabase".public.product_specification
	`
	rows, err := r.session.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch all product specifications: %w", err)
	}
	defer rows.Close()

	var specs []internal.ProductSpecification
	for rows.Next() {
		var spec internal.ProductSpecification
		err := rows.Scan(&spec.SpecificationID, &spec.Name, &spec.ProductionDuration, &spec.EquipmentID, &spec.MaterialID, &spec.Quantity)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		specs = append(specs, spec)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %w", err)
	}
	return specs, nil
}

func (r *productSpecificationRepository) UpdateProductSpecification(id int, spec internal.ProductSpecification) error {
	query := `
		UPDATE "MyDatabase".public.product_specification
		SET name = $1, production_duration = $2, equipment_id = $3, material_id = $4, quantity = $5
		WHERE specification_id = $6
	`
	_, err := r.session.DB.Exec(query, spec.Name, spec.ProductionDuration, spec.EquipmentID, spec.MaterialID, spec.Quantity, id)
	if err != nil {
		return fmt.Errorf("failed to update product specification: %w", err)
	}
	return nil
}

func (r *productSpecificationRepository) DeleteProductSpecification(id int) error {
	query := `
		DELETE FROM "MyDatabase".public.product_specification
		WHERE specification_id = $1
	`
	_, err := r.session.DB.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete product specification: %w", err)
	}
	return nil
>>>>>>> main
}
