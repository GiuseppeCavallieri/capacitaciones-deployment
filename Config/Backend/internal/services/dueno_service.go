package services

import (
	"veterinaria/backend/internal/models"
	"veterinaria/backend/internal/repository"
)

type DuenoService struct {
	repo *repository.DuenoRepository
}

func NewDuenoService(repo *repository.DuenoRepository) *DuenoService {
	return &DuenoService{repo: repo}
}

func (s *DuenoService) GetAll() ([]models.Dueno, error) {
	return s.repo.GetAll()
}

func (s *DuenoService) GetByID(id int) (models.Dueno, error) {
	return s.repo.GetByID(id)
}

func (s *DuenoService) Create(dueno models.Dueno) (models.Dueno, error) {
	return s.repo.Create(dueno)
}

func (s *DuenoService) Update(id int, dueno models.Dueno) error {
	return s.repo.Update(id, dueno)
}

func (s *DuenoService) Delete(id int) error {
	return s.repo.Delete(id)
}
