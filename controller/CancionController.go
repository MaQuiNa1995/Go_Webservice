package controller

import (
	"log"
	"maquina1995/webservice/model"
	"maquina1995/webservice/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

const id = "id"

// Create godoc
// @Summary Crea un nuevo registro en Base de datos
// @Produce json
// @Param cancion body model.CancionCreateDto true "Dto con los datos necesarios para crear una entidad"
// @Success 202
// @Router /cancion [post]
func Create(c *gin.Context) {

	var dto model.CancionCreateDto
	if err := c.ShouldBindJSON(&dto); err != nil {
		log.Println("Se ha producido un error: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	entity := model.CancionEntity{
		Nombre:   dto.Nombre,
		Duracion: dto.Duracion,
		Genero:   dto.Genero,
	}
	repository.Create(&entity)
	c.JSON(http.StatusAccepted, &entity.Id)
}

// FindById godoc
// @Summary Buscar por Id en Base de datos
// @Produce json
// @Param id path string true "Id de la entidad a buscar"
// @Success 200 {object} model.CancionEntity
// @Router /cancion/{id} [get]
func FindById(c *gin.Context) {

	entity, err := repository.FindById(c.Param(id))
	if err != nil {
		log.Println("Se ha producido un error: ", err)
		c.JSON(http.StatusNoContent, gin.H{})
		return
	}
	c.JSON(http.StatusOK, &entity)
}

// FindAll godoc
// @Summary Busca todos los registros de la tabla
// @Produce json
// @Success 200 {array} model.CancionEntity
// @Router /cancion [get]
func FindAll(c *gin.Context) {
	dtos := repository.FindAll()
	c.JSON(http.StatusOK, dtos)
}

// Update godoc
// @Summary Actualiza un registro en Base de datos por id
// @Produce json
// @Param cancion body model.CancionUpdateDto true "Dto con los datos necesarios para actualizar por id"
// @Success 200
// @Router /cancion [patch]
func Update(c *gin.Context) {

	var dto model.CancionUpdateDto
	err := c.ShouldBindJSON(&dto)
	if err != nil {
		log.Println("Se ha producido un error: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = repository.Update(&dto)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

// Delete godoc
// @Summary Eliminar por Id en Base de datos
// @Produce json
// @Param id path string true "Id de la base de datos a eliminar"
// @Success 202,204
// @Router /cancion/{id} [delete]
func Delete(c *gin.Context) {

	err := repository.Delete(c.Param(id))
	if err != nil {
		c.JSON(http.StatusNoContent, gin.H{})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{})
}
