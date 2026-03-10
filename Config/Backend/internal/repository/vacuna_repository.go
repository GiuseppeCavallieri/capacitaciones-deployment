package repository

import (
	"database/sql"
	"veterinaria/backend/internal/models"
)

type VacunaRepository struct {
	db *sql.DB
}

func NewVacunaRepo(db *sql.DB) *VacunaRepository {
	return &VacunaRepository{db: db}
}

func (r *VacunaRepository) GetAll() ([]models.Vacuna, error) {
	rows, err := r.db.Query("SELECT id, fecha, nombrevacuna, id_perro FROM vacunas")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	vacunas := make([]models.Vacuna, 0)
	for rows.Next() {
		var v models.Vacuna
		if err := rows.Scan(&v.ID, &v.Fecha, &v.NombreVacuna, &v.IDPerro); err != nil {
			return nil, err
		}
		vacunas = append(vacunas, v)
	}
	return vacunas, rows.Err()
}

func (r *VacunaRepository) GetByID(id int) (models.Vacuna, error) {
	var v models.Vacuna
	err := r.db.QueryRow("SELECT id, fecha, nombrevacuna, id_perro FROM vacunas WHERE id=$1", id).Scan(&v.ID, &v.Fecha, &v.NombreVacuna, &v.IDPerro)
	return v, err
}

func (r *VacunaRepository) GetByPerroID(idPerro int) ([]models.Vacuna, error) {
	rows, err := r.db.Query("SELECT id, fecha, nombrevacuna, id_perro FROM vacunas WHERE id_perro=$1", idPerro)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	vacunas := make([]models.Vacuna, 0)
	for rows.Next() {
		var v models.Vacuna
		if err := rows.Scan(&v.ID, &v.Fecha, &v.NombreVacuna, &v.IDPerro); err != nil {
			return nil, err
		}
		vacunas = append(vacunas, v)
	}
	return vacunas, rows.Err()
}

func (r *VacunaRepository) Create(vacuna models.Vacuna) (models.Vacuna, error) {
	var id int
	err := r.db.QueryRow("INSERT INTO vacunas (fecha, nombrevacuna, id_perro) VALUES ($1, $2, $3) RETURNING id", vacuna.Fecha, vacuna.NombreVacuna, vacuna.IDPerro).Scan(&id)
	if err != nil {
		return vacuna, err
	}
	vacuna.ID = id
	return vacuna, nil
}

func (r *VacunaRepository) Update(id int, vacuna models.Vacuna) error {
	_, err := r.db.Exec("UPDATE vacunas SET fecha=$1, nombrevacuna=$2, id_perro=$3 WHERE id=$4", vacuna.Fecha, vacuna.NombreVacuna, vacuna.IDPerro, id)
	return err
}

func (r *VacunaRepository) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM vacunas WHERE id=$1", id)
	return err
}
