package repository

import (
	"veterinaria/backend/internal/models"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

type DuenoRepository struct {
	col *mgo.Collection
}

func NewDuenoRepo(db *mgo.Database) *DuenoRepository {
	return &DuenoRepository{col: db.C("Duenos")}
}

func (r *DuenoRepository) GetAll() ([]models.Dueno, error) {
	var duenos []models.Dueno
	err := r.col.Find(nil).All(&duenos)
	return duenos, err
}

func (r *DuenoRepository) GetByID(id int) (models.Dueno, error) {
	var dueno models.Dueno
	err := r.col.FindId(id).One(&dueno)
	return dueno, err
}

func (r *DuenoRepository) Create(dueno models.Dueno) error {
	return r.col.Insert(dueno)
}

func (r *DuenoRepository) Update(id int, dueno models.Dueno) error {
	return r.col.UpdateId(id, bson.M{"$set": dueno})
}

func (r *DuenoRepository) Delete(id int) error {
	return r.col.RemoveId(id)
}
