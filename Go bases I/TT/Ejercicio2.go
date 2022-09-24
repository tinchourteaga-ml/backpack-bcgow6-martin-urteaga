package main

import "fmt"

func main() {
	const edadMinima, antiguedadMinima, sueldoMinimo = 22, 1, 100000
	var edad, antiguedad int = 23, 2
	var estaEmpleado bool = true
	var sueldo float64 = 150000

	if edad > edadMinima && antiguedad > antiguedadMinima && estaEmpleado {
		fmt.Println("Préstamo otorgado")
		if sueldo > sueldoMinimo {
			fmt.Println("No se cobrará interés")
		} else {
			fmt.Println("Se cobrará interés")
		}
	} else {
		fmt.Println("Préstamo no otorgado")
	}
}
