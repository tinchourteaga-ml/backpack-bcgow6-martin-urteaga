package main

import "fmt"

func main() {
	var palabra string = "bootcamp"

	fmt.Println("Cantidad de letras: ", len(palabra))

	for i := 0; i < len(palabra); i++ {
		fmt.Printf("%c\n", palabra[i])
	}
}
