package test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/tinchourteaga-ml/backpack-bcgow6-martin-urteaga/Go-testing/Go-testing-II/cmd/server/handler"
	"github.com/tinchourteaga-ml/backpack-bcgow6-martin-urteaga/Go-testing/Go-testing-II/internal/products"
)

func createServer(mockStorage products.MockStorage) *gin.Engine {

	repo := products.NewRepository(&mockStorage)
	service := products.NewService(repo)
	prod := handler.NewProduct(service)
	router := gin.Default()

	pr := router.Group("/products")
	pr.POST("/add", prod.Store())
	pr.GET("/catalog", prod.GetAll())
	pr.GET("/catalog/:id", prod.GetSpecific())
	pr.DELETE("catalog/:id", prod.Delete())
	pr.PUT("/catalog/:id", prod.Update())
	pr.PATCH("/catalog/:id", prod.UpdateNameAndPrice())

	return router
}

func createRequestTest(method, url, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")

	return req, httptest.NewRecorder()
}

func TestDeleteProduct(t *testing.T) {
	initialDb := []products.Product{
		{Id: 1, Name: "volleyball", Color: "white", Price: "5000", Stock: "38", Code: "AB231F", Published: "true", CreationDate: "10-09-2022"},
		{Id: 2, Name: "football", Color: "black", Price: "8000", Stock: "15", Code: "AB233F", Published: "false", CreationDate: "11-12-2021"},
	}

	mockStorage := products.MockStorage{
		DataMock: initialDb,
	}

	router := createServer(mockStorage)
	req, recorder := createRequestTest(http.MethodDelete, "/products/catalog/1", "")
	router.ServeHTTP(recorder, req)
	assert.Equal(t, 200, recorder.Code)
}
