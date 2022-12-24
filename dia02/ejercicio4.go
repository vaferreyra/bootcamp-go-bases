package main

import (
	"errors"
)

const (
	minimum = "minimum"
	maximum = "maximum"
	average = "average"
)

type Operacion func(...int) int

// func main() {
// 	calcularEstadistica(average, 1, 2, 3, 4, 5)
// a, b := calcularEstadistica(minimum, 1, 2, 3, 4, 5, -10)
// fmt.Println(a, b)
// calcularEstadistica(maximum, 1, 2, 3, 4, 5)
// calcularEstadistica("asdadsasd", 1, 2)
// calcularEstadistica(minimum)
// calcularEstadistica(maximum)
// calcularEstadistica(average)
// }

func calcularEstadistica(calculo string, notas ...int) (nota int, status bool) {
	function, err := handlerCalculo(calculo)
	statusNotas := comprobarNotas(notas...)

	if err != nil || len(notas) == 0 || !statusNotas {
		return
	}

	nota = function(notas...)
	status = !status
	return
}

func handlerCalculo(calculo string) (operacion Operacion, err error) {
	switch calculo {
	case minimum:
		operacion = opMinimum
	case maximum:
		operacion = opMaximum
	case average:
		operacion = opAverage
	default:
		err = errors.New("El calculo es inválido")
	}
	return
}

func opMaximum(numbers ...int) (max int) {
	if len(numbers) == 0 {
		return
	}

	max = numbers[0]

	for _, number := range numbers {
		if max < number {
			max = number
		}
	}

	return
}

func opMinimum(numbers ...int) (min int) {
	if len(numbers) == 0 {
		return
	}

	min = numbers[0]
	for _, number := range numbers {
		if min > number {
			min = number
		}
	}

	return
}

func opAverage(numbers ...int) (avg int) {
	if len(numbers) == 0 {
		return
	}

	for _, number := range numbers {
		avg += number
	}

	avg = avg / len(numbers)
	return
}

func comprobarNotas(numbers ...int) (status bool) {
	for _, numero := range numbers {
		if numero <= 0 {
			return
		}
	}
	status = !status
	return
}
