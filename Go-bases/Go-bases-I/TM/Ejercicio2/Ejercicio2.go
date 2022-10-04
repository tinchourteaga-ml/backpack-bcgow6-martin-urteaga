package main

import "fmt"

func main() {
	var temperatura, humedad, presion float64
	temperatura = 17
	humedad = 45
	presion = 1022

	fmt.Print("Temperatura: ", temperatura, "º", "\nHumedad: ", humedad, "%", "\nPresion: ", presion, "hpa\n")
}

// Asignaria el tipo de dato Float dado que son valores que no necesariamente van a ser enteros.
