package constants

import (
	"errors"
	"regexp"
	"strings"
)

const (
	SizeTLV         int    = 2
	SumSizeTag      int    = 5
	InitialPosition int    = 2
	regex           string = "^[a-zA-Z0-9]*$"
	prefixAlpha     string = "A"
	prefixNumeric   string = "N"
)

func ValidateLength(largo int) (bool, error) {
	if largo <= 0 {
		return false, errors.New("Debe ingresar cadena a checkear")
	}
	return true, nil
}

func ValidateValue(valor string, largo int) (bool, error) {
	re := regexp.MustCompile(regex)
	regex := re.MatchString(valor)
	if regex == false {
		return false, errors.New("Valor no permite caracteres especiales")
	}
	if len(strings.TrimSpace(valor)) < largo {
		return false, errors.New("Valor no puede ser de largo menor al definido en el campo Largo")
	}
	return true, nil
}

func ValidateType(tipo string) (bool, error) {

	if strings.HasPrefix(tipo, prefixAlpha) == false && strings.HasPrefix(tipo, prefixNumeric) == false {
		err1 := errors.New("Valor inesperado para variable Tipo")
		return false, err1
	}
	return true, nil
}
