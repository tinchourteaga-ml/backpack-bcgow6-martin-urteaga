package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/tinchourteaga-ml/backpack-bcgow6-martin-urteaga/Go-web/Go-web-III/cmd/server/handler"
	"github.com/tinchourteaga-ml/backpack-bcgow6-martin-urteaga/Go-web/Go-web-III/docs"
	"github.com/tinchourteaga-ml/backpack-bcgow6-martin-urteaga/Go-web/Go-web-III/internal/products"
	"github.com/tinchourteaga-ml/backpack-bcgow6-martin-urteaga/Go-web/Go-web-III/pkg/store"
)

// @title MELI Bootcamp API
// @version 1.0
// @description This API Handle MELI Products
// @termsOfService https://developers.mercadolibre.com.ar/es_ar/terminos-y-condiciones

// @contact.name API Support
// @contact.url https://developers.mercadolibre.com.ar/support

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	err := godotenv.Load()

	if err != nil {
		fmt.Println(errors.New("error: no se han podido leer las variables de entorno"))
	}

	db := store.New(store.FileType, "./products.json")
	repo := products.NewRepository(db)
	service := products.NewService(repo)
	prod := handler.NewProduct(service)

	router := gin.Default()

	docs.SwaggerInfo.Host = os.Getenv("HOST")
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	pr := router.Group("/products")
	pr.POST("/add", prod.Store())
	pr.GET("/catalog", prod.GetAll())
	pr.GET("/catalog/:id", prod.GetSpecific())
	pr.DELETE("catalog/:id", prod.Delete())
	pr.PUT("/catalog/:id", prod.Update())
	pr.PATCH("/catalog/:id", prod.UpdateNameAndPrice())
	router.Run()
}
