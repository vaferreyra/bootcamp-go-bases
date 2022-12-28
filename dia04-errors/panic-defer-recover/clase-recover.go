package main

import "fmt"

func main2() {
	fmt.Println("Comienza el programa")
	function()
	fmt.Println("Finalizó el programa")
}

func function() {
	var apiSlice []string
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("Ocurrió un error")
		}
	}()

	_ = apiSlice[0]

	// Ocurre un panic, pero al tener un recover, se "disfraza"
	// y ejecuta lo que escribí dentro de su función, para tener un panic
	// un poco más elegante en consola.
}
