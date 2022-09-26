package main

import (
	"errors"
	"fmt"
)

const (
	minimum = "minimum"
	average = "average"
	maximum = "maximum"
)

func calcMin(valores ...float64) float64 {
	min := valores[0]
	for _, valor := range valores {
		if valor < min {
			min = valor
		}
	}
	return min
}

func calcMax(valores ...float64) float64 {
	max := valores[0]
	for _, valor := range valores {
		if valor > max {
			max = valor
		}
	}
	return max
}

func calcAvg(valores ...float64) float64 {
	var sumatoria float64
	for _, valor := range valores {
		sumatoria += valor
	}
	return sumatoria / float64(len(valores))
}

func operation(operacion string) (func(valores ...float64) float64, error) {
	switch operacion {
	case "minimum":
		return calcMin, nil
	case "maximum":
		return calcMax, nil
	case "average":
		return calcAvg, nil
	}
	return nil, errors.New("Operacion desconocida")
}

func main() {
	minFunc, err := operation(minimum)
	averageFunc, err := operation(average)
	maxFunc, err := operation(maximum)

	if err != nil {
		fmt.Println(err.Error())
	}

	minValue := minFunc(2, 3, 3, 4, 10, 2, 4, 5)
	averageValue := averageFunc(2, 3, 3, 4, 1, 2, 4, 5)
	maxValue := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)

	fmt.Println("Minimo:", minValue)
	fmt.Println("Promedio:", averageValue)
	fmt.Println("Maximo:", maxValue)
}
