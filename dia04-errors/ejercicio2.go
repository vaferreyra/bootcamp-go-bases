package main

import (
	"errors"
	"fmt"
)

var (
	ErrLowSalary = SalaryError{
		msg: "Error: el salario es menor a 10.000",
	}
	ErrHighSalary = SalaryError{
		msg: "Error: tu salario excede lo pedido",
	}
)

type SalaryError struct {
	msg string
}

func (e *SalaryError) Error() string {
	return e.msg
}

func consultarSalario(numero int) (result bool, err error) {
	if numero < 10000 {
		result = !result
		err = &ErrLowSalary
	}
	return
}

func main() {
	_, err := consultarSalario(5000)

	if err != nil {
		if errors.Is(err, &ErrLowSalary) {
			fmt.Println(err)
		}
	}
}
