package internal

type Equipment struct {
	Name            string `json:"name"`
	Manufacturer    string `json:"manufacturer"`
	StartDate       string `json:"start_date"`
	LifeTime        int    `json:"lifeTime"`
	InventoryNumber int    `json:"inventory"`
}

type Material struct {
	MaterialID        int     `json:"material_id"`
	Name              string  `json:"name"`
	Type              string  `json:"type"`
	UnitPrice         float64 `json:"unit_price"`
	UnitOfMeasurement string  `json:"unit_of_measurement"`
	Alternative       string  `json:"alternative"`
}

type ProductSpecification struct {
	SpecificationID    int    `json:"specification_id"`
	Name               string `json:"name"`
	ProductionDuration int    `json:"production_duration"`
	EquipmentID        int    `json:"equipment_id"`
	MaterialID         int    `json:"material_id"`
	Quantity           int    `json:"quantity"`
}
