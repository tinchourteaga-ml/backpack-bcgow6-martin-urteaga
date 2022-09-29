package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	path, _ := os.Getwd()
	filePath := fmt.Sprint(path, "/myFile.csv")
	content, _ := os.ReadFile(filePath)
	prodLists := strings.Split(string(content), "\n")
	var sum float64
	fmt.Printf("ID\t Precio Cantidad\n")
	for _, prod := range prodLists[0:3] {
		p := strings.Split(prod, ";")
		price, _ := strconv.ParseFloat(p[1], 64)
		sum += price
		fmt.Printf("%s\t %s\t %s\n", p[0], p[1], p[2])
	}
	fmt.Printf("\t %.2f\t \n", sum)
}

// Solucion mas friendly
/* func main() {
	read, err := os.ReadFile("../ejercicio1/productos.csv")
	if err != nil {
		fmt.Println(err)
		return
	}

	//fmt.Println(string(read)) // formato byte
	data := string(read)
	fmt.Println(strings.ReplaceAll(data, ";", "\t\t\t"))
} */
