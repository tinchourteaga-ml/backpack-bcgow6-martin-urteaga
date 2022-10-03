package main

import (
	"encoding/json"
	"io/ioutil"
	"strconv"

	"github.com/gin-gonic/gin"
)

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

var Catalog = ProductsCatalog{}

// Hidratamos el catalogo con los productos el archivo json
func readJSON() error {
	file, err := ioutil.ReadFile("products.json")

	if err != nil {
		return err
	}

	json.Unmarshal([]byte(file), &Catalog)
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

	for _, prod := range Catalog.Products {
		if prod.Id == id || prod.Name == name || prod.Color == color || prod.Price == price || prod.Stock == stock || prod.Code == code || prod.Published == published || prod.CreationDate == creationDate {
			filteredCatalog.Products = append(filteredCatalog.Products, prod)
		}
	}

	if len(filteredCatalog.Products) > 0 {
		ctx.JSON(200, filteredCatalog)
	} else {
		ctx.JSON(200, Catalog)
	}
}

func productFilterHandler(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	found := false
	for _, prod := range Catalog.Products {
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

func main() {
	err := readJSON()

	if err != nil {
		panic("error: no se pudo leer el archivo json")
	}

	router := gin.Default()
	router.GET("/greetings", greetingsHandler)
	router.GET("/products", getAllHandler)
	router.GET("/products/:id", productFilterHandler)
	router.Run()
}
