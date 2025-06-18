package validators

import (
	"fmt"
	"time"
)

func ValidarLegajo(nroLegajo int) (int, error) {
	if nroLegajo <= 0 {
		return 0, fmt.Errorf("legajo inválido, debe ser mayor a 0")
	}
	return nroLegajo, nil
}

func ValidarFecha(fecha time.Time, campo string) (time.Time, error) {
	if fecha.IsZero() {
		return time.Time{}, fmt.Errorf("fecha %s inválida: fecha vacía", campo)
	}

	if fecha.After(time.Now()) {
		return time.Time{}, fmt.Errorf("fecha %s inválida: no puede ser una fecha futura", campo)
	}

	return fecha, nil
}

func ValidarSexo(sexo string) (string, error) {
	if sexo != "M" && sexo != "F" {
		return "", fmt.Errorf("sexo inválido: %v", sexo)
	}
	return sexo, nil
}

func ValidarTipoDocumento(tipoDocumento string) (string, error) {
	if tipoDocumento != "DNI" && tipoDocumento != "LibretaEnrolamiento" && tipoDocumento != "Pasaporte" && tipoDocumento != "LibretaCivica" {
		return "", fmt.Errorf("tipo de documento inválido: %v", tipoDocumento)
	}
	return tipoDocumento, nil
}
