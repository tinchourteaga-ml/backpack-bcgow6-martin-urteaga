package main

import "fmt"

type Producto struct {
	Nombre   string
	Precio   float64
	Cantidad int
}

type Usuario struct {
	Nombre    string
	Apellido  string
	Correo    string
	Productos []Producto
}

func nuevoProducto(nombre string, precio float64) Producto {
	return Producto{
		Nombre: nombre,
		Precio: precio,
	}
}

func agregarProducto(usuario *Usuario, producto *Producto, cantidad int) {
	producto.Cantidad = cantidad
	usuario.Productos = append(usuario.Productos, *producto)
}

func borrarProductos(usuario *Usuario) {
	usuario.Productos = []Producto{}
}

func main() {
	u1 := Usuario{Nombre: "John", Apellido: "Doe", Correo: "jdoe@gmail.com"}
	p1 := nuevoProducto("Heladera", 100000)
	p2 := nuevoProducto("Notebook", 45000)
	agregarProducto(&u1, &p1, 1)
	agregarProducto(&u1, &p2, 2)
	fmt.Println(u1)
	borrarProductos(&u1)
	fmt.Println(u1)
}
