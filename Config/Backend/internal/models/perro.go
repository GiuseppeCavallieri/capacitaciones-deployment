package models

type Perro struct {
	ID      int    `json:"_id"`
	Nombre  string `json:"nombre"`
	Raza    string `json:"raza"`
	Color   string `json:"color"`
	Edad    int    `json:"edad"`
	IDDueno int    `json:"id_dueno"`
}
