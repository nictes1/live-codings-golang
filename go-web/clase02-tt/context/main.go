package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.GET("/hola", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Hola :) \n",
		})
	})

	router.POST("/hola", func(context *gin.Context) {
		//El body, header y method están contenidos en el contexto de gin.
		body := context.Request.Body
		header := context.Request.Header
		metodo := context.Request.Method

		fmt.Println("¡He recibido algo!")
		fmt.Printf("\tMetodo: %s\n", metodo)
		fmt.Printf("\tContenido del header:\n")

		for key, value := range header {
			fmt.Printf("\t\t%s -> %s\n", key, value)
		}
		fmt.Printf("\tEl body es un io.ReadCloser:(%s), y para trabajar con el vamos a tener queleerlo luego\n", body)
		fmt.Println("¡Yay!")
		context.String(http.StatusOK, "¡Lo recibí!") //Respondemos al cliente con 200 OK y un mensaje.

	})

	router.Run()

}
