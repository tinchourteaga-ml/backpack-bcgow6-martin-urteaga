package products

import (
	"fmt"

	"github.com/tinchourteaga-ml/backpack-bcgow6-martin-urteaga/Go-web/Go-web-III/pkg"
	"github.com/tinchourteaga-ml/backpack-bcgow6-martin-urteaga/Go-web/Go-web-III/pkg/store"
)

var Catalog = ProductsCatalog{}

type Repository interface {
	GetAll(filter pkg.Filter) ([]Product, error)
	Store(id int, name, color, price, stock, code, published, creationDate string) (Product, error)
	LastId() (int, error)
	Delete(id int) error
	Update(id int, name, color, price, stock, code, published, creationDate string) (Product, error)
	UpdateNameAndPrice(id int, name, price string) (Product, error)
}

type ProductsCatalog struct {
	Products []Product
}

type Product struct {
	Id           int
	Name         string
	Color        string
	Price        string
	Stock        string
	Code         string
	Published    string
	CreationDate string
}

type repository struct {
	db store.Store
}

func NewRepository(db store.Store) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetAll(filter pkg.Filter) ([]Product, error) {
	var filteredCatalog ProductsCatalog
	var catalog ProductsCatalog

	r.db.Read(&catalog.Products)

	for _, prod := range catalog.Products {
		if prod.Id == filter.Id || prod.Name == filter.Name || prod.Color == filter.Color || prod.Price == filter.Price || prod.Stock == filter.Stock || prod.Code == filter.Code || prod.Published == filter.Published || prod.CreationDate == filter.CreationDate {
			filteredCatalog.Products = append(filteredCatalog.Products, prod)
		}
	}

	if len(filteredCatalog.Products) > 0 {
		return filteredCatalog.Products, nil
	} else {
		return catalog.Products, nil
	}
}

func (r *repository) LastId() (int, error) {
	var catalog ProductsCatalog
	var lastId int

	r.db.Read(&catalog.Products)

	if len(catalog.Products) > 0 {
		lastId = catalog.Products[len(catalog.Products)-1].Id
	} else {
		lastId = 0
	}
	return lastId, nil
}

func (r *repository) Store(id int, name, color, price, stock, code, published, creationDate string) (Product, error) {
	var catalog ProductsCatalog
	prod := Product{id, name, color, price, stock, code, published, creationDate}

	r.db.Read(&catalog.Products)
	catalog.Products = append(catalog.Products, prod)

	if err := r.db.Write(catalog.Products); err != nil {
		return Product{}, err
	}

	return prod, nil
}

func (r *repository) Delete(id int) error {
	var catalog ProductsCatalog
	var index int
	deleted := false

	r.db.Read(&catalog.Products)

	for i, prod := range catalog.Products {
		if prod.Id == id {
			index = i
			deleted = true
		}
	}

	if !deleted {
		return fmt.Errorf("no existe el producto con id %d", id)
	}

	catalog.Products = append(catalog.Products[:index], catalog.Products[index+1:]...)

	if err := r.db.Write(catalog.Products); err != nil {
		return err
	}

	return nil
}

func (r *repository) Update(id int, name, color, price, stock, code, published, creationDate string) (Product, error) {
	var catalog ProductsCatalog
	p := Product{Name: name, Color: color, Price: price, Stock: stock, Code: code, Published: published, CreationDate: creationDate}
	updated := false

	r.db.Read(&catalog.Products)

	for i, prod := range catalog.Products {
		if prod.Id == id {
			p.Id = id
			catalog.Products[i] = p
			updated = true
		}
	}

	if !updated {
		return Product{}, fmt.Errorf("no existe el producto con id %d", id)
	}

	if err := r.db.Write(catalog.Products); err != nil {
		return Product{}, err
	}

	return p, nil
}

func (r *repository) UpdateNameAndPrice(id int, name, price string) (Product, error) {
	var catalog ProductsCatalog
	var p Product
	updated := false

	r.db.Read(&catalog.Products)

	for i, prod := range catalog.Products {
		if prod.Id == id {
			prod.Name = name
			prod.Price = price
			p = prod
			catalog.Products[i] = p
			updated = true
		}
	}

	if !updated {
		return Product{}, fmt.Errorf("no existe el producto con id %d", id)
	}

	if err := r.db.Write(catalog.Products); err != nil {
		return Product{}, err
	}

	return p, nil
}
