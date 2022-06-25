package main

import (
	controller "maquina1995/webservice/controller"

	"maquina1995/webservice/docs"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

const basePath = "/cancion/"

// @contact.name MaQuiNa1995
// @contact.url https://github.com/MaQuiNa1995
// @contact.email maquina1995@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	r := gin.New()

	docs.SwaggerInfo.Title = "MaQuiNa1995 WebService En Golang"
	docs.SwaggerInfo.Description = "Ejemplo de Crud con Gorm"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8081"
	docs.SwaggerInfo.BasePath = ""
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	// Endpoint para la documentación de swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	cancion := r.Group(basePath)
	{
		cancion.POST("", controller.Create)
		cancion.GET("", controller.FindAll)
		cancion.GET(":id", controller.FindById)
		cancion.PATCH("", controller.Update)
		cancion.DELETE(":id", controller.Delete)
	}
	r.Run("localhost:8081")
}
