package main

import "fmt"

func main() {
	var empleados = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}
	var mayoresDe21 int

	fmt.Println("Edad de Benjamin:", empleados["Benjamin"])

	for _, elemento := range empleados {
		if elemento > 21 {
			mayoresDe21++
		}
	}

	fmt.Println("Empleados mayores de 21 años:", mayoresDe21)

	empleados["Federico"] = 25
	delete(empleados, "Pedro")
}
