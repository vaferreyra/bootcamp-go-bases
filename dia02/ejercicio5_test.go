package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Con TestDog compruebo que el cálculo para la comida de los perros funciona correctamente
func TestDog(t *testing.T) {
	// Arrange
	animalDog, err := animal(Dog)
	alimentoEsperado := float64(100)

	// Act
	alimentoObtenido := animalDog(10)

	// Assert
	assert.Equal(t, alimentoEsperado, alimentoObtenido)
	assert.Nil(t, err)
}

// Con TestCat compruebo que el cálculo para la comida de los gatos funciona correctamente
func TestCat(t *testing.T) {
	// Arrange
	animalCat, err := animal(Cat)
	alimentoEsperado := float64(50)

	// Act
	alimentoObtenido := animalCat(10)

	// Assert
	assert.Equal(t, alimentoEsperado, alimentoObtenido)
	assert.Nil(t, err)
}

// Con TestHamster compruebo que el cálculo para la comida de los hamsters funciona correctamente
func TestHamster(t *testing.T) {
	// Arrange
	animalHamster, err := animal(Hamster)
	alimentoEsperado := float64(2.5)

	// Act
	alimentoObtenido := animalHamster(10)

	// Assert
	assert.Equal(t, alimentoEsperado, alimentoObtenido)
	assert.Nil(t, err)
}

// Con TestSpider compruebo que el cálculo para la comida de las tarántulas funciona correctamente
func TestSpider(t *testing.T) {
	// Arrange
	animalSpider, err := animal(Spider)
	alimentoEsperado := float64(1.5)

	// Act
	alimentoObtenido := animalSpider(10)

	// Assert
	assert.Equal(t, alimentoEsperado, alimentoObtenido)
	assert.Nil(t, err)
}

// En este test comprueno que cuando se pasa un animal inválido como argumento,
// este levanta un error y no pasa ninguna función como respuesta
func TestAnimal_conAnimalInvalido(t *testing.T) {
	// Act
	animalInvalido, err := animal("Cocodrilo")

	// Assert
	assert.NotNil(t, err)
	assert.Nil(t, animalInvalido)
}
