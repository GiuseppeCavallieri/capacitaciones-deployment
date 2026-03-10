package repository

import (
	"database/sql"
	"veterinaria/backend/internal/models"
)

type PerroRepository struct {
	db *sql.DB
}

func NewPerroRepo(db *sql.DB) *PerroRepository {
	return &PerroRepository{db: db}
}

func (r *PerroRepository) GetAll() ([]models.Perro, error) {
	rows, err := r.db.Query("SELECT id, nombre, raza, color, edad, id_dueno FROM perros")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	perros := make([]models.Perro, 0)
	for rows.Next() {
		var p models.Perro
		if err := rows.Scan(&p.ID, &p.Nombre, &p.Raza, &p.Color, &p.Edad, &p.IDDueno); err != nil {
			return nil, err
		}
		perros = append(perros, p)
	}
	return perros, rows.Err()
}

func (r *PerroRepository) GetByID(id int) (models.Perro, error) {
	var p models.Perro
	err := r.db.QueryRow("SELECT id, nombre, raza, color, edad, id_dueno FROM perros WHERE id=$1", id).Scan(&p.ID, &p.Nombre, &p.Raza, &p.Color, &p.Edad, &p.IDDueno)
	return p, err
}

func (r *PerroRepository) Create(perro models.Perro) (models.Perro, error) {
	var id int
	err := r.db.QueryRow("INSERT INTO perros (nombre, raza, color, edad, id_dueno) VALUES ($1, $2, $3, $4, $5) RETURNING id", perro.Nombre, perro.Raza, perro.Color, perro.Edad, perro.IDDueno).Scan(&id)
	if err != nil {
		return perro, err
	}
	perro.ID = id
	return perro, nil
}

func (r *PerroRepository) Update(id int, perro models.Perro) error {
	_, err := r.db.Exec("UPDATE perros SET nombre=$1, raza=$2, color=$3, edad=$4, id_dueno=$5 WHERE id=$6", perro.Nombre, perro.Raza, perro.Color, perro.Edad, perro.IDDueno, id)
	return err
}

func (r *PerroRepository) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM perros WHERE id=$1", id)
	return err
}
