package test

import (
	"importar_alumnos/config"
	"importar_alumnos/models"
	"importar_alumnos/repository"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestInsertarBatchAlumnos_DB(t *testing.T) {
	db, err := config.ConectarDB("../.env")
	assert.NoError(t, err)
	defer db.Close()

	alumnos := []models.Alumno{
		{
			Apellido:        "Ignacio",
			Nombre:          "Patiño",
			NroDocumento:    "46397680",
			TipoDocumento:   "DNI",
			FechaNacimiento: time.Date(1990, 3, 3, 0, 0, 0, 0, time.UTC),
			Sexo:            "M",
			NroLegajo:       10175,
			FechaIngreso:    time.Now(),
		},
		{
			Apellido:        "Fernández",
			Nombre:          "Luis",
			NroDocumento:    "20202020",
			TipoDocumento:   "DNI",
			FechaNacimiento: time.Date(1988, 4, 4, 0, 0, 0, 0, time.UTC),
			Sexo:            "M",
			NroLegajo:       9900,
			FechaIngreso:    time.Now(),
		},
	}

	err = repository.InsertarBatchAlumnos(db, alumnos)

	assert.NoError(t, err, "no debería haber error al insertar alumnos únicos")

	for _, a := range alumnos {
		_, _ = db.Exec("DELETE FROM alumnos WHERE nro_legajo = $1", a.NroLegajo)
	}
}
