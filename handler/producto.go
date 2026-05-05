package handler

import (
	"mi-api-go/domain"
	"mi-api-go/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductoHandler struct {
	service *service.ProductosService
}

func NewProductoHandler(s *service.ProductosService) *ProductoHandler {
	return &ProductoHandler{service: s}
}

func (h *ProductoHandler) GetAll(c *gin.Context) {
	productos, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, productos)
}

func (h *ProductoHandler) GetById(c *gin.Context) {
	id := c.Param("id")
	producto, err := h.service.GetById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, producto)
}

func (h *ProductoHandler) Create(c *gin.Context) {
	var p domain.Producto
	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	nuevo, err := h.service.Create(p)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, nuevo)
}

func (h *ProductoHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var p domain.Producto
	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	actualizado, err := h.service.Update(id, p)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, actualizado)
}

func (h *ProductoHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	if err := h.service.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
