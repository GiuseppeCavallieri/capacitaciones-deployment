package repository

import (
	"veterinaria/backend/internal/models"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

type PerroRepository struct {
	col *mgo.Collection
}

func NewPerroRepo(db *mgo.Database) *PerroRepository {
	return &PerroRepository{col: db.C("Perros")}
}

func (r *PerroRepository) GetAll() ([]models.Perro, error) {
	perros := make([]models.Perro, 0)
	err := r.col.Find(nil).All(&perros)
	return perros, err
}

func (r *PerroRepository) GetByID(id int) (models.Perro, error) {
	var perro models.Perro
	err := r.col.FindId(id).One(&perro)
	return perro, err
}

func (r *PerroRepository) Create(perro models.Perro) error {
	return r.col.Insert(perro)
}

func (r *PerroRepository) GetMaxID() (int, error) {
	var perro models.Perro
	err := r.col.Find(nil).Sort("-_id").One(&perro)
	if err != nil {
		return 0, err
	}
	return perro.ID, nil
}

func (r *PerroRepository) Update(id int, perro models.Perro) error {
	return r.col.UpdateId(id, bson.M{"$set": perro})
}

func (r *PerroRepository) Delete(id int) error {
	return r.col.RemoveId(id)
}
