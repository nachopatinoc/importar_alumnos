package main

import (
	"fmt"
	"importar_alumnos/config"
	"importar_alumnos/models"
	"importar_alumnos/repository"
	"importar_alumnos/service"
	"log"
	"sync"
	"time"
)

const (
	csvPath   = "data/alumnos.csv"
	batchSize = 5000
)

func main() {

	db, err := config.ConectarDB(".env")
	if err != nil {
		log.Fatalf("No se pudo conectar a la base de datos: %v", err)
	}
	defer db.Close()

	alumnos, err := service.ParsearCSV(csvPath)
	if err != nil {
		log.Fatalf("Error al parsear CSV: %v", err)
	}

	if len(alumnos) == 0 {
		log.Fatal("El archivo CSV está vacío o no se pudo parsear correctamente.")
	}

	fmt.Printf("Se leyeron %d alumnos\n", len(alumnos))

	tiempoInicio := time.Now()

	var wg sync.WaitGroup
	maxConcurrent := 18
	semaforo := make(chan struct{}, maxConcurrent)

	for i := 0; i < len(alumnos); i += batchSize {
		fin := i + batchSize
		if fin > len(alumnos) {
			fin = len(alumnos)
		}
		batch := alumnos[i:fin]

		wg.Add(1)
		semaforo <- struct{}{}

		go func(i, fin int, b []models.Alumno) {
			defer wg.Done()
			defer func() { <-semaforo }()

			err := repository.InsertarBatchAlumnos(db, b)
			if err != nil {
				log.Fatalf("Error al insertar batch %d - %d: %v", i, fin, err)
			}
		}(i, fin, batch)
	}
	wg.Wait()
	tiempoFin := time.Since(tiempoInicio)
	fmt.Printf("Tiempo total de ejecución: %s\n", tiempoFin)

	alumnosTotales, err := repository.ContarAlumnos(db)
	if err != nil {
		log.Fatalf("Error al contar alumnos: %v", err)
	}
	fmt.Printf("Total de alumnos insertados: %d\n", alumnosTotales)
}
