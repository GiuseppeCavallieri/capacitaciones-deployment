package database

import (
	"log"
	"time"

	"veterinaria/backend/internal/config"

	"github.com/globalsign/mgo"
)


func ConnectDB(config *config.Config) (*mgo.Session, error) {
	session, err := mgo.DialWithTimeout(config.MongoURI, 10*time.Second)
	if err != nil {
		log.Fatalf("Error conectando a MongoDB: %v", err)
	}

	session.SetMode(mgo.Monotonic, true)
	log.Println("✅ Conexión a MongoDB exitosa")
	return session, nil
}