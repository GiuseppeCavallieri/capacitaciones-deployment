package controllers

import (
	"net/http"
	"strconv"
	"veterinaria/backend/internal/models"
	"veterinaria/backend/internal/services"

	"github.com/gin-gonic/gin"
)

type VacunaController struct {
	svc *services.VacunaService
}

func NewVacunaController(svc *services.VacunaService) *VacunaController {
	return &VacunaController{svc: svc}
}

func (c *VacunaController) GetAll(ctx *gin.Context) {
	vacunas, err := c.svc.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, vacunas)
}

func (c *VacunaController) GetByID(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	vacuna, err := c.svc.GetByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Vacuna no encontrada"})
		return
	}
	ctx.JSON(http.StatusOK, vacuna)
}

func (c *VacunaController) Create(ctx *gin.Context) {
	var vacuna models.Vacuna
	if err := ctx.ShouldBindJSON(&vacuna); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.svc.Create(vacuna); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, vacuna)
}

func (c *VacunaController) Update(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	var vacuna models.Vacuna
	if err := ctx.ShouldBindJSON(&vacuna); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.svc.Update(id, vacuna); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, vacuna)
}

func (c *VacunaController) Delete(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	if err := c.svc.Delete(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Vacuna eliminada"})
}
