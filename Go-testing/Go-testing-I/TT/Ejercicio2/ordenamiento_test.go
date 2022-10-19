package Ejercicio2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrdenarSliceAsc(t *testing.T) {
	valoresOrdenados := OrdenarSliceAsc([]int{5, 3, 4, 7, 10, 9})
	valoresOrdenadosEsperados := []int{3, 4, 5, 7, 9, 10}

	assert.Equal(t, valoresOrdenadosEsperados, valoresOrdenados, "deben ser iguales")
}
