package main

import (
	"errors"
	"fmt"
)

var (
	ErrCantidadDeHorasNegativa = fmt.Errorf("El número de cantidad de horas trabajadas no puede ser negativo")
	ErrCantidadDeHorasBaja     = fmt.Errorf("El mínimo de horas mensuales trabajadas debe ser mayor a 80")
)

func main() {
	salario, err := calcularSalarioMensual(10, -2)

	if err != nil {
		if errors.Is(err, ErrCantidadDeHorasNegativa) {
			panic("Número de horas negativo")
		}
		if errors.Is(err, ErrCantidadDeHorasBaja) {
			panic("Cantidad de horas menor a 80")
		}
	}

	fmt.Println("Tu salario total es de", salario)
}

func calcularSalarioMensual(precioHora int, cantidadHoras int) (salario int, err error) {
	if cantidadHoras < 0 {
		err = ErrCantidadDeHorasNegativa
		return
	}

	if cantidadHoras < 80 {
		err = ErrCantidadDeHorasBaja
		return
	}

	salario = precioHora * cantidadHoras

	if salario > 150000 {
		salario = descontarConceptoDeImpuesto(salario)
	}
	return
}

func descontarConceptoDeImpuesto(salario int) int {
	return salario * 90 / 100
}
