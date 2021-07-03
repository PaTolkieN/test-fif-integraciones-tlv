package parser

import (
	"fmt"
	"strconv"
	"strings"
	"test-fif-integraciones-tlv/pkg/model"
	constants "test-fif-integraciones-tlv/pkg/utils"
)

func ParseTLV(tlvByte []byte) (map[string]string, error) {

	tlvStr := string(tlvByte)
	lenStr := len(strings.TrimSpace(tlvStr))
	_, err := constants.ValidateLength(lenStr)
	if err != nil {
		return nil, err
	}
	fmt.Println("Arreglo de bytes desde parámetros: ", tlvByte)
	fmt.Println("Largo del string: ", lenStr)
	cont := 2
	m := make(map[string]string)
	var tlvStruct model.Tlv
	for cont < lenStr {
		fmt.Println("Valor de contador ==> ", cont)
		if cont == constants.InitialPosition {
			leng, err := strconv.Atoi(tlvStr[:constants.SizeTLV])
			fmt.Println("Largo TLV => ", leng)
			if err != nil {
				return nil, err
			}
			val := tlvStr[constants.InitialPosition : constants.SizeTLV+leng]
			fmt.Println("valor TLV => ", val)
			valid, err := constants.ValidateValue(val, leng)
			fmt.Println("valor es válido? ", valid)
			if err != nil {
				return nil, err
			}
			tag := tlvStr[constants.SizeTLV+leng : leng+constants.SumSizeTag]
			fmt.Println("Tipo TLV => ", tag)
			valid, err = constants.ValidateType(tag)
			fmt.Println("Tipo es válido? ", valid)
			if err != nil {
				return nil, err
			}
			cont = constants.SumSizeTag + leng
			tlvStruct = model.Tlv{
				Largo: leng,
				Tipo:  tag,
				Valor: val,
			}
			m[tlvStruct.Tipo] = tlvStruct.Valor
			fmt.Println("Valor de contador ==> ", cont)
			if lenStr-cont < 2 {
				cont = lenStr
			}
		} else {
			len2, err := strconv.Atoi(tlvStr[cont : cont+constants.SizeTLV])
			fmt.Println("Largo TLV => ", len2)
			if err != nil {
				return nil, err
			}
			val2 := tlvStr[cont+constants.SizeTLV : cont+constants.SizeTLV+len2]
			fmt.Println("valor TLV => ", val2)
			valid, err := constants.ValidateValue(val2, len2)
			fmt.Println("valor es válido? ", valid)
			if err != nil {
				return nil, err
			}
			tag2 := tlvStr[cont+constants.SizeTLV+len2 : cont+len2+constants.SumSizeTag]
			valid, err = constants.ValidateType(tag2)
			fmt.Println("Tipo es válido? ", valid)
			if err != nil {
				return nil, err
			}
			cont = cont + len2 + constants.SumSizeTag
			tlvStruct = model.Tlv{
				Largo: len2,
				Tipo:  tag2,
				Valor: val2,
			}
			m[tlvStruct.Tipo] = tlvStruct.Valor
		}
	}
	return m, nil
}
