package products

import (
	"context"
	"database/sql"
	"errors"
	"log"

	"github.com/tinchourteaga-ml/backpack-bcgow6-martin-urteaga/Databases/Go-storage/Go-storage-II-TT/internal/products/domain"
)

type Repository interface {
	Store(ctx context.Context, product *domain.Product) (*domain.Product, error)
	GetOne(ctx context.Context, productID int) domain.Product
	GetAll(ctx context.Context) ([]domain.Product, error)
	Delete(ctx context.Context, id int) error
	Update(ctx context.Context, product domain.Product) (domain.Product, error)
}

type repository struct {
	db *sql.DB
}

var (
	storeProduct   = "INSERT INTO products(name, qty, price, id_warehouse) VALUES (?, ?, ?, ?)"
	getProductByID = "SELECT name FROM products WHERE products.id = ?"
	getAllProducts = "SELECT id, name, qty, price, id_warehouse FROM products"
	deleteProduct  = "DELETE FROM products WHERE products.id = ?"
	updateProduct  = "UPDATE products AS p SET p.name = ?, p.qty = ?, p.price = ?, p.id_warehouse = ? WHERE p.id = ?"
)

func newRepository(storage *sql.DB) Repository {
	return &repository{
		db: storage,
	}
}

func (repo *repository) Store(ctx context.Context, product *domain.Product) (*domain.Product, error) {
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

func (repo *repository) GetOne(ctx context.Context, productID int) domain.Product {
	var product domain.Product

	rows, err := repo.db.Query(getProductByID, productID)

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

func (repo *repository) GetAll(ctx context.Context) ([]domain.Product, error) {
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

func (repo *repository) Delete(ctx context.Context, id int) error {
	stmt, err := repo.db.Prepare(deleteProduct)

	if err != nil {
		log.Println(err)
		return err
	}

	defer stmt.Close()

	result, err := stmt.Exec(id)

	if err != nil {
		log.Println(err)
		return err
	}

	affect, err := result.RowsAffected()

	if err != nil {
		log.Println(err)
		return err
	}

	if affect < 1 {
		return errors.New("no se encontro el producto indicado")
	}

	return nil
}

func (repo *repository) Update(ctx context.Context, product domain.Product) (domain.Product, error) {
	stmt, err := repo.db.Prepare(updateProduct)

	if err != nil {
		log.Println(err)
		return domain.Product{}, err
	}

	defer stmt.Close()

	_, err = stmt.Exec(product.Name, product.Qty, product.Price, product.WarehouseID, product.ID)

	if err != nil {
		log.Println(err)
		return domain.Product{}, err
	}

	return product, nil
}
