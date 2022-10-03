package main

import "fmt"

type Matrix struct {
	Valores    []float64
	Ancho      int
	Alto       int
	Max        float64
	Cuadratica bool
}

func (matrix *Matrix) setValues(valores ...float64) {
	max := valores[0]
	matrix.Valores = valores
	for _, valor := range valores {
		if valor > max {
			matrix.Max = valor
		}
	}
}

func (matrix *Matrix) setDimensions(ancho, alto int) {
	matrix.Ancho = ancho
	matrix.Alto = alto
	if ancho == alto {
		matrix.Cuadratica = true
	} else {
		matrix.Cuadratica = false
	}
}

func (matrix Matrix) print() {
	if (matrix.Ancho * matrix.Alto) > len(matrix.Valores) {
		zerosToAdd := matrix.Ancho*matrix.Alto - len(matrix.Valores)
		i := 0
		for i < zerosToAdd {
			matrix.Valores = append(matrix.Valores, 0)
			i++
		}
	}
	pos := 0
	for i := 0; i < matrix.Alto; i++ {
		j := 0
		for j < matrix.Ancho {
			fmt.Print(matrix.Valores[pos+j])
			j++
		}
		pos += j
		fmt.Print("\n")
	}
}

func main() {
	var matrix Matrix
	matrix.setValues(4, 5, 6, 1, 2, 3, 1, 4, 3, 6, 0, 4, 0, 1, 7)
	// matrix.setValues(4, 5, 6, 1, 2, 3, 1, 4, 3, 6) menos valores que la dimension, completo con ceros
	matrix.setDimensions(5, 3)
	matrix.print()
}
