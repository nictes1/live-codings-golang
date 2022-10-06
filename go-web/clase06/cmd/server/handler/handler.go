package handler

import (
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nictes1/live-codings-golang/go-web/clase06/internal/products"
	"github.com/nictes1/live-codings-golang/go-web/clase06/pkg/web"
)

type ProductRequest struct {
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

// ListProducts godoc
// @Summary  Show list products
// @Tags     Products
// @Produce  json
// @Param    token  header    string        true  "token"
// @Success  200    {object}  web.Response  "List products"
// @Failure  401    {object}  web.Response  "Unauthorized"
// @Failure  500    {object}  web.Response  "Internal server error "
// @Failure  404    {object}  web.Response  "Not found products"
// @Router   /products [GET]
func (p *Product) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			c.JSON(http.StatusUnauthorized, web.NewResponse(http.StatusUnauthorized, nil, "token inválido"))
			return
		}

		p, err := p.service.GetAll()
		if err != nil {
			c.JSON(http.StatusInternalServerError, web.NewResponse(http.StatusInternalServerError, nil, "ha ocurrido un error inesperado"))
			return
		}

		if len(p) == 0 {
			c.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, "no hay productos registrados"))
			return
		}

		c.JSON(http.StatusOK, web.NewResponse(http.StatusOK, p, ""))
	}
}

// Store Product godoc
// @Summary  Store product
// @Tags     Products
// @Accept   json
// @Produce  json
// @Param    token    header    string          true  "token requerido"
// @Param    product  body      ProductRequest  true  "Product to Store"
// @Success  200      {object}  web.Response
// @Failure  401      {object}  web.Response
// @Failure  400      {object}  web.Response
// @Failure  409      {object}  web.Response
// @Router   /products [POST]
func (p *Product) Store() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			c.JSON(http.StatusUnauthorized, web.NewResponse(http.StatusUnauthorized, nil, "token inválido"))
			return
		}

		var req ProductRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, err.Error()))
			return
		}

		if req.Nombre == "" {
			c.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "campo nombre es requerido"))
			return
		}

		if req.Tipo == "" {
			c.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "campo tipo es requerido"))
			return
		}

		if req.Cantidad <= 0 {
			c.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "campo cantidad es requerido"))
			return
		}

		if req.Precio <= 0 {
			c.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "campo precio es requerido"))
			return
		}

		p, err := p.service.Store(req.Nombre, req.Tipo, req.Cantidad, req.Precio)
		if err != nil {
			c.JSON(http.StatusConflict, web.NewResponse(http.StatusConflict, nil, err.Error()))
			return
		}

		c.JSON(http.StatusOK, p)
	}
}

// UpdateProduct godoc
// @Summary  Update product
// @Tags     Products
// @Accept   json
// @Produce  json
// @Param    id       path      int             true   "Id product"
// @Param    token    header    string          false  "Token"
// @Param    product  body      ProductRequest  true   "Product to update"
// @Success  200      {object}  web.Response
// @Failure  401      {object}  web.Response
// @Failure  400      {object}  web.Response
// @Failure  404      {object}  web.Response
// @Router   /products/{id} [PUT]
func (p *Product) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			c.JSON(http.StatusUnauthorized, web.NewResponse(http.StatusUnauthorized, nil, "token inválido"))
			return
		}

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "Id invalido - "+err.Error()))
			return
		}

		var req ProductRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, err.Error()))
			return
		}

		if req.Nombre == "" {
			c.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "campo nombre es requerido"))
			return
		}

		if req.Tipo == "" {
			c.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "campo tipo es requerido"))
			return
		}

		if req.Cantidad <= 0 {
			c.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "campo cantidad es requerido"))
			return
		}

		if req.Precio <= 0 {
			c.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "campo precio es requerido"))
			return
		}

		p, err := p.service.Update(id, req.Nombre, req.Tipo, req.Cantidad, req.Precio)
		if err != nil {
			c.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, err.Error()))
			return
		}

		c.JSON(http.StatusOK, web.NewResponse(http.StatusOK, p, ""))
	}
}

// Update Name Product godoc
// @Summary      Update name product
// @Tags         Products
// @Accept       json
// @Produce      json
// @Description  This endpoint update field name from an product
// @Param        token  header    string               true  "Token header"
// @Param        id     path      int                  true  "Product Id"
// @Param        name   body      ProductRequestPatch  true  "Product name"
// @Success      200    {object}  web.Response
// @Failure      401    {object}  web.Response
// @Failure      400    {object}  web.Response
// @Failure      404    {object}  web.Response
// @Router       /products/{id} [PATCH]
func (p *Product) UpdateName() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			c.JSON(http.StatusUnauthorized, web.NewResponse(http.StatusUnauthorized, nil, "token inválido"))
			return
		}

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "Id invalido - "+err.Error()))
			return
		}

		var req ProductRequestPatch
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, err.Error()))
			return
		}

		p, err := p.service.UpdateName(id, req.Nombre)
		if err != nil {
			c.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, err.Error()))
			return
		}

		c.JSON(http.StatusOK, web.NewResponse(http.StatusOK, p, ""))
	}
}

/* // Test godoc
// @Produce plain
// @Success 200 {string} string
// @Router /products/test [GET]
func (p *Product) Test() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(http.StatusOK, "Contenido de repuesto en formato text/plain")
	}
} */

// Delete Product
// @Summary  Delete product
// @Tags     Products
// @Param    id     path      int     true  "Product id"
// @Param    token  header    string  true  "Token"
// @Success  204
// @Failure  401    {object}  web.Response
// @Failure  400    {object}  web.Response
// @Failure  404    {object}  web.Response
// @Router   /products/{id} [DELETE]
func (p *Product) Delete(c *gin.Context) {
	token := c.GetHeader("token")
	if token != os.Getenv("TOKEN") {
		c.JSON(http.StatusUnauthorized, web.NewResponse(http.StatusUnauthorized, nil, "token inválido"))
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "Id invalido - "+err.Error()))
		return
	}

	if err := p.service.Delete(id); err != nil {
		c.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, err.Error()))
		return
	}

	c.JSON(http.StatusNoContent, web.NewResponse(http.StatusNoContent, nil, ""))
}
