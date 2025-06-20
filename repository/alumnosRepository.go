package repository

import (
	"database/sql"
	"fmt"
	"importar_alumnos/models"
	"strings"
)

func InsertarBatchAlumnos(db *sql.DB, alumnos []models.Alumno) error {
	if len(alumnos) == 0 {
		return nil
	}

	var queryBuilder strings.Builder
	queryBuilder.WriteString("INSERT INTO alumnos (")
	queryBuilder.WriteString("nro_legajo, apellido, nombre, nro_documento, tipo_documento, fecha_nacimiento, sexo, fecha_ingreso")
	queryBuilder.WriteString(") VALUES ")

	args := []interface{}{}
	for i, a := range alumnos {
		start := i*8 + 1
		queryBuilder.WriteString(fmt.Sprintf("($%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d),",
			start, start+1, start+2, start+3, start+4, start+5, start+6, start+7))

		args = append(args,
			a.NroLegajo,
			a.Apellido,
			a.Nombre,
			a.NroDocumento,
			a.TipoDocumento,
			a.FechaNacimiento,
			a.Sexo,
			a.FechaIngreso,
		)
	}

	query := strings.TrimSuffix(queryBuilder.String(), ",")

	_, err := db.Exec(query, args...)
	if err != nil {
		return fmt.Errorf("error insertando batch: %v", err)
	}

	return nil
}

func ContarAlumnos(db *sql.DB) (int, error) {
	var total int
	err := db.QueryRow("SELECT COUNT(*) FROM alumnos").Scan(&total)
	if err != nil {
		return 0, err
	}
	return total, nil
}
