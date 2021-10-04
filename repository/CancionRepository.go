package repository

import (
	model "maquina1995/webservice/model"
)

func Create(entity *model.CancionEntity) {
	db := Connect()
	defer Close(db)

	db.Create(&entity)
}

func FindById(id string) (*model.CancionEntity, error) {
	db := Connect()
	defer Close(db)

	entity := model.CancionEntity{}

	if err := db.Where("id = ?", id).First(&entity).Error; err != nil {

		return &entity, err
	}

	return &entity, nil
}

func FindAll() []model.CancionEntity {

	db := Connect()
	defer Close(db)

	var entities []model.CancionEntity
	db.Find(&entities)

	return entities
}

func Update(dto *model.CancionUpdateDto) error {
	db := Connect()
	defer Close(db)

	var entity model.CancionUpdateDto
	if err := db.Where("id = ?", dto.Id).First(&entity).Error; err != nil {
		return err
	}

	db.Model(&entity).Updates(dto)
	return nil
}

func Delete(id string) error {
	db := Connect()
	defer Close(db)

	var entity model.CancionUpdateDto
	if err := db.Where("id = ?", id).First(&entity).Error; err != nil {
		return err
	}

	db.Delete(&entity)
	return nil
}
