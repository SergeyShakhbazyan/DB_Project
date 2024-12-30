package handler

import (
	"ProjectDB/internal"
	"ProjectDB/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
<<<<<<< HEAD
	"strconv"
=======
>>>>>>> main
)

type ProductSpecificationHandler struct {
	service service.ProductSpecificationService
}

func NewProductSpecificationHandler(service service.ProductSpecificationService) *ProductSpecificationHandler {
	return &ProductSpecificationHandler{service: service}
}

func (h *ProductSpecificationHandler) CreateProductSpecification(c *gin.Context) {
	var spec internal.ProductSpecification
	if err := c.ShouldBindJSON(&spec); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.service.CreateProductSpecification(spec); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Product specification created successfully"})
}

<<<<<<< HEAD
func (h *ProductSpecificationHandler) GetSpecificationsByEquipment(c *gin.Context) {
	equipmentIDStr := c.Param("equipment_id")
	equipmentID, err := strconv.Atoi(equipmentIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid equipment_id"})
		return
	}

	specs, err := h.service.GetSpecificationsByEquipment(equipmentID)
=======
func (h *ProductSpecificationHandler) GetAllProductSpecifications(c *gin.Context) {
	specs, err := h.service.GetAllProductSpecifications()
>>>>>>> main
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
<<<<<<< HEAD

	c.JSON(http.StatusOK, specs)
}

func (h *ProductSpecificationHandler) UpdateProductionDuration(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid specification ID"})
		return
	}

	var request struct {
		NewDuration int `json:"new_duration"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.UpdateProductionDuration(id, request.NewDuration); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Production duration updated successfully"})
}

func (h *ProductSpecificationHandler) GetProductionCountByEquipment(c *gin.Context) {
	productionCounts, err := h.service.GetProductionCountByEquipment()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, productionCounts)
}
=======
	c.JSON(http.StatusOK, specs)
}
>>>>>>> main
