package repository

import (
	"database/sql"
	"veterinaria/backend/internal/models"
)

type DuenoRepository struct {
	db *sql.DB
}

func NewDuenoRepo(db *sql.DB) *DuenoRepository {
	return &DuenoRepository{db: db}
}

func (r *DuenoRepository) GetAll() ([]models.Dueno, error) {
	rows, err := r.db.Query("SELECT id, nombre, edad, sexo FROM duenos")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	duenos := make([]models.Dueno, 0)
	for rows.Next() {
		var d models.Dueno
		if err := rows.Scan(&d.ID, &d.Nombre, &d.Edad, &d.Sexo); err != nil {
			return nil, err
		}
		duenos = append(duenos, d)
	}
	return duenos, rows.Err()
}

func (r *DuenoRepository) GetByID(id int) (models.Dueno, error) {
	var d models.Dueno
	err := r.db.QueryRow("SELECT id, nombre, edad, sexo FROM duenos WHERE id=$1", id).Scan(&d.ID, &d.Nombre, &d.Edad, &d.Sexo)
	return d, err
}

func (r *DuenoRepository) Create(dueno models.Dueno) (models.Dueno, error) {
	var id int
	err := r.db.QueryRow("INSERT INTO duenos (nombre, edad, sexo) VALUES ($1, $2, $3) RETURNING id", dueno.Nombre, dueno.Edad, dueno.Sexo).Scan(&id)
	if err != nil {
		return dueno, err
	}
	dueno.ID = id
	return dueno, nil
}

func (r *DuenoRepository) Update(id int, dueno models.Dueno) error {
	_, err := r.db.Exec("UPDATE duenos SET nombre=$1, edad=$2, sexo=$3 WHERE id=$4", dueno.Nombre, dueno.Edad, dueno.Sexo, id)
	return err
}

func (r *DuenoRepository) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM duenos WHERE id=$1", id)
	return err
}
