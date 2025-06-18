package test

import (
	"importar_alumnos/service"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCrearAlumno(t *testing.T) {
	alumnos, err := service.ParsearCSV("testdata/alumnos_test.csv")

	assert.NoError(t, err)
	assert.Equal(t, 5, len(alumnos))

	alumno := alumnos[0]

	assert.Equal(t, "VonRueden", alumno.Apellido)
	assert.Equal(t, "Emely Maya Dare", alumno.Nombre)
	assert.Equal(t, 1, alumno.NroLegajo)
	fechaNacimientoEsp := time.Date(1943, 11, 4, 0, 0, 0, 0, time.UTC)
	fechaIngresoEsp := time.Date(2014, 4, 16, 0, 0, 0, 0, time.UTC)
	assert.Equal(t, fechaNacimientoEsp, alumno.FechaNacimiento)
	assert.Equal(t, fechaIngresoEsp, alumno.FechaIngreso)
	assert.Equal(t, "LibretaEnrolamiento", alumno.TipoDocumento)
}
