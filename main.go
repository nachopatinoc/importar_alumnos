package main

import (
	"fmt"
	"time"
)

func main() {
	tiempoInicio := time.Now()

	// Acá va tu lógica principal (importar CSV, insertar en DB, etc.)
	// ejemplo:
	// err := service.ImportarAlumnos("alumnos.csv", db)

	// Simulamos que el programa hace algo lento
	time.Sleep(3 * time.Second)

	tiempoFin := time.Since(tiempoInicio)
	fmt.Println("El programa tardó:", tiempoFin)
}
