package Ejercicio1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRestar(t *testing.T) {
	n1 := 8
	n2 := 5
	resultadoEsperado := 3

	resultado := Restar(n1, n2)

	assert.Equal(t, resultadoEsperado, resultado, "deben ser iguales")
}
