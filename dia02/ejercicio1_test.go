package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// El siguiente test comprueba que cuando se calcula un salario que está por debajo de los
// 50.000, no se cobra ningun impuesto
func TestCalcularImpuesto_salarioDebajoDe50000(t *testing.T) {
	// Arrange
	var montoEsperado float64 = 40000

	// Act
	montoObtenido, statusOK := calcularImpuesto(montoEsperado)

	// Assert
	assert.Equal(t, montoEsperado, montoObtenido)
	assert.True(t, statusOK)
}

// El siguiente test comprueba que cuando se caluka un salario que está entre 50000 y 150000,
// se cobra un impuesto del 17% del salario total
func TestCalcularImpuesto_salarioEntre50KY150K(t *testing.T) {
	// Arrange
	var salario float64 = 100000
	var montoEsperado float64 = salario * 17 / 100

	// Act
	montoObtenido, statusOK := calcularImpuesto(salario)

	// Assert
	assert.Equal(t, montoEsperado, montoObtenido)
	assert.True(t, statusOK)
}

// El siguiente test comprueba que cuando se caluka un salario que está entre 50000 y 150000,
// se cobra un impuesto del 17% del salario total
func TestCalcularImpuesto_salarioMayorA150K(t *testing.T) {
	// Arrange
	var salario float64 = 200000
	var montoEsperado float64 = salario * 27 / 100

	// Act
	montoObtenido, statusOK := calcularImpuesto(salario)

	// Assert
	assert.Equal(t, montoEsperado, montoObtenido)
	assert.True(t, statusOK)
}

// El siguiente test comprueba que cuando se intenta calcular el impuesto de un salario inválido,
// como un número negativo, se obtiene un false en el status y el montoObtenido es vacío o 0
func TestCalcularImpuesto_conNumeroInvalido(t *testing.T) {
	// Arrange
	var salario float64 = -342323

	// Act
	montoObtenido, statusOK := calcularImpuesto(salario)

	// Assert
	assert.Empty(t, montoObtenido)
	assert.False(t, statusOK)
}
