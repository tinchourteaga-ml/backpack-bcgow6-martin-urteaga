package handler

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/tinchourteaga-ml/backpack-bcgow6-martin-urteaga/Go-web/Go-web-III/internal/products"
	"github.com/tinchourteaga-ml/backpack-bcgow6-martin-urteaga/Go-web/Go-web-III/pkg"
	"github.com/tinchourteaga-ml/backpack-bcgow6-martin-urteaga/Go-web/Go-web-III/pkg/web"
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

	if token != os.Getenv("AUTH_TOKEN") {
		ctx.JSON(401, web.NewResponse(401, nil, "no tiene permisos para realizar la petición solicitada"))
		return false
	}
	return true
}

// ListProducts godoc
// @Summary List products
// @Tags Products
// @Description get products
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Success 200 {object} web.Response
// @Router /products/catalog [get]
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
			ctx.JSON(400, web.NewResponse(400, nil, fmt.Sprint(err)))
			return
		}

		ctx.JSON(200, web.NewResponse(200, prods, ""))
	}
}

func (p *Product) GetSpecific() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, _ := strconv.Atoi(ctx.Param("id"))

		if !validateAuthToken(ctx) {
			return
		}

		prod, err := p.service.GetSpecific(id)

		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, fmt.Sprint(err)))
			return
		}

		ctx.JSON(200, web.NewResponse(200, prod, ""))
	}
}

// StoreProducts godoc
// @Summary Store products
// @Tags Products
// @Description store products
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param product body string true "Product to store"
// @Success 200 {object} web.Response
// @Router /products/add [post]
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
			ctx.JSON(400, web.NewResponse(400, nil, fmt.Sprint(err)))
			return
		}

		ctx.JSON(200, web.NewResponse(200, prod, ""))
	}
}

func (p *Product) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if !validateAuthToken(ctx) {
			return
		}

		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, fmt.Sprint(err)))
			return
		}

		err = p.service.Delete(int(id))

		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, fmt.Sprint(err)))
			return
		}
		ctx.JSON(200, web.NewResponse(200, fmt.Sprintf("El producto con id %d ha sido eliminado", id), ""))
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
			ctx.JSON(400, web.NewResponse(400, nil, fmt.Sprint(err)))
			return
		}

		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, fmt.Sprintf("el campo %s es requerido", strings.Split(err.Error(), "'")[3])))
			return
		}

		prod, err := p.service.Update(int(id), req.Name, req.Color, req.Price, req.Stock, req.Code, req.Published, req.CreationDate)

		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, fmt.Sprint(err)))
			return
		}
		ctx.JSON(200, web.NewResponse(200, prod, ""))
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
			ctx.JSON(404, web.NewResponse(404, nil, "id inválido"))
			return
		}

		if err := ctx.ShouldBindJSON(&req); err != nil {
			if req.Name == "" {
				ctx.JSON(400, web.NewResponse(400, nil, "el campo Name es requerido"))
				return
			}
			if req.Price == "" {
				ctx.JSON(400, web.NewResponse(400, nil, "el campo Price es requerido"))
				return
			}
		}

		prod, err := p.service.UpdateNameAndPrice((int(id)), req.Name, req.Price)

		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, fmt.Sprint(err)))
			return
		}

		ctx.JSON(200, web.NewResponse(200, prod, ""))
	}
}
