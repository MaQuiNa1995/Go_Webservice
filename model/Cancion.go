package model

type CancionEntity struct {
	Id       uint    `json:"ID" gorm:"primary_key"`
	Nombre   string  `json:"NOMBRE"`
	Duracion float32 `json:"DURACION"`
	Genero   string  `json:"GENERO"`
}

type CancionCreate struct {
	Nombre   string  `json:"NOMBRE" binding:"required"`
	Duracion float32 `json:"DURACION" binding:"required"`
	Genero   string  `json:"GENERO" binding:"required"`
}

type CancionUpdate struct {
	Nombre   string  `json:"NOMBRE"`
	Duracion float32 `json:"DURACION"`
	Genero   string  `json:"GENERO"`
}
