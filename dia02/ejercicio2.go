package main

// func main() {
// 	result, err := promedioNotas(4, 5, 2, 10, 10, 9, 0)
// 	if err == nil {
// 		fmt.Println("El promedio de las notas es de:", result)
// 	} else {
// 		fmt.Println(err)
// 	}
// }

func promedioNotas(notas ...int) (promedio int, status bool) {
	if len(notas) == 0 {
		status = !status
		return
	}

	for _, nota := range notas {
		if nota <= 0 {
			promedio = 0
			return
		}
		promedio += nota
	}

	status = !status
	promedio = promedio / len(notas)
	return
}
