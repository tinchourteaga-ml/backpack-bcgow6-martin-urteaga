package main

import (
	"fmt"
	"math"
)

type Producto struct {
	Nombre   string
	Precio   float64
	Cantidad int
}

type Servicio struct {
	Nombre  string
	Precio  float64
	Minutos int
}

type Mantenimiento struct {
	Nombre string
	Precio float64
}

func sumarProductos(productos []Producto, canal chan float64) {
	var sum float64
	for _, prod := range productos {
		sum += prod.Precio * float64(prod.Cantidad)
	}
	canal <- sum
	close(canal)
}

func sumarServicios(servicios []Servicio, canal chan float64) {
	var sum float64
	for _, serv := range servicios {
		sum += serv.Precio * math.Ceil(float64(serv.Minutos)/30)
	}
	canal <- sum
	close(canal)
}

func sumarMantenimiento(mantenimientos []Mantenimiento, canal chan float64) {
	var sum float64
	for _, mant := range mantenimientos {
		sum += mant.Precio
	}
	canal <- sum
	close(canal)
}

func main() {
	c1 := make(chan float64)
	c2 := make(chan float64)
	c3 := make(chan float64)

	p1 := Producto{Nombre: "Heladera", Precio: 1000, Cantidad: 1}
	p2 := Producto{Nombre: "Mouse", Precio: 3000, Cantidad: 5}
	p3 := Producto{Nombre: "HDD", Precio: 10000, Cantidad: 1}
	s1 := Servicio{Nombre: "Reparacion", Precio: 8000, Minutos: 60}
	s2 := Servicio{Nombre: "Cambio pantalla", Precio: 12000, Minutos: 125}
	m1 := Mantenimiento{Nombre: "Limpieza teclado", Precio: 3000}

	productos := []Producto{p1, p2, p3}
	servicios := []Servicio{s1, s2}
	mantenimientos := []Mantenimiento{m1}

	go sumarProductos(productos, c1)
	go sumarServicios(servicios, c2)
	go sumarMantenimiento(mantenimientos, c3)

	t1 := <-c1
	t2 := <-c2
	t3 := <-c3

	fmt.Println("Monto total: $", t1+t2+t3)
}
