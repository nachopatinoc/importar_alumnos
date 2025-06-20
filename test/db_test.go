package test

import (
	"importar_alumnos/config"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConectarDB(t *testing.T) {
	db, err := config.ConectarDB("../.env")

	assert.NoError(t, err)
	assert.NotNil(t, db)

	err = db.Ping()
	assert.NoError(t, err)

	defer db.Close()
}
