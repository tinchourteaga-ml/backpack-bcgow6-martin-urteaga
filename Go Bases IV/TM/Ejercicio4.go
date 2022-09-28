package main

import (
	"errors"
	"fmt"
	"os"
)

func salarioMensual(cantHoras int, valorHora float64) (float64, error) {
	if cantHoras < 80 {
		return 0, errors.New("error: el trabajador no puede haber trabajado menos de 80hs mensuales")
	}

	salario := float64(cantHoras) * valorHora

	if salario >= 150000 {
		salario -= salario * 0.1
	}
	return salario, nil
}

func calcAguinaldo(mejorSalario float64, cantMesesTrabajados int) (float64, error) {
	if mejorSalario < 0 {
		return 0, errors.New("error: el valor ingresado es negativo")
	} else if cantMesesTrabajados < 0 {
		return 0, errors.New("error: el valor de los meses es negativo")
	} else if cantMesesTrabajados > 6 {
		return 0, errors.New("error: el valor de los meses es mayor a un semestre (6)")
	}
	aguinaldo := mejorSalario / 12 * float64(cantMesesTrabajados)
	return aguinaldo, nil
}

func main() {
	var cantHoras int = 90
	var cantMesesTrabajados int = 5
	var valorHora float64 = 620
	var mejorSalario float64 = 120000

	salario, err := salarioMensual(cantHoras, valorHora)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	aguinaldo, err := calcAguinaldo(mejorSalario, cantMesesTrabajados)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("El salario mensual es: $%.2f y el aguinaldo es: $%.2f\n", salario, aguinaldo)
}
