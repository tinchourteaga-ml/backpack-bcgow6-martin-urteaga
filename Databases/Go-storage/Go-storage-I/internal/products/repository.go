package products

import (
	"database/sql"
	"log"

	"github.com/tinchourteaga-ml/backpack-bcgow6-martin-urteaga/Databases/Go-storage/Go-storage-I/internal/products/domain"
)

type Repository interface {
	Store(product *domain.Product) (*domain.Product, error)
	GetByName(productName string) domain.Product
	GetAll() ([]domain.Product, error)
}

type repository struct {
	db *sql.DB
}

var (
	storeProduct     = "INSERT INTO products(name, qty, price, id_warehouse) VALUES (?, ?, ?, ?)"
	getProductByName = "SELECT name FROM products WHERE products.name = ?"
	getAllProducts   = "SELECT id, name, qty, price, id_warehouse FROM products"
)

func newRepository(storage *sql.DB) Repository {
	return &repository{
		db: storage,
	}
}

func (repo *repository) Store(product *domain.Product) (*domain.Product, error) {
	stmt, err := repo.db.Prepare(storeProduct)

	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	result, err := stmt.Exec(product.Name, product.Qty, product.Price, product.WarehouseID)

	if err != nil {
		return nil, err
	}

	insertedID, err := result.LastInsertId()

	if err != nil {
		return nil, err
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

func (repo *repository) GetAll() ([]domain.Product, error) {
	var products []domain.Product

	rows, err := repo.db.Query(getAllProducts)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	for rows.Next() {
		var product domain.Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Qty, &product.Price, &product.WarehouseID); err != nil {
			log.Println(err)
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}
