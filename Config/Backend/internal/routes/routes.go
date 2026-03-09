package routes

import (
	"veterinaria/backend/internal/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, duenoCtrl *controllers.DuenoController, perroCtrl *controllers.PerroController, vacunaCtrl *controllers.VacunaController) {
	// Dueños
	r.GET("/duenos", duenoCtrl.GetAll)
	r.GET("/duenos/:id", duenoCtrl.GetByID)
	r.POST("/duenos", duenoCtrl.Create)
	r.PUT("/duenos/:id", duenoCtrl.Update)
	r.DELETE("/duenos/:id", duenoCtrl.Delete)

	// Perros
	r.GET("/perros", perroCtrl.GetAll)
	r.GET("/perros/:id", perroCtrl.GetByID)
	r.POST("/perros", perroCtrl.Create)
	r.PUT("/perros/:id", perroCtrl.Update)
	r.DELETE("/perros/:id", perroCtrl.Delete)
	r.GET("/perros/:id/dueno", perroCtrl.GetDueno)
	r.GET("/perros/:id/vacunas", perroCtrl.GetVacunas)

	// Vacunas
	r.GET("/vacunas", vacunaCtrl.GetAll)
	r.GET("/vacunas/:id", vacunaCtrl.GetByID)
	r.POST("/vacunas", vacunaCtrl.Create)
	r.PUT("/vacunas/:id", vacunaCtrl.Update)
	r.DELETE("/vacunas/:id", vacunaCtrl.Delete)
}
