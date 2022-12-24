package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// El siguiente test comprueba que el promedio de las notas se realiza correctamente cuando
// todas las notas son válidas
func TestPromedioNotas(t *testing.T) {
	// Arrange
	var notas = []int{8, 9, 7}
	var promedioEsperado = (8 + 9 + 7) / 3

	// Act
	promedioObtenido, statusOK := promedioNotas(notas...)

	// Assert
	assert.Equal(t, promedioEsperado, promedioObtenido, "El promedio obtenido no es el esperado")
	assert.True(t, statusOK, "El status obtenido no es el esperado")
}

// El siguiente test comprueba que el promedio de notas no se calcula y el status retorna false
// cuando alguna de las notas a calculas es inválida
func TestPromedioNotas_conArgumentoInvalido(t *testing.T) {
	// Arrange
	var notas = []int{3, 4, -1}

	// Act
	promedioObtenido, statusOK := promedioNotas(notas...)

	// Assert
	assert.Empty(t, promedioObtenido)
	assert.False(t, statusOK)
}

// EL siguiente test comprueba que cuando se pide calcular promedio con un argumento sin
// valores, retorna 0 como promedio obtenido, y el status es true, o sea que no hay error (es subjetivo)
func TestPromedioNotas_conArgumentoVacio(t *testing.T) {
	// Arrange
	var notas = []int{}

	// Act
	promedioObtenido, statusOK := promedioNotas(notas...)

	// Assert
	assert.Empty(t, promedioObtenido)
	assert.True(t, statusOK)
}
