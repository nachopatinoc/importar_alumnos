package main

import (
	"fmt"
	"log"
	"time"

	"importar_alumnos/config"
	"importar_alumnos/repository"
	"importar_alumnos/service"
)

func main() {
	tiempoInicio := time.Now()

	db, err := config.ConectarDB(".env")
	if err != nil {
		log.Fatalf("No se pudo conectar a la base de datos: %v", err)
	}
	defer db.Close()

	fmt.Println("Leyendo archivo CSV...")
	alumnos, err := service.ParsearCSV("data/alumnos.csv")
	if err != nil {
		log.Fatalf("Error al parsear CSV: %v", err)
	}
	fmt.Printf("Se leyeron %d alumnos\n", len(alumnos))

	fmt.Println("Insertando alumnos en la base...")
	err = repository.InsertarBatchAlumnos(db, alumnos)
	if err != nil {
		log.Fatalf("Error al insertar alumnos: %v", err)
	}

	tiempoFinal := time.Since(tiempoInicio)
	fmt.Printf("Todos los alumnos cargados en: %v\n", tiempoFinal)

	totalAlumnos, err := repository.ContarAlumnos(db)
	if err != nil {
		log.Fatalf("Error al contar alumnos: %v", err)
	}
	fmt.Printf("Total de alumnos cargados: %d\n", totalAlumnos)
}
