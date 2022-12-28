package main

import "fmt"

func main() {
	salario, err := calcularSalarioMensual(10, 1000)

	if err != nil {
		panic(err)
	}

	fmt.Println("Tu salario total es de", salario)
}

func calcularSalarioMensual(precioHora int, cantidadHoras int) (salario int, err error) {
	if cantidadHoras < 80 {
		err = fmt.Errorf("El mínimo de horas mensuales trabajadas debe ser mayor a 80, el de usted fue %d", cantidadHoras)
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
