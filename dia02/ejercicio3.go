package main

// func main() {
// 	calcularSalario(2, "A")
// }

func calcularSalario(horasTrabajadas int, categoria string) (salario int, status bool) {
	if horasTrabajadas < 0 {
		return
	}

	status = !status

	switch categoria {
	case "A":
		salario = 3000 * horasTrabajadas
		salario += salario / 2
	case "B":
		salario = 1500 * horasTrabajadas
		salario += salario * 20 / 100
	case "C":
		salario = 1000 * horasTrabajadas
	default:
		status = !status
	}
	return
}
