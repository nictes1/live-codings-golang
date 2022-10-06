package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/nictes1/live-codings-golang/go-web/clase06/cmd/server/handler"
	"github.com/nictes1/live-codings-golang/go-web/clase06/docs"
	"github.com/nictes1/live-codings-golang/go-web/clase06/internal/products"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Bootcamp Go Wave 6 - API
// @version         1.0
// @description     This is a simple API development by Digital House's team.
// @termsOfService  https://developers.mercadolibre.com.ar/es_ar/terminos-y-condiciones

// @contact.name   API Support Digital House
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1
func main() {
	loadEnd()
	repo := products.NewRepository()
	service := products.NewService(repo)
	p := handler.NewProduct(service)

	r := gin.Default()
	api := r.Group("/api/v1")

	// Documentación swagger
	docs.SwaggerInfo.Host = os.Getenv("HOST")
	api.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Middlewares
	//api.Use(handler.MiddlewareUno)
	pr := api.Group("products")
	{
		pr.GET("/", handler.MiddlewareList(p.GetAll())...)
		pr.POST("/", p.Store())
		pr.PUT("/:id", p.Update())
		pr.PATCH("/:id", p.UpdateName())
		pr.DELETE("/:id", p.Delete)
		//pr.GET("/test", p.Test())
	}

	if err := r.Run(); err != nil {
		panic(err)
	}
}

func loadEnd() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("No se pudo cargar las variables de entorno - error: ", err)
	}
}
