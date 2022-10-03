package main

import (
	"fmt"
	"os"
)

func checkSalary(salary int) (string, error) {
	if salary <= 150000 {
		return "", fmt.Errorf("error: el mínimo imponible es de $150000 y el salario ingresado es de: $%d", salary)
	}
	return "Debe pagar impuesto", nil
}

func main() {
	salary := 80000
	str, err := checkSalary(salary)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(str)
}
