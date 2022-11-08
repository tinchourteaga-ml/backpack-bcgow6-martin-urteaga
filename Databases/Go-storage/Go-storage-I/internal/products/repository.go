package products

import (
	"log"

	"github.com/tinchourteaga-ml/backpack-bcgow6-martin-urteaga/Databases/Go-storage/Go-storage-I/internal/products/domain"
	"github.com/tinchourteaga-ml/backpack-bcgow6-martin-urteaga/Databases/Go-storage/Go-storage-I/pkg/store"
)

type Repository interface {
	Store(product domain.Product) (domain.Product, error)
	GetByName(productName string) domain.Product
}

type repository struct{}

var (
	db               = store.StorageDB
	storeProduct     = "INSERT INTO products(name, qty, price) VALUES (?, ?, ?)"
	getProductByName = "SELECT * FROM products WHERE name = ?"
)

func newRepository() Repository {
	return &repository{}
}

func (repo *repository) Store(product domain.Product) (domain.Product, error) {
	stmt, err := db.Prepare(storeProduct)

	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	result, err := stmt.Exec(product.Name, product.Qty, product.Price)

	if err != nil {
		return domain.Product{}, err
	}

	insertedID, err := result.LastInsertId()

	if err != nil {
		return domain.Product{}, err
	}

	product.ID = int(insertedID)

	return product, nil
}

func (repo *repository) GetByName(productName string) domain.Product {
	var product domain.Product

	rows, err := db.Query(getProductByName, productName)

	if err != nil {
		log.Println(err)
		return product
	}

	for rows.Next() {
		if err := rows.Scan(&product.ID, &product.Name, &product.Qty, &product.Price); err != nil {
			log.Println(err)
			return product
		}
	}

	return product
}
