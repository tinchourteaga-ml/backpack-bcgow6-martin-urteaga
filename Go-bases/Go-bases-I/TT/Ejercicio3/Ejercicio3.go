package main

import "fmt"

func main() {
	var numMes int = 11

	var meses = []string{"Enero", "Febrero", "Marzo", "Abril", "Mayo", "Junio", "Julio",
		"Agosto", "Septiembre", "Octubre", "Noviembre", "Diciembre"}

	fmt.Println("Mes:", meses[numMes-1])
}

// La desventaja de esta forma de hacerlo es que el array con los meses tiene que estar ordenado, ademas deberia contemplar
// la posibilidad de recibir un numero que no corresponda a algun mes y catchearlo.
// Otra opción (mejor) es realizarlo con un switch, de esta forma no tengo el problema mencionado anteriormente: no necesito que
// este ordenado y puedo recurrir al "default" para manejar los inputs invalidos.
