package repository

import (
	"veterinaria/backend/internal/models"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

type VacunaRepository struct {
	col *mgo.Collection
}

func NewVacunaRepo(db *mgo.Database) *VacunaRepository {
	return &VacunaRepository{col: db.C("Vacunas")}
}

func (r *VacunaRepository) GetAll() ([]models.Vacuna, error) {
	vacunas := make([]models.Vacuna, 0)
	err := r.col.Find(nil).All(&vacunas)
	return vacunas, err
}

func (r *VacunaRepository) GetByID(id int) (models.Vacuna, error) {
	var vacuna models.Vacuna
	err := r.col.FindId(id).One(&vacuna)
	return vacuna, err
}

func (r *VacunaRepository) GetByPerroID(idPerro int) ([]models.Vacuna, error) {
	vacunas := make([]models.Vacuna, 0)
	err := r.col.Find(bson.M{"id_perro": idPerro}).All(&vacunas)
	return vacunas, err
}

func (r *VacunaRepository) Create(vacuna models.Vacuna) error {
	return r.col.Insert(vacuna)
}

func (r *VacunaRepository) Update(id int, vacuna models.Vacuna) error {
	return r.col.UpdateId(id, bson.M{"$set": vacuna})
}

func (r *VacunaRepository) Delete(id int) error {
	return r.col.RemoveId(id)
}
