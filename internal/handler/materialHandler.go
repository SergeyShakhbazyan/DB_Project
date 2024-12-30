package handler

import (
	"ProjectDB/internal"
	"ProjectDB/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
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

func (h *MaterialHandler) GetAllMaterials(c *gin.Context) {
	materials, err := h.service.GetAllMaterials()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, materials)
}
