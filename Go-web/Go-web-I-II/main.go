package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

var AUTH_TOKEN = "123ABC"
var catalog = ProductsCatalog{}
var products []ProductDTO

type ProductsCatalog struct {
	Products []Product
}

type Product struct {
	Id           int
	Name         string
	Color        string
	Price        string
	Stock        string
	Code         string
	Published    string
	CreationDate string
}

type ProductDTO struct {
	Id           int    `json:"id" binding:"required"`
	Name         string `json:"name" binding:"required"`
	Color        string `json:"color" binding:"required"`
	Price        string `json:"price" binding:"required"`
	Stock        string `json:"stock" binding:"required"`
	Code         string `json:"code" binding:"required"`
	Published    string `json:"published" binding:"required"`
	CreationDate string `json:"creationDate" binding:"required"`
}

// Hidratamos el catalogo con los productos el archivo json
func readJSON() error {
	file, err := ioutil.ReadFile("products.json")

	if err != nil {
		return err
	}

	json.Unmarshal([]byte(file), &catalog)
	return nil
}

func greetingsHandler(ctx *gin.Context) {
	name := "Martin"
	ctx.JSON(200, gin.H{
		"message": "Hola " + name,
	})
}

func getAllHandler(ctx *gin.Context) {
	filteredCatalog := ProductsCatalog{}
	id, _ := strconv.Atoi(ctx.Query("id"))
	name := ctx.Query("name")
	color := ctx.Query("color")
	price := ctx.Query("price")
	stock := ctx.Query("stock")
	code := ctx.Query("code")
	published := ctx.Query("published")
	creationDate := ctx.Query("creationDate")

	for _, prod := range catalog.Products {
		// Si quisieramos que no nos tire el valor default, podemos hacer que el tipo sea puntero, entonces devuelve nil
		if prod.Id == id || prod.Name == name || prod.Color == color || prod.Price == price || prod.Stock == stock || prod.Code == code || prod.Published == published || prod.CreationDate == creationDate {
			filteredCatalog.Products = append(filteredCatalog.Products, prod)
		}
	}

	if len(filteredCatalog.Products) > 0 {
		ctx.JSON(200, filteredCatalog)
	} else {
		ctx.JSON(200, catalog)
	}
}

func productFilterHandler(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	found := false
	for _, prod := range catalog.Products {
		if prod.Id == id {
			ctx.JSON(200, prod)
			found = true
			break
		}
	}

	if !found {
		ctx.JSON(404, gin.H{
			"message": "El producto con id " + ctx.Param("id") + " no fue hallado en el catalogo",
		})
	}
}

func validateAuthToken(ctx *gin.Context) bool {

	token := ctx.GetHeader("token")

	if token != AUTH_TOKEN {
		ctx.JSON(401, gin.H{
			"error": "no tiene permisos para realizar la petición solicitada",
		})
		return false
	}
	ctx.JSON(200, "OK")
	return true
}

func addHandler(ctx *gin.Context) {
	var req ProductDTO
	var lastId int

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{
			"error": fmt.Sprintf("el campo %s es requerido", strings.Split(err.Error(), "'")[3]),
		})
		return
	}

	if !validateAuthToken(ctx) {
		return
	}

	if len(products) > 0 {
		lastId = products[len(products)-1].Id
	} else {
		lastId = catalog.Products[len(catalog.Products)-1].Id
	}
	lastId++
	req.Id = lastId
	products = append(products, req)

	ctx.JSON(200, req)
}

func main() {
	err := readJSON()

	if err != nil {
		panic("error: no se pudo leer el archivo json")
	}

	router := gin.Default()
	router.GET("/greetings", greetingsHandler)
	router.GET("/products", getAllHandler)
	router.GET("/products/:id", productFilterHandler)
	router.POST("/products/add", addHandler)
	router.Run()
}
