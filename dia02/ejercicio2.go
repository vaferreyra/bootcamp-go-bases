package main

import (
	"errors"
)

// func main() {
// 	result, err := promedioNotas(4, 5, 2, 10, 10, 9, 0)
// 	if err == nil {
// 		fmt.Println("El promedio de las notas es de:", result)
// 	} else {
// 		fmt.Println(err)
// 	}
// }

func promedioNotas(notas ...int) (int, error) {
	var totalNota int
	for _, nota := range notas {
		if nota <= 0 {
			return 0, errors.New("Hay notas inválidas")
		}
		totalNota += nota
	}
	return totalNota / len(notas), nil
}
