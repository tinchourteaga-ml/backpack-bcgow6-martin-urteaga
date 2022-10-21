package fibonacci

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFibonacci(t *testing.T) {
	var iterations int = 6
	var expectedResult int = 8

	result := fibonacci(iterations)

	assert.Equal(t, expectedResult, result, "deben ser iguales")
}
