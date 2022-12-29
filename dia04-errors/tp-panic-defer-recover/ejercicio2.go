package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.ReadFile("customers.txt")

	defer func() {
		fmt.Println("Ejecución finalizada")
	}()

	if err != nil {
		panic("El archivo indicado no fue encontrado o está dañado")
	}

	fmt.Printf("%s\n", file)
}
