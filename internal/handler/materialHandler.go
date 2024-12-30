package handler

import (
	"ProjectDB/internal"
	"ProjectDB/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type MaterialHandler struct {
	service service.MaterialService
}

func NewMaterialHandler(service service.MaterialService) *MaterialHandler {
	return &MaterialHandler{service: service}
}

func (h *MaterialHandler) CreateMaterial(c *gin.Context) {
	var material internal.Material
	if err := c.ShouldBindJSON(&material); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.service.CreateMaterial(material); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Material created successfully"})
}

func (h *MaterialHandler) GetFilteredMaterials(c *gin.Context) {
	name := c.Query("name")
	materialType := c.Query("type")

	if name == "" || materialType == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Name and Type parameters are required"})
		return
	}

	materials, err := h.service.GetFilteredMaterials(name, materialType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, materials)
}

func (h *MaterialHandler) UpdateMaterialPrice(c *gin.Context) {
	percentageStr := c.Query("percentage")
	if percentageStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Percentage parameter is required"})
		return
	}

	percentage, err := strconv.ParseFloat(percentageStr, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid percentage value"})
		return
	}

	if err := h.service.UpdateMaterialPrice(percentage); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Material prices updated successfully"})
}

func (h *MaterialHandler) GetMaterialCountByType(c *gin.Context) {
	materialCounts, err := h.service.GetMaterialCountByType()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, materialCounts)
}
