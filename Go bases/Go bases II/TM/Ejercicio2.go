package main

import (
	"errors"
	"fmt"
)

func calcularPromedioCalificaciones(notas ...int) (float64, error) {
	var sumatoria int
	for _, nota := range notas {
		if nota < 0 {
			return 0, errors.New("No pueden existir notas negativas")
		}
		sumatoria += nota
	}
	return float64(sumatoria / len(notas)), nil
}

func main() {
	promedio, err := calcularPromedioCalificaciones(8, 6, 8, 9, 2, 10, 6, 6, 7)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Promedio:", promedio)
}
