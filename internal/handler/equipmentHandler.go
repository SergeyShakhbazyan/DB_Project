package handler

import (
	"ProjectDB/internal"
	"ProjectDB/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
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

func (h *EquipmentHandler) GetAllEquipment(c *gin.Context) {
	equipment, err := h.service.GetAllEquipment()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, equipment)
}
