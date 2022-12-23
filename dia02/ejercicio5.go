package main

import (
	"errors"
)

const (
	Dog     = "dog"
	Cat     = "cat"
	Hamster = "hamster"
	Spider  = "tarantula"
)

// func main() {
// 	// Cuantos kg de alimento necesito para alimentar 10 perros y 10 gatos?
// 	var amount float64

// 	animalDog, err1 := animal(Dog)
// 	animalCat, err2 := animal(Cat)

// 	if err1 != nil || err2 != nil {
// 		fmt.Println("Animal inválido")
// 		return
// 	}

// 	amount += animalDog(10)
// 	amount += animalCat(10)

// 	fmt.Printf("Se necesitan %v kg de alimento para alimentar los animales \n", amount)
// }

func animal(animal string) (funcion func(int) float64, err error) {
	switch animal {
	case Dog:
		funcion = animalDog
	case Cat:
		funcion = animalCat
	case Hamster:
		funcion = animalHamster
	case Spider:
		funcion = animalSpider
	default:
		err = errors.New("Animal inválido")
	}
	return
}

func animalDog(cantidadAnimales int) float64 {
	return float64(cantidadAnimales) * cantidadDeAlimento(Dog)
}

func animalCat(cantidadAnimales int) float64 {
	return float64(cantidadAnimales) * cantidadDeAlimento(Cat)
}

func animalHamster(cantidadAnimales int) float64 {
	return float64(cantidadAnimales) * cantidadDeAlimento(Hamster)
}

func animalSpider(cantidadAnimales int) float64 {
	return float64(cantidadAnimales) * cantidadDeAlimento(Spider)
}

func cantidadDeAlimento(animal string) (pesoAlimento float64) {
	switch animal {
	case Dog:
		pesoAlimento = 10
	case Cat:
		pesoAlimento = 5
	case Hamster:
		pesoAlimento = 0.25
	case Spider:
		pesoAlimento = 0.150
	}
	return
}
