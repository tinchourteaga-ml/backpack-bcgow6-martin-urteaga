package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
)

type Cliente struct {
	Legajo          int
	NombreYApellido string
	Dni             string
	Telefono        string
	Domicilio       string
}

func generarId() (Cliente, error) {
	cliente := Cliente{}
	cliente.Legajo = rand.Int()

	if cliente == (Cliente{}) {
		return Cliente{}, errors.New("Fallo en la creación")
	}
	return cliente, nil
}

func validarDatos(cliente Cliente) (string, error) {
	if cliente == (Cliente{}) {
		return "", errors.New("Datos del cliente incompletos")
	}
	return "Cliente validado", nil
}

func verificarCliente(filePath string) {
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
	cliente, err := generarId()

	if err != nil {
		panic("error: no se pudo registrar al nuevo cliente")
	}

	verificarCliente("customers.txt")
	_, err = validarDatos(cliente)

	if err != nil {
		fmt.Println(err)
	}

	defer fmt.Println("No han quedado archivos abiertos")
	defer fmt.Println("Se detectaron varios errores en tiempo de ejecución")
	defer fmt.Println("Fin de la ejecución")
}
