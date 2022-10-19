package Ejercicio3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDividir(t *testing.T) {
	n1 := 6
	denominador := 0
	_, err := Dividir(n1, denominador)

	assert.NotNil(t, err)
}
