package model

type CancionEntity struct {
	Id       uint    `json:"id" gorm:"primary_key"`
	Nombre   string  `json:"nombre"`
	Duracion float32 `json:"duracion"`
	Genero   string  `json:"genero"`
}

type CancionCreateDto struct {
	Nombre   string  `json:"nombre" binding:"required"`
	Duracion float32 `json:"duracion" binding:"required"`
	Genero   string  `json:"genero" binding:"required"`
}

type CancionUpdateDto struct {
	Id       string  `json:"id"`
	Nombre   string  `json:"nombre"`
	Duracion float32 `json:"duracion"`
	Genero   string  `json:"genero"`
}
