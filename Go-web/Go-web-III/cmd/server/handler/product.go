package handler

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/tinchourteaga-ml/backpack-bcgow6-martin-urteaga/Go-web/Go-web-III/internal/products"
	"github.com/tinchourteaga-ml/backpack-bcgow6-martin-urteaga/Go-web/Go-web-III/pkg"
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
		var filter pkg.Filter
		filter.Id, _ = strconv.Atoi(ctx.Query("id"))
		filter.Name = ctx.Query("name")
		filter.Color = ctx.Query("color")
		filter.Price = ctx.Query("price")
		filter.Stock = ctx.Query("stock")
		filter.Code = ctx.Query("code")
		filter.Published = ctx.Query("published")
		filter.CreationDate = ctx.Query("creationDate")

		if !validateAuthToken(ctx) {
			return
		}

		prods, err := p.service.GetAll(filter)

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

		if !validateAuthToken(ctx) {
			return
		}

		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(400, gin.H{
				"error": fmt.Sprintf("el campo %s es requerido", strings.Split(err.Error(), "'")[3]),
			})
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

func (p *Product) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if !validateAuthToken(ctx) {
			return
		}

		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

		if err != nil {
			ctx.JSON(400, gin.H{
				"error": "id inválido",
			})
			return
		}

		err = p.service.Delete(int(id))

		if err != nil {
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(200, gin.H{
			"msg": fmt.Sprintf("El producto con id %d ha sido eliminado", id),
		})
	}
}

func (p *Product) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req ProductDTO

		if !validateAuthToken(ctx) {
			return
		}

		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

		if err != nil {
			ctx.JSON(400, gin.H{
				"error": "id inválido",
			})
			return
		}

		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(400, gin.H{
				"error": fmt.Sprintf("el campo %s es requerido", strings.Split(err.Error(), "'")[3]),
			})
			return
		}

		prod, err := p.service.Update(int(id), req.Name, req.Color, req.Price, req.Stock, req.Code, req.Published, req.CreationDate)

		if err != nil {
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(200, prod)
	}
}

func (p *Product) UpdateNameAndPrice() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req ProductDTO

		if !validateAuthToken(ctx) {
			return
		}

		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

		if err != nil {
			ctx.JSON(400, gin.H{
				"error": "id inválido",
			})
			return
		}

		if err := ctx.ShouldBindJSON(&req); err != nil {
			if req.Name == "" {
				ctx.JSON(400, gin.H{
					"error": fmt.Sprint("el campo Name es requerido"),
				})
				return
			}
			if req.Price == "" {
				ctx.JSON(400, gin.H{
					"error": fmt.Sprint("el campo Price es requerido"),
				})
				return
			}
		}

		prod, err := p.service.UpdateNameAndPrice((int(id)), req.Name, req.Price)

		if err != nil {
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(200, prod)
	}
}
