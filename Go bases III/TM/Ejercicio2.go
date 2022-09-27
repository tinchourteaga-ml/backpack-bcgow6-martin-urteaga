package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	path, _ := os.Getwd()
	filePath := fmt.Sprint(path, "/myFile.csv")
	content, _ := os.ReadFile(filePath)
	prodLists := strings.Split(string(content), "\n")

	fmt.Printf("ID\t Precio Cantidad\n")
	for _, prod := range prodLists[0:3] {
		p := strings.Split(prod, ";")
		fmt.Printf("%s\t %s\t %s\n", p[0], p[1], p[2])
	}
}
