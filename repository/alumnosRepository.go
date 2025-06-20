package repository

import (
	"database/sql"
	"fmt"
	"importar_alumnos/models"
)

func InsertarBatchAlumnos(db *sql.DB, alumnos []models.Alumno) error {
	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("error iniciando transacción: %v", err)
	}

	stmt, err := tx.Prepare(`
		INSERT INTO alumnos 
		(nro_legajo, apellido, nombre, nro_documento, tipo_documento, fecha_nacimiento, sexo, fecha_ingreso)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error insertando batch: %v", err)
	}
	defer stmt.Close()

	for _, a := range alumnos {
		_, err := stmt.Exec(
			a.NroLegajo, a.Apellido, a.Nombre,
			a.NroDocumento, a.TipoDocumento,
			a.FechaNacimiento, a.Sexo, a.FechaIngreso,
		)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("error insertando batch: %v", err)
		}
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("error al confirmar transacción: %v", err)
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
