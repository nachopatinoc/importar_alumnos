package service

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"importar_alumnos/models"
	"importar_alumnos/validators"
	"os"
	"strconv"
	"time"
)

func ParsearCSV(path string) ([]models.Alumno, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("error abriendo archivo: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(bufio.NewReader(file))
	var alumnos []models.Alumno

	_, err = reader.Read()
	if err != nil {
		return nil, fmt.Errorf("error leyendo encabezados: %v", err)
	}

	for {
		line, err := reader.Read()
		if err != nil {
			break
		}

		nroLegajo, err := strconv.Atoi(line[6])
		if err != nil {
			return nil, fmt.Errorf("legajo inválido: %v", err)
		}
		nroLegajo, err = validators.ValidarLegajo(nroLegajo)
		if err != nil {
			return nil, err
		}

		fechaNacimiento, err := time.Parse("2006-01-02", line[4])
		if err != nil {
			return nil, fmt.Errorf("fecha de nacimiento inválida: %v", err)
		}
		fechaNacimiento, err = validators.ValidarFecha(fechaNacimiento, "fecha_nacimiento")
		if err != nil {
			return nil, err
		}

		fechaIngreso, err := time.Parse("2006-01-02", line[7])
		if err != nil {
			return nil, fmt.Errorf("fecha de ingreso inválida: %v", err)
		}
		fechaIngreso, err = validators.ValidarFecha(fechaIngreso, "fecha_ingreso")
		if err != nil {
			return nil, err
		}

		sexo, err := validators.ValidarSexo(line[5])
		if err != nil {
			return nil, err
		}

		tipoDocumento, err := validators.ValidarTipoDocumento(line[3])
		if err != nil {
			return nil, err
		}

		alumno := models.Alumno{
			Apellido:        line[0],
			Nombre:          line[1],
			NroDocumento:    line[2],
			TipoDocumento:   tipoDocumento,
			FechaNacimiento: fechaNacimiento,
			Sexo:            sexo,
			NroLegajo:       nroLegajo,
			FechaIngreso:    fechaIngreso,
		}

		alumnos = append(alumnos, alumno)
	}

	return alumnos, nil
}
