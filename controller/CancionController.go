package controller

import (
	"maquina1995/webservice/model"
	"maquina1995/webservice/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {

	var dto model.CancionCreate
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	entity := model.CancionEntity{
		Nombre:   dto.Nombre,
		Duracion: dto.Duracion,
		Genero:   dto.Genero,
	}

	repository.Create(&entity)

	c.JSON(http.StatusAccepted, nil)
}

func FindById(c *gin.Context) {

	cancion, err := repository.FindById(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusNoContent, nil)
	}

	c.JSON(http.StatusOK, gin.H{"data": cancion})

}

func FindAll(c *gin.Context) {

	canciones := repository.FindAll()
	c.JSON(http.StatusOK, gin.H{"data": canciones})
}

func Update(c *gin.Context) {

}

func Delete(c *gin.Context) {

}
