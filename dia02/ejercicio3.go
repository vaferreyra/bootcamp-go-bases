package main

import (
	"fmt"
)

func main() {
	calcularSalario(2, "A")

}

func calcularSalario(horasTrabajadas int, categoria string) {
	var calculo int
	switch categoria {
	case "A":
		calculo = 3000 * horasTrabajadas
		fmt.Printf("Tu salario mensual es de: %v, pero te corresponde un aumento del 50, salario total de: %v \n", calculo, calculo+(calculo/2))
	case "B":
		calculo = 1500 * horasTrabajadas
		fmt.Printf("Tu salario mensual es de: %v, pero te corresponde un aumento del 20, salario total de: %v \n", calculo, calculo+(calculo*20/100))
	case "C":
		fmt.Printf("Tu salario mensual es de: %v \n", 1000*horasTrabajadas)
	default:
		fmt.Println("Categoria inválida")
	}
}
