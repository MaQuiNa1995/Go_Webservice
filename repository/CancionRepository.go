package repository

import (
	model "maquina1995/webservice/model"
)

func Create(entity *model.CancionEntity) {
	db := Connect()
	defer Close(db)

	db.Create(&entity)
}

func FindById() {
	db := Connect()
	defer Close(db)
}

func FindAll() []model.CancionEntity {

	db := Connect()
	defer Close(db)

	var canciones []model.CancionEntity
	db.Find(&canciones)

	return canciones
}

func Update() {
	db := Connect()
	defer Close(db)

}

func Delete() {
	db := Connect()
	defer Close(db)
}
