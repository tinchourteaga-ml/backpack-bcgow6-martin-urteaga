package products

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tinchourteaga-ml/backpack-bcgow6-martin-urteaga/Databases/Go-storage/Go-storage-I/internal/products/domain"
	"github.com/tinchourteaga-ml/backpack-bcgow6-martin-urteaga/Databases/Go-storage/Go-storage-I/pkg/store"
)

/* func TestStore(t *testing.T) {
	db := store.Init()
	product := domain.Product{
		Name:        "Heladera",
		Qty:         1,
		Price:       120000,
		WarehouseID: 1,
	}

	repo := newRepository(db)

	result, err := repo.Store(&product)

	if err != nil {
		log.Println(err)
	}

	assert.Equal(t, &product, result)
}

func TestGetByName(t *testing.T) {
	db := store.Init()
	product := domain.Product{
		Name: "televisor",
	}

	repo := newRepository(db)

	result := repo.GetByName(product.Name)

	assert.Equal(t, product.Name, result.Name)
} */

func TestGetAll(t *testing.T) {
	db := store.Init()

	var products = []domain.Product{
		{
			ID:          1,
			Name:        "heladera",
			Qty:         1,
			Price:       120000,
			WarehouseID: 0,
		},
		{
			ID:          2,
			Name:        "televisor",
			Qty:         3,
			Price:       75000,
			WarehouseID: 0,
		},
	}

	repo := newRepository(db)

	result, err := repo.GetAll()

	assert.Nil(t, err)
	assert.Equal(t, products, result)
}
