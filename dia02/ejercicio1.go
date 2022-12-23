package main

import "fmt"

func main() {
	calcularImpuesto()
}

func calcularImpuesto() {
	var salario float64
	fmt.Scan(&salario)

	if salario < 0 {
		fmt.Println("El monto ingresado no es válido")
		return
	}

	if salario > 150000 {
		fmt.Println("Tu impuesto es del 27%, y el total de tu impuesto es de:", obtenerPorcentaje(salario, 27))
	} else if salario > 50000 {
		fmt.Println("Tu impuesto es del 17%, y el total de tu impuesto es de:", obtenerPorcentaje(salario, 17))
	} else {
		fmt.Println("No te corresponden impuestos")
	}
}

func obtenerPorcentaje(salario, porcentaje float64) float64 {
	return salario * porcentaje / 100
}
