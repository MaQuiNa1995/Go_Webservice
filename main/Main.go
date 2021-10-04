package main

import (
	controller "maquina1995/webservice/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world"})
	})

	basepath := "/canciones"

	r.POST(basepath, controller.Create)
	r.GET(basepath, controller.FindAll)

	r.Run("localhost:8081")
}