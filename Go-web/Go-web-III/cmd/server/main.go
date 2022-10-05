package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tinchourteaga-ml/backpack-bcgow6-martin-urteaga/Go-web/Go-web-III/cmd/server/handler"
	"github.com/tinchourteaga-ml/backpack-bcgow6-martin-urteaga/Go-web/Go-web-III/internal/products"
)

func main() {
	repo := products.NewRepository()
	service := products.NewService(repo)
	prod := handler.NewProduct(service)

	router := gin.Default()
	pr := router.Group("/products")
	pr.POST("/add", prod.Store())
	pr.GET("/catalog", prod.GetAll())
	pr.GET("/catalog/:id", prod.GetAll())
	pr.DELETE("catalog/:id", prod.Delete())
	pr.PUT("/catalog/:id", prod.Update())
	pr.PATCH("/catalog/:id", prod.UpdateNameAndPrice())
	router.Run()
}
