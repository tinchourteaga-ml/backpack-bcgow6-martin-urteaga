package products

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tinchourteaga-ml/backpack-bcgow6-martin-urteaga/Databases/Go-storage/Go-storage-I/internal/products/domain"
	"github.com/tinchourteaga-ml/backpack-bcgow6-martin-urteaga/Databases/Go-storage/Go-storage-I/pkg/store"
)

func TestStore(t *testing.T) {
	db := store.Init()
	product := domain.Product{
		ID:          7,
		Name:        "Heladera",
		Qty:         1,
		Price:       120000,
		WarehouseID: 1,
	}

	repo := newRepository(db)

	result, err := repo.Store(product)

	if err != nil {
		log.Println(err)
	}

	assert.Equal(t, product, result)
}

func TestGetByName(t *testing.T) {
	db := store.Init()
	product := domain.Product{
		Name: "televisor",
	}

	repo := newRepository(db)

	result := repo.GetByName(product.Name)

	assert.Equal(t, product.Name, result.Name)
}
