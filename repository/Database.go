package repository

import (
	"log"
	model "maquina1995/webservice/model"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func Connect() *gorm.DB {
	db, err := gorm.Open("mysql", "root:pass@tcp(localhost:3306)/go_application")

	if err != nil {
		log.Println("fallo al conectar con la base de datos mas info: ", err)
		panic("fallo al conectar con la base de datos")
	}

	db.AutoMigrate(&model.CancionEntity{})

	return db
}

func Close(db *gorm.DB) {

	err := db.Close()

	if err != nil {
		log.Println("fallo al cerrar la conexion con la base de datos mas info: ", err)
		panic("fallo al desconectar con la base de datos")
	}
}
