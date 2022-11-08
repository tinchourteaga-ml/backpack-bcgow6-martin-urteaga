package products

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tinchourteaga-ml/backpack-bcgow6-martin-urteaga/Databases/Go-storage/Go-storage-I/internal/products/domain"
)

func TestStore(t *testing.T) {
	product := domain.Product{
		Name:  "Heladera",
		Qty:   1,
		Price: 120000,
	}

	repo := newRepository()

	result, err := repo.Store(product)

	if err != nil {
		log.Println(err)
	}

	assert.Equal(t, product, result)
}

func TestGetByName(t *testing.T) {
	product := domain.Product{
		Name: "Televisor",
	}

	repo := newRepository()

	result := repo.GetByName(product.Name)

	assert.Equal(t, product.Name, result.Name)
}
