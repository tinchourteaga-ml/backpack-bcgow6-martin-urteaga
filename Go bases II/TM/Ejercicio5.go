package main

import (
	"errors"
	"fmt"
)

const (
	dog       = "dog"
	cat       = "cat"
	hamster   = "hamster"
	tarantula = "tarantula"

	// En gramos
	comidaPerro     = 10000
	comidaGato      = 5000
	comidaHamster   = 250
	comidaTarantula = 150
)

func animalDog(cantAnimal float64) float64 {
	return cantAnimal * comidaPerro
}

func animalCat(cantAnimal float64) float64 {
	return cantAnimal * comidaGato
}

func animalHamster(cantAnimal float64) float64 {
	return cantAnimal * comidaHamster
}

func animalTarantula(cantAnimal float64) float64 {
	return cantAnimal * comidaTarantula
}

func Animal(tipoAnimal string) (func(cantAnimal float64) float64, error) {
	switch tipoAnimal {
	case "dog":
		return animalDog, nil
	case "cat":
		return animalCat, nil
	case "hamster":
		return animalHamster, nil
	case "tarantula":
		return animalTarantula, nil
	}
	return nil, errors.New("No existe el animal")
}

func main() {
	animalDog, msg := Animal(dog)
	animalCat, msg := Animal(cat)
	animalHamster, msg := Animal(hamster)
	animalTarantula, msg := Animal(tarantula)

	if msg != nil {
		fmt.Println(msg.Error())
	}

	var amount float64
	amount += animalDog(5)
	amount += animalCat(8)
	amount += animalHamster(10)
	amount += animalTarantula(3)

	fmt.Println("Cantidad de comida a comprar:", amount, "gramos")
}
