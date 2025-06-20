package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func ConectarDB(path string) (*sql.DB, error) {
	err := godotenv.Load(path)
	if err != nil {
		log.Println("Advertencia: no se pudo cargar .env, se usarán variables del entorno")
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	name := os.Getenv("DB_NAME")

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, pass, name)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, fmt.Errorf("error abriendo conexión a DB: %v", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("no se pudo conectar a la base de datos: %v", err)
	}

	return db, nil
}
