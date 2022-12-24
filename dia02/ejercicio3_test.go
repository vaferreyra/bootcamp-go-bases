package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// El siguiente test comprubea que cuando un empleado es categoria A, cobra 3000 por cada hora trabajada,
// y se le suma el 50% al salario total
func TestCalcularSalario_conCategoriaA(t *testing.T) {
	// Arrange
	horasTrabajadas := 5
	categoria := "A"
	salarioEsperado := (15000) + (7500)

	// Act
	salarioObtenido, statusOK := calcularSalario(horasTrabajadas, categoria)

	// Assert
	assert.Equal(t, salarioEsperado, salarioObtenido)
	assert.True(t, statusOK)
}

// El siguiente test comprubea que cuando un empleado es categoria B, cobra 1500 por cada hora trabajada,
// y se le suma el 20% al salario total
func TestCalcularSalario_conCategoriaB(t *testing.T) {
	// Arrange
	horasTrabajadas := 5
	categoria := "B"
	salarioEsperado := (7500) + (7500 * 20 / 100)

	// Act
	salarioObtenido, statusOK := calcularSalario(horasTrabajadas, categoria)

	// Assert
	assert.Equal(t, salarioEsperado, salarioObtenido)
	assert.True(t, statusOK)
}

// El siguiente test comprubea que cuando un empleado es categoria C, cobra 1000 por cada hora trabajada,
// y no se le suma ningun aumento
func TestCalcularSalario_conCategoriaC(t *testing.T) {
	// Arrange
	horasTrabajadas := 5
	categoria := "C"
	salarioEsperado := 5000

	// Act
	salarioObtenido, statusOK := calcularSalario(horasTrabajadas, categoria)

	// Assert
	assert.Equal(t, salarioEsperado, salarioObtenido)
	assert.True(t, statusOK)
}

// El siguiente test comprueba que cuando se intenta calcular el salario con una categoria
// invalida, este no se calcula y retorna false en el status
func TestCalcularSalario_conCategoriaInvalida(t *testing.T) {
	horasTrabajadas := 5
	categoria := "D"

	// Act
	salarioObtenido, statusOK := calcularSalario(horasTrabajadas, categoria)

	// Assert
	assert.Empty(t, salarioObtenido)
	assert.False(t, statusOK)
}

// El siguiente test comprueba que cuando se intenta calcular el salario con una cantidad de horas
// trabajadas invalida, como numero en negativo, este no se calcula y retorna false en el status
func TestCalcularSalario_conCantidadDeHorasInvalida(t *testing.T) {
	horasTrabajadas := -5
	categoria := "A"

	// Act
	salarioObtenido, statusOK := calcularSalario(horasTrabajadas, categoria)

	// Assert
	assert.Empty(t, salarioObtenido)
	assert.False(t, statusOK)
}
