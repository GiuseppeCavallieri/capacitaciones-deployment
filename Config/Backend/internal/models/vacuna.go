package models

type Vacuna struct {
	ID           int    `json:"_id"`
	Fecha        string `json:"fecha"`
	NombreVacuna string `json:"nombrevacuna"`
	IDPerro      int    `json:"id_perro"`
}
