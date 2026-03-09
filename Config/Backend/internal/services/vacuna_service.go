package services

import (
	"veterinaria/backend/internal/models"
	"veterinaria/backend/internal/repository"
)

type VacunaService struct {
	repo *repository.VacunaRepository
}

func NewVacunaService(repo *repository.VacunaRepository) *VacunaService {
	return &VacunaService{repo: repo}
}

func (s *VacunaService) GetAll() ([]models.Vacuna, error) {
	return s.repo.GetAll()
}

func (s *VacunaService) GetByID(id int) (models.Vacuna, error) {
	return s.repo.GetByID(id)
}

func (s *VacunaService) Create(vacuna models.Vacuna) error {
	return s.repo.Create(vacuna)
}

func (s *VacunaService) Update(id int, vacuna models.Vacuna) error {
	return s.repo.Update(id, vacuna)
}

func (s *VacunaService) Delete(id int) error {
	return s.repo.Delete(id)
}
