package products

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tinchourteaga-ml/backpack-bcgow6-martin-urteaga/Databases/Go-storage/Go-storage-I-II-TM/internal/products/domain"
	"github.com/tinchourteaga-ml/backpack-bcgow6-martin-urteaga/Databases/Go-storage/Go-storage-I-II-TM/pkg/store"
)

var db = store.Init()

func TestStore(t *testing.T) {
	product := domain.Product{
		Name:        "Heladera",
		Qty:         1,
		Price:       120000,
		WarehouseID: 1,
	}

	repo := newRepository(db)

	result, err := repo.Store(&product)

	assert.Nil(t, err)
	assert.Equal(t, &product, result)
}

func TestGetByName(t *testing.T) {
	product := domain.Product{
		Name: "televisor",
	}

	repo := newRepository(db)

	result := repo.GetByName(product.Name)

	assert.Equal(t, product.Name, result.Name)
}

func TestGetAll(t *testing.T) {
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

func TestDelete(t *testing.T) {
	repo := newRepository(db)

	err := repo.Delete(1)

	assert.Nil(t, err)
}

func TestUpdate(t *testing.T) {
	product := domain.Product{
		ID:          2,
		Name:        "parlante",
		Qty:         6,
		Price:       38500,
		WarehouseID: 7,
	}

	repo := newRepository(db)

	result, err := repo.Update(product)

	assert.Nil(t, err)
	assert.Equal(t, product, result)
}
