package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	//Cada vez que llamamos a GET y le pasamos una ruta, definimos un nuevo endpoint.
	router.GET("/", HandlerRaiz)
	router.GET("/gophers", HandlerGophers)
	router.GET("/gophers/info", HandlerGetInfo)
	router.GET("/about", HandlerAbout)
	router.Run(":8080")
}

func HandlerRaiz(ctx *gin.Context) {
	ctx.String(http.StatusOK, "Bienvenidos/as a la pagina gopher \n")
}
func HandlerGophers(ctx *gin.Context) {
	type response struct {
		Gophers []string `json:"gophers"`
	}
	r := response{Gophers: []string{"gopher-1", "gopher-2 \n"}}
	ctx.JSON(http.StatusOK, r)
}
func HandlerGetInfo(ctx *gin.Context) {
	ctx.String(http.StatusOK, "Gophher: Mascota de Golang \n")
}
func HandlerAbout(ctx *gin.Context) {
	ctx.String(http.StatusOK, "Informacion \n")
}
