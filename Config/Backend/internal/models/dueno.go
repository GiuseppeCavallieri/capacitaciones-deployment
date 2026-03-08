package models

type Dueno struct {
	ID     int    `bson:"_id" json:"_id"`
	Nombre string `bson:"nombre" json:"nombre"`
	Edad   int    `bson:"edad" json:"edad"`
	Sexo   string `bson:"sexo" json:"sexo"`
}
