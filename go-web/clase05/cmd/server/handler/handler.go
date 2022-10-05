package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nictes1/live-codings-golang/go-web/clase05/internal/products"
)

type ProductRequest struct {
	Id       int     `json:"id"`
	Nombre   string  `json:"nombre" binding:"required"`
	Tipo     string  `json:"tipo" binding:"required"`
	Cantidad int     `json:"cantidad" binding:"required"`
	Precio   float64 `json:"precio" binding:"required"`
}

type ProductRequestPatch struct {
	Nombre string `json:"nombre" binding:"required"`
}

type Product struct {
	service products.Service
}

func NewProduct(s products.Service) *Product {
	return &Product{service: s}
}

func (p *Product) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token != "123456" || token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "token invalido"})
			return
		}

		p, err := p.service.GetAll()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if len(p) == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "no hay products registrados"})
			return
		}

		c.JSON(http.StatusOK, p)
	}
}

func (p *Product) Store() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token != "123456" || token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "token invalido"})
			return
		}

		var req ProductRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if req.Nombre == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "campo nombre es requerido"})
			return
		}

		if req.Tipo == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "campo tipo es requerido"})
			return
		}

		if req.Cantidad <= 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "campo cantidad es requerido"})
			return
		}

		if req.Precio <= 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "campo precio es requerido"})
			return
		}

		p, err := p.service.Store(req.Nombre, req.Tipo, req.Cantidad, req.Precio)
		if err != nil {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, p)
	}
}

func (p *Product) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token != "123456" || token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "token invalido"})
			return
		}

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Id invalido - " + err.Error()})
			return
		}

		var req ProductRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if req.Nombre == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "campo nombre es requerido"})
			return
		}

		if req.Tipo == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "campo tipo es requerido"})
			return
		}

		if req.Cantidad <= 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "campo cantidad es requerido"})
			return
		}

		if req.Precio <= 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "campo precio es requerido"})
			return
		}

		p, err := p.service.Update(id, req.Nombre, req.Tipo, req.Cantidad, req.Precio)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, p)
	}
}

func (p *Product) UpdateName() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token != "123456" || token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "token invalido"})
			return
		}

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Id invalido - " + err.Error()})
			return
		}

		var req ProductRequestPatch
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		p, err := p.service.UpdateName(id, req.Nombre)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, p)
	}
}

func (p *Product) Delete(c *gin.Context) {
	token := c.GetHeader("token")
	if token != "123456" || token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "token invalido"})
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Id invalido - " + err.Error()})
		return
	}

	if err := p.service.Delete(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "algo"})

}
