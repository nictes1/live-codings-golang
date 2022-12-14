package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/nictes1/live-codings-golang/go-web/clase04/cmd/server/handler"
	"github.com/nictes1/live-codings-golang/go-web/clase04/internal/products"
	"github.com/nictes1/live-codings-golang/go-web/clase04/pkg/store"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("error al intentar cargar archivo .env")
	}
	db := store.New(store.FileType, "./products.json")
	repo := products.NewRepository(db)
	service := products.NewService(repo)
	p := handler.NewProduct(service)

	r := gin.Default()

	pr := r.Group("/products")
	pr.POST("/", p.Store())
	pr.GET("/", p.GetAll())
	pr.PUT("/:id", p.Update())
	r.Run()
}
