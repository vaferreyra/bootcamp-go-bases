package main

import (
	"fmt"
)

func consultarSalarioErrorf(numero int) (result string, err error) {
	if numero < 10000 {
		err = fmt.Errorf("Error: el minimo imponible es de 10.000, y tu salario es de %d", numero)
		return
	}
	result = fmt.Sprintf("Tu salario %d excede de los 10.000", numero)
	return
}

func main4() {
	result, err := consultarSalarioErrorf(15000)

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(result)
}
