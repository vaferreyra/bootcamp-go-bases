package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumar(t *testing.T) {
	num1 := 5
	num2 := 3

	resultadoEsperado := 8

	resultado := num1 + num2

	assert.Equal(t, resultadoEsperado, resultado, "Hubo un error en el test sumar")
}

func TestSumarFalla(t *testing.T) {
	num1 := 5
	num2 := 3

	resultadoEsperado := 8

	resultado := num1 - num2

	assert.NotEqual(t, resultadoEsperado, resultado, "Hubo un error en el test sumar")
}
