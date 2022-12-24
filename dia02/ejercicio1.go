package main

// func main() {
// 	calcularImpuesto()
// }

func calcularImpuesto(salario float64) (montoTotal float64, ok bool) {

	if salario < 0 {
		return
	}

	ok = !ok

	if salario > 150000 {
		montoTotal = obtenerPorcentaje(salario, 27)
	} else if salario > 50000 {
		montoTotal = obtenerPorcentaje(salario, 17)
	} else {
		montoTotal = salario
	}
	return
}

func obtenerPorcentaje(salario, porcentaje float64) float64 {
	return salario * porcentaje / 100
}
