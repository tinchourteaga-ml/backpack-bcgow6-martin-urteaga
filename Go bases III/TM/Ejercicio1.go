package main

import (
	"fmt"
	"os"
)

type Producto struct {
	Id       int
	Precio   float64
	Cantidad int
}

func main() {
	path, _ := os.Getwd()
	filePath := fmt.Sprint(path, "/myFile.csv")

	p1 := Producto{Id: 111, Precio: 1500, Cantidad: 2}
	p2 := Producto{Id: 222, Precio: 3000, Cantidad: 5}
	p3 := Producto{Id: 333, Precio: 10000, Cantidad: 1}

	listaCompras := []Producto{p1, p2, p3}

	file, _ := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

	for _, prod := range listaCompras {
		csvLine := fmt.Sprint(prod.Id, ";", prod.Precio, ";", prod.Cantidad, "\n")
		file.Write([]byte(csvLine))
	}

	file.Close()
}
