package models

type Perro struct {
	ID      int    `bson:"_id" json:"_id"`
	Nombre  string `bson:"nombre" json:"nombre"`
	Raza    string `bson:"raza" json:"raza"`
	Color   string `bson:"color" json:"color"`
	Edad    int    `bson:"edad" json:"edad"`
	IDDueno int    `bson:"id_dueno" json:"id_dueno"`
}
