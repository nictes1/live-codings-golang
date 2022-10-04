package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Empleado struct {
	Nombre   string `form:"name" json:"name"`
	Password string `form:"password" json:"password"`
	Id       string `form:"id" json:"id"`
	Activo   string `form:"active" json:"active" binding:"required"`
}

// "Nombre" : "nicolas"
// "name": "nicolas"

func main() {
	router := gin.Default()

	router.POST("/auth", AutorizarEmpleado)

	router.Run()
}

func AutorizarEmpleado(ctx *gin.Context) {

	var empleado Empleado

	if err := ctx.ShouldBindJSON(&empleado); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if empleado.Nombre != "Nicolas" || empleado.Password != "Password" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"status": "no esta autorizado"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "autorizado"})
}
