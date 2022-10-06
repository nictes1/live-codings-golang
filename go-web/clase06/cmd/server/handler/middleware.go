package handler

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/nictes1/live-codings-golang/go-web/clase06/pkg/web"
)

func MiddlewareUno(c *gin.Context) {
	log.Println("Este es el primer middleware")
	mwdUno := c.GetHeader("token_uno")
	if mwdUno != "middlewareUno" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "falló en la primer capa de middleware"})
		return
	}
	c.Next()
}

func MiddlewareSegundo(c *gin.Context) {
	log.Println("Este es el segundo middleware")
	mwdTwo := c.GetHeader("token_dos")
	if mwdTwo != "middlewareDos" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "falló en la segunda capa de middleware"})
		return
	}

	c.Next()
}

func MiddlewareList(f gin.HandlerFunc) []gin.HandlerFunc {
	list := []gin.HandlerFunc{
		MiddlewareUno,
		MiddlewareSegundo,
	}

	list = append(list, f)
	return list
}

func AuthMiddlewareToken() gin.HandlerFunc {
	tokenAPI := os.Getenv("TOKEN")
	if tokenAPI == "" {
		log.Fatal("No se ha registrado el token para la API")
	}

	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if tokenAPI != token {
			c.AbortWithStatusJSON(http.StatusForbidden, web.NewResponse(http.StatusForbidden, nil, "Forbidden 403"))
			//c.AbortWithStatusJSON(http.StatusUnauthorized, web.NewResponse(http.StatusUnauthorized, nil, "token inválido"))
			return
		}

		c.Next()
	}
}
