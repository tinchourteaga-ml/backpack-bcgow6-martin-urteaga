package main

import (
	"fmt"
	"os"
)

type SalaryError struct {
	msg string
}

func (e SalaryError) Error() string {
	return e.msg
}

func checkSalary(salary int) (string, error) {
	if salary <= 150000 {
		return "", SalaryError{msg: "error: el salario ingresado no alcanza el mínimo imponible"}
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
