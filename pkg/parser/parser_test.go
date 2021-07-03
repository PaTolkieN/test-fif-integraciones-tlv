package parser_test

import (
	"test-fif-integraciones-tlv/pkg/parser"
	"testing"
)

const (
	value       string = "11AB123456789A050"
	valBadRegex string = "11AB12345678-A050"
	lenBadRegex string = "A1AB12345678-A050"
	TagBadRegex string = "A1AB12345678-R050"
)

func TestParseTLVValueFail(t *testing.T) {

	byteString := []byte(valBadRegex)
	_, err := parser.ParseTLV(byteString)
	if err == nil {
		t.Error("Test value fail OK")
	}
}

func TestParseTLVTagFail(t *testing.T) {

	byteString := []byte(TagBadRegex)
	_, err := parser.ParseTLV(byteString)
	if err == nil {
		t.Error("Test Tag Fail OK")
	}
}

func TestParseTLVLenghtFail(t *testing.T) {

	byteString := []byte(lenBadRegex)
	_, err := parser.ParseTLV(byteString)
	if err == nil {
		t.Error("Test length fail OK")
	}
}

func TestParseTLVOK(t *testing.T) {

	byteString := []byte(value)
	val, _ := parser.ParseTLV(byteString)
	if val != nil {
		t.Log("Test OK")
	}
}
