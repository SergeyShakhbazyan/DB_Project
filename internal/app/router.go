package app

<<<<<<< HEAD
import "ProjectDB/internal/handler"

func (a *App) setRoutersForEquipments(equipmentHandler *handler.EquipmentHandler) {
	a.Router.POST("/api/equipment", equipmentHandler.CreateEquipment)
	a.Router.GET("/api/equipment/filter", equipmentHandler.GetFilteredEquipment)
	a.Router.GET("/api/equipment/materials", equipmentHandler.GetEquipmentWithMaterials)
	a.Router.PUT("/api/equipment/update-lifetime", equipmentHandler.UpdateEquipmentLifeTime)
	a.Router.DELETE("/api/equipment/:id", equipmentHandler.DeleteEquipment)
}

func (a *App) setRoutersForMaterials(materialHandler *handler.MaterialHandler) {
	a.Router.POST("/api/material", materialHandler.CreateMaterial)
	a.Router.GET("/api/material", materialHandler.GetFilteredMaterials)
	a.Router.PUT("/api/material/update-price", materialHandler.UpdateMaterialPrice)
}

func (a *App) setRoutersForProductSpecifications(productSpecificationHandler *handler.ProductSpecificationHandler) {
	a.Router.POST("/api/product-specification", productSpecificationHandler.CreateProductSpecification)
	a.Router.GET("/api/product-specification/equipment/:id", productSpecificationHandler.GetSpecificationsByEquipment)
	a.Router.PUT("/api/product-specification/:id", productSpecificationHandler.UpdateProductionDuration)
=======
func (a *App) initializeRoutes() {

>>>>>>> main
}
