package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// En este test se comprueba que cuando se pide calcular la nota minima de una lista de notas,
// la funcion devuelve la nota más baja
func TestCalcularEstadistica_notaMinima(t *testing.T) {
	// Arrange
	var notas = []int{5, 2, 10, 5, 9, 9, 7}
	resultadoEsperado := 2

	// Act
	resultadoObtenido, statusOK := calcularEstadistica(minimum, notas...)

	// Assert
	assert.Equal(t, resultadoEsperado, resultadoObtenido)
	assert.True(t, statusOK)
}

// En este test se comprueba que cuando se pide calcular la nota máxima de una lista de notas,
// la funcion devuelve la nota más alta
func TestCalcularEstadistica_notaMaxima(t *testing.T) {
	// Arrange
	var notas = []int{5, 2, 10, 5, 9, 9, 7}
	resultadoEsperado := 10

	// Act
	resultadoObtenido, statusOK := calcularEstadistica(maximum, notas...)

	// Assert
	assert.Equal(t, resultadoEsperado, resultadoObtenido)
	assert.True(t, statusOK)
}

// En este test se comprueba que cuando se pide calcular el promedio de una lista de notas,
// la funcion devuelve el promedio total de todas las notas indicadas
func TestCalcularEstadistica_promedioNotas(t *testing.T) {
	// Arrange
	var notas = []int{5, 2, 10, 5, 9}
	resultadoEsperado := (5 + 2 + 10 + 5 + 9) / 5

	// Act
	resultadoObtenido, statusOK := calcularEstadistica(average, notas...)

	// Assert
	assert.Equal(t, resultadoEsperado, resultadoObtenido)
	assert.True(t, statusOK)
}

// Con este test demuestro que cuando se pasa una estadistica inválida como argumento,
// devuelve un status false y el resultadoObtenido es 0
func TestCalcularEstadistica_conEstadisticaInvalida(t *testing.T) {
	// Arrange
	var notas = []int{5, 2, 10, 5, 9}
	estadisticaInvalida := "Hola"

	// Act
	resultadoObtenido, statusOK := calcularEstadistica(estadisticaInvalida, notas...)

	// Assert
	assert.Empty(t, resultadoObtenido)
	assert.False(t, statusOK)
}

// Con este test demuestro que cuando se pasa una lista de notas vacia, devuelve un error
func TestCalcularEstadistica_sinNingunaNotaComoArgumento(t *testing.T) {
	// Arrange
	var notas = []int{}

	// Act
	resultadoObtenido, statusOK := calcularEstadistica(minimum, notas...)

	// Assert
	assert.Empty(t, resultadoObtenido)
	assert.False(t, statusOK)
}

// Con este test demuestro que cuando se pasa una lista de notas con algún número inválido, devuelve un error
func TestCalcularEstadistica_conNotasInvalidas(t *testing.T) {
	// Arrange
	var notas = []int{3, 4, 5, -10, 0, -2}

	// Act
	resultadoObtenido, statusOK := calcularEstadistica(minimum, notas...)

	// Assert
	assert.Empty(t, resultadoObtenido)
	assert.False(t, statusOK)
}
