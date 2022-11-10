package products

import (
	"context"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/tinchourteaga-ml/backpack-bcgow6-martin-urteaga/Databases/Go-storage/Go-storage-II-TT/internal/products/domain"
)

// Error: no se pudo cumplir con mock.ExpectationsWereMet() de la sentencia SELECT
/*
	func TestStoreMock(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectPrepare("INSERT INTO products")
		mock.ExpectExec("INSERT INTO products").WillReturnResult(sqlmock.NewResult(5, 1))

		columns := []string{"id", "name", "qty", "price", "id_warehouse"}
		rows := sqlmock.NewRows(columns)
		productID := 5

		rows.AddRow(productID, "", "", "", "")
		mock.ExpectQuery("SELECT .* FROM products").WithArgs(productID).WillReturnRows(rows)

		repository := newRepository(db)
		ctx := context.TODO()
		product := domain.Product{
			ID: productID,
		}

		getResult := repository.GetOne(ctx, productID)
		assert.Empty(t, getResult)

		result, err := repository.Store(ctx, &product)
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, product.ID, result.ID)
		assert.NoError(t, mock.ExpectationsWereMet())
	}
*/
func Test_sqlRepositoryStore(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	p := &domain.Product{
		ID:          1,
		Name:        "test",
		Qty:         2,
		Price:       1200,
		WarehouseID: 3,
	}

	ctx := context.TODO()
	repo := newRepository(db)

	mock.ExpectPrepare(regexp.QuoteMeta("INSERT INTO products(name, qty, price, id_warehouse) VALUES (?, ?, ?, ?)")).
		ExpectExec().
		WithArgs(p.Name, p.Qty, p.Price, p.WarehouseID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	// Store
	prod, err := repo.Store(ctx, p)
	assert.NoError(t, err)
	assert.Equal(t, p.ID, prod.ID)

	columns := []string{"id", "name", "qty", "price", "id_warehouse"}
	rows := sqlmock.NewRows(columns)
	rows.AddRow(p.ID, p.Name, p.Qty, p.Price, p.WarehouseID)
	mock.ExpectPrepare(regexp.QuoteMeta("SELECT name FROM products WHERE products.id = ?")).
		ExpectQuery().
		WithArgs(p.ID).
		WillReturnRows(rows)

	assert.Equal(t, p, prod)
}
