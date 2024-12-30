package handler

import (
	"ProjectDB/internal"
	"ProjectDB/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type EquipmentHandler struct {
	service service.EquipmentService
}

func NewEquipmentHandler(service service.EquipmentService) *EquipmentHandler {
	return &EquipmentHandler{service: service}
}

func (h *EquipmentHandler) CreateEquipment(c *gin.Context) {
	var equipment internal.Equipment
	if err := c.ShouldBindJSON(&equipment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.service.CreateEquipment(equipment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Equipment created successfully"})
}

func (h *EquipmentHandler) GetFilteredEquipment(c *gin.Context) {
	name := c.Query("name")
	manufacturer := c.Query("manufacturer")

	if name == "" || manufacturer == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Name and Manufacturer parameters are required"})
		return
	}

	equipments, err := h.service.GetFilteredEquipment(name, manufacturer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, equipments)
}

func (h *EquipmentHandler) GetEquipmentWithMaterials(c *gin.Context) {
	results, err := h.service.GetEquipmentWithMaterials()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, results)
}

func (h *EquipmentHandler) UpdateEquipmentLifeTime(c *gin.Context) {
	date := c.Query("start_date")
	if date == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "start_date parameter is required"})
		return
	}

	err := h.service.UpdateEquipmentLifeTime(date)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Equipment updated successfully"})
}

func (h *EquipmentHandler) DeleteEquipment(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Equipment ID"})
		return
	}

	err = h.service.DeleteEquipment(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Equipment deleted successfully"})
}
