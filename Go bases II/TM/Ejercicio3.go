package main

import (
	"errors"
	"fmt"
)

func salarioPorHora(salario float64, minutos int) float64 {
	return salario * float64(minutos) / 60
}

func salarioExtra(salario float64, minutos int, porcentaje float64) float64 {
	return salarioPorHora(salario, minutos) * porcentaje
}

func calcularSalario(minutos int, categoria string) (float64, error) {
	switch categoria {
	case "A":
		return salarioPorHora(3000, minutos) + salarioExtra(3000, minutos, 0.5), nil
	case "B":
		return salarioPorHora(1500, minutos) + salarioExtra(1500, minutos, 0.2), nil
	case "C":
		return salarioPorHora(1000, minutos), nil
	}
	return 0, errors.New("Categoría inválida")
}

func main() {
	salarioA, err := calcularSalario(1000, "A")
	salarioB, err := calcularSalario(1000, "B")
	salarioC, err := calcularSalario(1000, "C")

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Salario A: $", salarioA)
	fmt.Println("Salario B: $", salarioB)
	fmt.Println("Salario C: $", salarioC)
}
