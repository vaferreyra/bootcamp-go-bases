package main

import (
	"errors"
	"fmt"
)

var (
	ErrLowSalaryWithErrors  = errors.New("Error: el salario es menor a 10.000")
	ErrHighSalaryWithErrors = errors.New("Error: tu salario excede lo pedido")
)

func consultarSalarioConErrors(numero int) (result bool, err error) {
	if numero < 10000 {
		result = !result
		err = ErrLowSalaryWithErrors
	}
	return
}

func main3() {
	_, err := consultarSalarioConErrors(5000)

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(errors.Is(err, ErrLowSalaryWithErrors))
	fmt.Println(errors.Is(err, ErrHighSalaryWithErrors))
}
