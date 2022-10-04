package handler

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/tinchourteaga-ml/backpack-bcgow6-martin-urteaga/Go-web/Go-web-III/internal/products"
)

var AUTH_TOKEN = "123ABC"

type ProductDTO struct {
	Name         string `json:"name" binding:"required"`
	Color        string `json:"color" binding:"required"`
	Price        string `json:"price" binding:"required"`
	Stock        string `json:"stock" binding:"required"`
	Code         string `json:"code" binding:"required"`
	Published    string `json:"published" binding:"required"`
	CreationDate string `json:"creationDate" binding:"required"`
}

type Product struct {
	service products.Service
}

func NewProduct(prod products.Service) *Product {
	return &Product{
		service: prod,
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
	return true
}

func (p *Product) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if !validateAuthToken(ctx) {
			return
		}

		prods, err := p.service.GetAll()

		if err != nil {
			ctx.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(200, prods)
	}
}

func (p *Product) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req ProductDTO

		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(400, gin.H{
				"error": fmt.Sprintf("el campo %s es requerido", strings.Split(err.Error(), "'")[3]),
			})
			return
		}

		if !validateAuthToken(ctx) {
			return
		}

		prod, err := p.service.Store(req.Name, req.Color, req.Price, req.Stock, req.Code, req.Published, req.CreationDate)

		if err != nil {
			ctx.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(200, prod)
	}
}
