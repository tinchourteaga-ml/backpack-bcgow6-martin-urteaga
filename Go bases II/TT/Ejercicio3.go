package main

import "fmt"

type tienda struct {
	Productos []producto
}

type producto struct {
	Tipo   string
	Nombre string
	Precio float64
}

type Producto interface {
	CalcularCosto() float64
}

type Ecommerce interface {
	Total() float64
	Agregar(producto)
}

func (p producto) CalcularCosto() float64 {
	switch p.Tipo {
	case "Mediano":
		return p.Precio * 0.03
	case "Grande":
		return p.Precio*0.06 + 2500
	}
	return 0
}

func (t tienda) Total() float64 {
	var total float64 = 0
	for _, prod := range t.Productos {
		total += prod.Precio + prod.CalcularCosto()
	}
	return total
}

func (t *tienda) Agregar(prod producto) {
	t.Productos = append(t.Productos, prod)
}

func nuevoProducto(tipo string, nombre string, precio float64) Producto {
	return producto{
		Tipo:   tipo,
		Nombre: nombre,
		Precio: precio,
	}
}

func nuevaTienda() Ecommerce {
	return &tienda{}
}

func main() {
	p1 := nuevoProducto("Mediano", "Notebook", 45000)
	p2 := nuevoProducto("Grande", "Heladera", 100000)
	p3 := nuevoProducto("Pequeño", "Mouse", 6000)
	t1 := nuevaTienda()

	t1.Agregar(p1.(producto)) // Type assertion
	t1.Agregar(p2.(producto))
	t1.Agregar(p3.(producto))

	fmt.Println("Valor total:$", t1.Total())
}
