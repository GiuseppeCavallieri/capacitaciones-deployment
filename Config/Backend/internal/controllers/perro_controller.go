package controllers

import (
	"net/http"
	"strconv"
	"veterinaria/backend/internal/models"
	"veterinaria/backend/internal/services"

	"github.com/gin-gonic/gin"
)

type PerroController struct {
	svc *services.PerroService
}

func NewPerroController(svc *services.PerroService) *PerroController {
	return &PerroController{svc: svc}
}

func (c *PerroController) GetAll(ctx *gin.Context) {
	perros, err := c.svc.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, perros)
}

func (c *PerroController) GetByID(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	perro, err := c.svc.GetByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Perro no encontrado"})
		return
	}
	ctx.JSON(http.StatusOK, perro)
}

func (c *PerroController) Create(ctx *gin.Context) {
	var perro models.Perro
	if err := ctx.ShouldBindJSON(&perro); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	created, err := c.svc.Create(perro)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, created)
}

func (c *PerroController) Update(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	var perro models.Perro
	if err := ctx.ShouldBindJSON(&perro); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.svc.Update(id, perro); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, perro)
}

func (c *PerroController) Delete(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	if err := c.svc.Delete(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Perro eliminado"})
}

// GetDueno obtiene el dueño de un perro
func (c *PerroController) GetDueno(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	dueno, err := c.svc.GetDuenoDePerro(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, dueno)
}

// GetVacunas obtiene las vacunas de un perro
func (c *PerroController) GetVacunas(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	vacunas, err := c.svc.GetVacunasDePerro(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, vacunas)
}