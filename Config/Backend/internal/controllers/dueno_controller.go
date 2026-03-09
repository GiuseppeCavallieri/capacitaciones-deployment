package controllers

import (
	"net/http"
	"strconv"
	"veterinaria/backend/internal/models"
	"veterinaria/backend/internal/services"

	"github.com/gin-gonic/gin"
)

type DuenoController struct {
	svc *services.DuenoService
}

func NewDuenoController(svc *services.DuenoService) *DuenoController {
	return &DuenoController{svc: svc}
}

func (c *DuenoController) GetAll(ctx *gin.Context) {
	duenos, err := c.svc.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, duenos)
}

func (c *DuenoController) GetByID(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	dueno, err := c.svc.GetByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Dueño no encontrado"})
		return
	}
	ctx.JSON(http.StatusOK, dueno)
}

func (c *DuenoController) Create(ctx *gin.Context) {
	var dueno models.Dueno
	if err := ctx.ShouldBindJSON(&dueno); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.svc.Create(dueno); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, dueno)
}

func (c *DuenoController) Update(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	var dueno models.Dueno
	if err := ctx.ShouldBindJSON(&dueno); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.svc.Update(id, dueno); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, dueno)
}

func (c *DuenoController) Delete(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	if err := c.svc.Delete(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Dueño eliminado"})
}
