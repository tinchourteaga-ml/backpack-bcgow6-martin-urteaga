package main

import "fmt"

type Alumno struct {
	Nombre   string
	Apellido string
	Dni      string
	Fecha    string
}

func (alum Alumno) detalle() {
	fmt.Printf("Nombre: %s\nApellido: %s\nDni: %s\nFecha: %s\n", alum.Nombre, alum.Apellido, alum.Dni, alum.Fecha)
}

func main() {
	a1 := Alumno{
		Nombre:   "John",
		Apellido: "Doe",
		Dni:      "11111111",
		Fecha:    "24/11/2020",
	}
	a1.detalle()
}
