package app

import (
	"ProjectDB/config"
	"ProjectDB/internal/handler"
	"ProjectDB/internal/repository"
	"ProjectDB/internal/service"
	"github.com/gin-gonic/gin"
	"log"
)

type App struct {
	Router   *gin.Engine
	Database *config.Connection
}

func (a *App) Initialize() {
	a.Router = gin.Default()
	a.Database = config.NewConnection()

	equipmentRepo := repository.NewEquipmentRepository(a.Database)
	equipmentService := service.NewEquipmentService(equipmentRepo)
	equipmentHandler := handler.NewEquipmentHandler(equipmentService)

	materialRepo := repository.NewMaterialRepository(a.Database)
	materialService := service.NewMaterialService(materialRepo)
	materialHandler := handler.NewMaterialHandler(materialService)

	productSpecification := repository.NewProductSpecificationRepository(a.Database)
	productSpecificationService := service.NewProductSpecificationService(productSpecification)
	productSpecificationHandler := handler.NewProductSpecificationHandler(productSpecificationService)

	a.setRoutersForEquipments(equipmentHandler)
	a.setRoutersForMaterials(materialHandler)
	a.setRoutersForProductSpecifications(productSpecificationHandler)
}

func (a *App) Run(port string) {
	if err := a.Router.Run(":" + port); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
