package utils

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateLength(t *testing.T) {
	_, err := ValidateLength(0)
	assert.Equal(t, errors.New("Debe ingresar cadena a checkear"), err)
	valid, _ := ValidateLength(10)
	assert.Equal(t, true, valid)
}

func TestValidateValue(t *testing.T) {
	_, err := ValidateValue("AB12345...", 2)
	assert.Equal(t, errors.New("Valor no permite caracteres especiales"), err)
	valid, _ := ValidateValue("AB12345", 0)
	assert.Equal(t, true, valid)
}

func TestValidateType(t *testing.T) {

	_, err := ValidateType("F00")
	assert.Equal(t, errors.New("Valor inesperado para variable Tipo"), err)
	valid, _ := ValidateType("A00")
	assert.Equal(t, true, valid)
}
