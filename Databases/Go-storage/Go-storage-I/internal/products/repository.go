package products

import (
	"database/sql"
	"log"

	"github.com/tinchourteaga-ml/backpack-bcgow6-martin-urteaga/Databases/Go-storage/Go-storage-I/internal/products/domain"
)

type Repository interface {
	Store(product domain.Product) (domain.Product, error)
	GetByName(productName string) domain.Product
}

type repository struct {
	db *sql.DB
}

var (
	storeProduct     = "INSERT INTO products(name, qty, price, id_warehouse) VALUES (?, ?, ?, ?)"
	getProductByName = "SELECT name FROM products WHERE products.name = ?"
)

func newRepository(storage *sql.DB) Repository {
	return &repository{
		db: storage,
	}
}

func (repo *repository) Store(product domain.Product) (domain.Product, error) {
	stmt, err := repo.db.Prepare(storeProduct)

	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	result, err := stmt.Exec(product.Name, product.Qty, product.Price, product.WarehouseID)

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

	rows, err := repo.db.Query(getProductByName, productName)

	if err != nil {
		log.Println(err)
		return product
	}

	for rows.Next() {
		if err := rows.Scan(&product.Name); err != nil {
			log.Println(err)
			return product
		}
	}

	return product
}
