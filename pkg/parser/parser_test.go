package parser

import (
	//"test-fif-integraciones-tlv/pkg/parser"
	"errors"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	value        string = "11AB123456789A5005123V5N06"
	valBadRegex  string = "11AB12345678-A50"
	valBadRegex2 string = "11AB123456789A5005123V.N06"
	lenBadRegex  string = "A1AB123456789A50"
	lenBadRegex2 string = "11AB123456789A50V5123V5N06"
	TagBadRegex  string = "11AB123456789R50"
	TagBadRegex2 string = "11AB123456789A5005123V5R06"
	emptyValue   string = ""
)

func TestMain(m *testing.M) {
	exitVal := m.Run()

	os.Exit(exitVal)
}
func TestParseTLVEmptyValueFail(t *testing.T) {
	byteString := []byte(emptyValue)
	_, err := ParseTLV(byteString)
	assert.NotNil(t, err)
	assert.Equal(t, errors.New("Hubo un problema al convertir el largo de String a Entero"), err)
}

func TestParseTLVValueFail(t *testing.T) {
	byteString := []byte(valBadRegex)
	_, err := ParseTLV(byteString)
	assert.NotNil(t, err)
	assert.Equal(t, errors.New("Valor no permite caracteres especiales"), err)
}
func TestParseTLVTagFail(t *testing.T) {

	byteString := []byte(TagBadRegex)
	_, err := ParseTLV(byteString)
	assert.NotNil(t, err)
}

func TestParseTLVLenghtFail(t *testing.T) {

	byteString := []byte(lenBadRegex)
	_, err := ParseTLV(byteString)
	assert.NotNil(t, err)
	assert.Equal(t, errors.New("Hubo un problema al convertir el string a entero"), err)
}

func TestParseTLVValueFail2(t *testing.T) {
	byteString := []byte(valBadRegex2)
	_, err := ParseTLV(byteString)
	assert.NotNil(t, err)
	assert.Equal(t, errors.New("Valor no permite caracteres especiales"), err)
}
func TestParseTLVTagFail2(t *testing.T) {

	byteString := []byte(TagBadRegex2)
	_, err := ParseTLV(byteString)
	assert.NotNil(t, err)
}

func TestParseTLVLenghtFail2(t *testing.T) {

	byteString := []byte(lenBadRegex2)
	_, err := ParseTLV(byteString)
	assert.NotNil(t, err)
}
func TestParseTLVOK(t *testing.T) {

	expected := map[string]string{
		"A50": "AB123456789",
		"N06": "123V5",
	}
	byteString := []byte(value)
	val, _ := ParseTLV(byteString)
	assert.NotNil(t, val)
	assert.Equal(t, expected, val)
}
