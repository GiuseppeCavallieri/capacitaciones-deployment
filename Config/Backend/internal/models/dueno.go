package models

type Dueno struct {
	ID     int    `json:"_id"`
	Nombre string `json:"nombre"`
	Edad   int    `json:"edad"`
	Sexo   string `json:"sexo"`
}
