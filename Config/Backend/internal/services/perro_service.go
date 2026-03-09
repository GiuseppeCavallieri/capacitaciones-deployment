package services

import (
	"errors"
	"veterinaria/backend/internal/models"
	"veterinaria/backend/internal/repository"
)

var (
	ErrPerroNoEncontrado  = errors.New("Perro no encontrado")
	ErrDuenoNoEncontrado  = errors.New("Dueño no encontrado")
	ErrVacunaNoEncontrada = errors.New("Vacunas no encontradas")
)

type PerroService struct {
	perroRepo  *repository.PerroRepository
	duenoRepo  *repository.DuenoRepository
	vacunaRepo *repository.VacunaRepository
}

func NewPerroService(pr *repository.PerroRepository, dr *repository.DuenoRepository, vr *repository.VacunaRepository) *PerroService {
	return &PerroService{perroRepo: pr, duenoRepo: dr, vacunaRepo: vr}
}

// GetDuenoDePerro obtiene el dueño de un perro
func (s *PerroService) GetDuenoDePerro(idPerro int) (models.Dueno, error) {
	perro, err := s.perroRepo.GetByID(idPerro)
	if err != nil {
		return models.Dueno{}, ErrPerroNoEncontrado
	}
	dueno, err := s.duenoRepo.GetByID(perro.IDDueno)
	if err != nil {
		return models.Dueno{}, ErrDuenoNoEncontrado
	}
	return dueno, nil
}

// GetVacunasDePerro obtiene las vacunas de un perro
func (s *PerroService) GetVacunasDePerro(idPerro int) ([]models.Vacuna, error) {
	_, err := s.perroRepo.GetByID(idPerro)
	if err != nil {
		return nil, ErrPerroNoEncontrado
	}
	return s.vacunaRepo.GetByPerroID(idPerro)
}

func (s *PerroService) GetAll() ([]models.Perro, error) {
	return s.perroRepo.GetAll()
}

func (s *PerroService) GetByID(id int) (models.Perro, error) {
	return s.perroRepo.GetByID(id)
}

func (s *PerroService) Create(perro models.Perro) (models.Perro, error) {
	maxID, err := s.perroRepo.GetMaxID()
	if err != nil {
		maxID = 0
	}
	perro.ID = maxID + 1
	return perro, s.perroRepo.Create(perro)
}

func (s *PerroService) Update(id int, perro models.Perro) error {
	return s.perroRepo.Update(id, perro)
}

func (s *PerroService) Delete(id int) error {
	return s.perroRepo.Delete(id)
}
