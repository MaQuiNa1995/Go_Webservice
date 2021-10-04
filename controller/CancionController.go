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

	entity, err := repository.FindById(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusNoContent, nil)
	}

	c.JSON(http.StatusOK, gin.H{"data": entity})

}

func FindAll(c *gin.Context) {

	dtos := repository.FindAll()
	c.JSON(http.StatusOK, gin.H{"data": dtos})
}

func Update(c *gin.Context) {

	var dto model.CancionUpdateDto

	err := c.ShouldBindJSON(&dto)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = repository.Update(&dto)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, nil)
}

func Delete(c *gin.Context) {

	err := repository.Delete(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusNoContent, nil)
	}

	c.JSON(http.StatusAccepted, nil)
}
