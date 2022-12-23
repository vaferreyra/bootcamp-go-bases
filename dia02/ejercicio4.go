package main

import (
	"errors"
	"fmt"
)

const (
	minimum = "minimum"
	maximum = "maximum"
	average = "average"
)

type Operacion func(...int) int

func main() {
	calcularEstadistica(average, 1, 2, 3, 4, 5)
	calcularEstadistica(minimum, 1, 2, 3, 4, 5)
	calcularEstadistica(maximum, 1, 2, 3, 4, 5)
	calcularEstadistica("asdadsasd", 1, 2)
	calcularEstadistica(minimum)
	calcularEstadistica(maximum)
	calcularEstadistica(average)
}

func calcularEstadistica(calculo string, notas ...int) {
	function, err := handlerCalculo(calculo)

	if err != nil {
		fmt.Println(err)
		return
	}

	r := function(notas...)
	fmt.Println("El resultado del calculo es:", r)
}

func handlerCalculo(calculo string) (Operacion, error) {
	switch calculo {
	case minimum:
		return opMinimum, nil
	case maximum:
		return opMaximum, nil
	case average:
		return opAverage, nil
	default:
		return nil, errors.New("El calculo es inválido")
	}
}

func opMaximum(numbers ...int) int {
	var max int

	for _, number := range numbers {
		if max < number {
			max = number
		}
	}
	return max
}

func opMinimum(numbers ...int) int {
	var min int

	for _, number := range numbers {
		if min > number {
			min = number
		}
	}
	return min
}

func opAverage(numbers ...int) int {
	if len(numbers) == 0 {
		return 0
	}

	var sumaTotal int

	for _, number := range numbers {
		sumaTotal += number
	}

	return sumaTotal / len(numbers)
}
