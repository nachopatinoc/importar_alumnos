package models

import "time"

type Alumno struct {
	Id              int
	Apellido        string
	Nombre          string
	NroDocumento    string
	TipoDocumento   string
	FechaNacimiento time.Time
	Sexo            string
	NroLegajo       int
	FechaIngreso    time.Time
}
