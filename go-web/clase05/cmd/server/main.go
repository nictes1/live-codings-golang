package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nictes1/live-codings-golang/go-web/clase05/cmd/server/handler"
	"github.com/nictes1/live-codings-golang/go-web/clase05/internal/products"
)

func main() {
	repo := products.NewRepository()
	service := products.NewService(repo)
	p := handler.NewProduct(service)

	r := gin.Default()
	pr := r.Group("products")
	pr.POST("/", p.Store())
	pr.GET("/", p.GetAll())
	pr.PUT("/:id", p.Update())
	pr.PATCH("/:id", p.UpdateName())
	pr.DELETE("/:id", p.Delete)

	if err := r.Run(); err != nil {
		panic(err)
	}
}
