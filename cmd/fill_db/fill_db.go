package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const baseURL = "http://localhost:8000/api"

func addEquipment() {
	url := fmt.Sprintf("%s/equipment", baseURL)
	for i := 1; i <= 100; i++ {
		payload := map[string]interface{}{
			"name":         fmt.Sprintf("Equipment %d", i),
			"manufacturer": fmt.Sprintf("Manufacturer %d", i),
			"start_date":   "2023-01-01",
			"lifeTime":     i,
		}
		data, _ := json.Marshal(payload)
		resp, err := http.Post(url, "application/json", bytes.NewBuffer(data))
		if err != nil {
			fmt.Printf("Failed to create equipment: %v\n", err)
			continue
		}
		fmt.Printf("Response for equipment %d: %s\n", i, resp.Status)
	}
}

func addMaterial() {
	url := fmt.Sprintf("%s/material", baseURL)
	for i := 1; i <= 100; i++ {
		payload := map[string]interface{}{
			"name":                fmt.Sprintf("Material %d", i),
			"type":                "Type A",
			"unit_price":          100 + float64(i),
			"unit_of_measurement": "kg",
			"alternative":         fmt.Sprintf("Alternative %d", i),
		}
		data, _ := json.Marshal(payload)
		resp, err := http.Post(url, "application/json", bytes.NewBuffer(data))
		if err != nil {
			fmt.Printf("Failed to create material: %v\n", err)
			continue
		}
		fmt.Printf("Response for material %d: %s\n", i, resp.Status)
	}
}
func addProductSpecifications() {
	url := fmt.Sprintf("%s/product-specification", baseURL)
	for i := 1; i <= 100; i++ {
		payload := map[string]interface{}{
			"name":                fmt.Sprintf("Product Specification %d", i),
			"production_duration": 10 + i,
			"equipment_id":        (i % 50) + 1,
			"material_id":         (i % 50) + 1,
			"quantity":            100 + i,
		}
		data, _ := json.Marshal(payload)
		resp, err := http.Post(url, "application/json", bytes.NewBuffer(data))
		if err != nil {
			fmt.Printf("Failed to create product specification: %v\n", err)
			continue
		}
		fmt.Printf("Response for product specification %d: %s\n", i, resp.Status)
	}
}

func main() {
	fmt.Println("Filling database with test data...")
	addEquipment()
	addMaterial()
	addProductSpecifications()
	fmt.Println("Database population completed.")
}
