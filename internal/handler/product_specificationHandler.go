package handler

import (
	"ProjectDB/internal"
	"ProjectDB/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
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

func (h *ProductSpecificationHandler) GetAllProductSpecifications(c *gin.Context) {
	specs, err := h.service.GetAllProductSpecifications()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, specs)
}
