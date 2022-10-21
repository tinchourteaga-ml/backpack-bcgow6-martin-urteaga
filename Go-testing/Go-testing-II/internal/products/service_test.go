package products

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tinchourteaga-ml/backpack-bcgow6-martin-urteaga/Go-testing/Go-testing-II/pkg"
)

func TestServiceIntegrationGetAll(t *testing.T) {
	// Arrange.
	filter := pkg.Filter{}
	db := []Product{
		{Id: 1, Name: "volleyball", Color: "white", Price: "5000", Stock: "38", Code: "AB231F", Published: "true", CreationDate: "10-09-2022"},
		{Id: 2, Name: "football", Color: "black", Price: "8000", Stock: "15", Code: "AB233F", Published: "false", CreationDate: "11-12-2021"},
	}

	mockStorage := MockStorage{
		DataMock: db,
	}

	repository := NewRepository(&mockStorage)
	service := NewService(repository)

	// Act.
	results, err := service.GetAll(filter)

	// Assert.
	assert.Nil(t, err)
	assert.Equal(t, db, results)
}

func TestServiceIntegrationGetAllFail(t *testing.T) {
	// Arrange.
	filter := pkg.Filter{}
	expectedErr := errors.New("error")

	mockStorage := MockStorage{
		DataMock:   nil,
		errOnWrite: nil,
		errOnRead:  errors.New("error"),
	}

	repository := NewRepository(&mockStorage)
	service := NewService(repository)

	// Act.
	results, err := service.GetAll(filter)

	// Assert.
	assert.EqualError(t, err, expectedErr.Error())
	assert.Nil(t, results)
}

func TestServiceIntegrationStore(t *testing.T) {
	// Arrange.
	expectedDb := []Product{
		{Id: 1, Name: "volleyball", Color: "white", Price: "5000", Stock: "38", Code: "AB231F", Published: "true", CreationDate: "10-09-2022"},
		{Id: 2, Name: "football", Color: "black", Price: "8000", Stock: "15", Code: "AB233F", Published: "false", CreationDate: "11-12-2021"},
		{Id: 3, Name: "socks", Color: "red", Price: "1000", Stock: "20", Code: "AB235F", Published: "false", CreationDate: "10-10-2022"},
	}

	initialDb := []Product{
		{Id: 1, Name: "volleyball", Color: "white", Price: "5000", Stock: "38", Code: "AB231F", Published: "true", CreationDate: "10-09-2022"},
		{Id: 2, Name: "football", Color: "black", Price: "8000", Stock: "15", Code: "AB233F", Published: "false", CreationDate: "11-12-2021"},
	}

	mockStorage := MockStorage{
		DataMock: initialDb,
	}

	repository := NewRepository(&mockStorage)
	service := NewService(repository)

	// Act.
	productToCreate := Product{Id: 3, Name: "socks", Color: "red", Price: "1000", Stock: "20", Code: "AB235F", Published: "false", CreationDate: "10-10-2022"}

	result, err := service.Store(productToCreate.Name, productToCreate.Color, productToCreate.Price, productToCreate.Stock, productToCreate.Code, productToCreate.Published, productToCreate.CreationDate)

	// Assert.
	assert.Nil(t, err)
	assert.Equal(t, expectedDb, mockStorage.DataMock)
	assert.Equal(t, productToCreate, result)
}

func TestServiceIntegrationStoreFail(t *testing.T) {
	// Arrange.
	expectedErr := errors.New("error")

	mockStorage := MockStorage{
		DataMock:   nil,
		errOnRead:  nil,
		errOnWrite: errors.New("error"),
	}

	repository := NewRepository(&mockStorage)
	service := NewService(repository)

	// Act.
	productToCreate := Product{Id: 3, Name: "socks", Color: "red", Price: "1000", Stock: "20", Code: "AB235F", Published: "false", CreationDate: "10-10-2022"}

	result, err := service.Store(productToCreate.Name, productToCreate.Color, productToCreate.Price, productToCreate.Stock, productToCreate.Code, productToCreate.Published, productToCreate.CreationDate)

	// Assert.
	assert.EqualError(t, err, expectedErr.Error())
	assert.Empty(t, result)
}
