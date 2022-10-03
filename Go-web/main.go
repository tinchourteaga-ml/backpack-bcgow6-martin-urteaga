package main

import (
	"encoding/json"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

type ProductsCatalog struct {
	Products []Product
}

type Product struct {
	Id           int
	Name         string
	Color        string
	Price        float64
	Stock        int
	Code         string
	Published    bool
	CreationDate string
}

func greetingsHandler(ctx *gin.Context) {
	name := "Martin"
	ctx.JSON(200, gin.H{
		"message": "Hola " + name,
	})
}

func getAllHandler(ctx *gin.Context) {
	file, _ := ioutil.ReadFile("products.json")
	catalog := ProductsCatalog{}
	json.Unmarshal([]byte(file), &catalog)
	ctx.JSON(200, catalog)
}

func main() {
	router := gin.Default()
	router.GET("/greetings", greetingsHandler)
	router.GET("/products", getAllHandler)
	router.Run()
}
