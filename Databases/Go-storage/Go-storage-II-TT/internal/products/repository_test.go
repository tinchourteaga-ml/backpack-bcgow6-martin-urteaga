package products

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/tinchourteaga-ml/backpack-bcgow6-martin-urteaga/Databases/Go-storage/Go-storage-II-TT/internal/products/domain"
)

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
