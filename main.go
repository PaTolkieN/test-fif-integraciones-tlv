package main

import (
	"bufio"
	"fmt"
	"os"
	"test-fif-integraciones-tlv/pkg/parser"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Ingrese datos de entrada: ")

	strTlv, _ := reader.ReadString('\n')

	fmt.Println(strTlv)
	byteString := []byte(strTlv)
	result, err := parser.ParseTLV(byteString)
	if err != nil {
		fmt.Println("Ha ocurrido un error : ", err)
	} else {
		for key, value := range result {
			fmt.Printf("%s Valor es %v\n", key, value)
		}
	}
}

//11AB123456789A050200N2310AB87654321A3011AB123456789A07
