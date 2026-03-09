package main

import (
	"log"

	"veterinaria/backend/internal/config"
	"veterinaria/backend/internal/controllers"
	"veterinaria/backend/internal/database"
	"veterinaria/backend/internal/middleware"
	"veterinaria/backend/internal/repository"
	"veterinaria/backend/internal/routes"
	"veterinaria/backend/internal/services"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig()

	session, err := database.ConnectDB(cfg)
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}
	defer session.Close()

	db := session.DB(cfg.DBName)

	// Repositories
	duenoRepo := repository.NewDuenoRepo(db)
	perroRepo := repository.NewPerroRepo(db)
	vacunaRepo := repository.NewVacunaRepo(db)

	// Services
	duenoSvc := services.NewDuenoService(duenoRepo)
	perroSvc := services.NewPerroService(perroRepo, duenoRepo, vacunaRepo)
	vacunaSvc := services.NewVacunaService(vacunaRepo)

	// Controllers
	duenoCtrl := controllers.NewDuenoController(duenoSvc)
	perroCtrl := controllers.NewPerroController(perroSvc)
	vacunaCtrl := controllers.NewVacunaController(vacunaSvc)

	// Router
	router := gin.Default()
	router.Use(middleware.CORS())
	routes.SetupRoutes(router, duenoCtrl, perroCtrl, vacunaCtrl)

	log.Println("Servidor corriendo en", cfg.Port)
	router.Run(cfg.Port)
}