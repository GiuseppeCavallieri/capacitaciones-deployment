package models

type Vacuna struct {
	ID           int    `bson:"_id" json:"_id"`
	Fecha        string `bson:"fecha" json:"fecha"`
	NombreVacuna string `bson:"nombrevacuna" json:"nombrevacuna"`
	IDPerro      int    `bson:"id_perro" json:"id_perro"`
}
