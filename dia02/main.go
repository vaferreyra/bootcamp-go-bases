package main

import (
	"fmt"
)

func main() {
	fmt.Println(operaciones(4, 2))
}

func operaciones(numero1, numero2 float64) (suma float64, resta float64, multip float64, divis float64) {
	suma = numero1 + numero2
	resta = numero1 - numero2
	multip = numero1 * numero2

	if numero2 != 0 {
		divis = numero1 / numero2
	}
	return
}
