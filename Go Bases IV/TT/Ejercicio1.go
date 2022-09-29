package main

import (
	"fmt"
	"os"
)

func leerArchivo(filePath string) {
	_, err := os.ReadFile(filePath)

	defer func() {
		panicMsg := recover()

		if panicMsg != nil {
			fmt.Println(panicMsg)
		}
	}()

	if err != nil {
		panic("error: el archivo indicado no fue encontrado o está dañado")
	}
}

func main() {
	leerArchivo("customers.txt")
	fmt.Println("ejecución finalizada")
}
