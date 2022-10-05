package products

import (
	"encoding/json"
	"fmt"

	"github.com/tinchourteaga-ml/backpack-bcgow6-martin-urteaga/Go-web/Go-web-III/pkg"
)

var Catalog = ProductsCatalog{}
var lastId int

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

type repository struct{}

func NewRepository() Repository {
	file, _ := pkg.ReadJSON()
	json.Unmarshal(file, &Catalog)
	return &repository{}
}

func (r *repository) GetAll(filter pkg.Filter) ([]Product, error) {
	var filteredCatalog ProductsCatalog
	for _, prod := range Catalog.Products {
		if prod.Id == filter.Id || prod.Name == filter.Name || prod.Color == filter.Color || prod.Price == filter.Price || prod.Stock == filter.Stock || prod.Code == filter.Code || prod.Published == filter.Published || prod.CreationDate == filter.CreationDate {
			filteredCatalog.Products = append(filteredCatalog.Products, prod)
		}
	}

	if len(filteredCatalog.Products) > 0 {
		return filteredCatalog.Products, nil
	} else {
		return Catalog.Products, nil
	}
}

func (r *repository) LastId() (int, error) {
	if len(Catalog.Products) > 0 {
		lastId = Catalog.Products[len(Catalog.Products)-1].Id
	} else {
		lastId = 0
	}
	return lastId, nil
}

func (r *repository) Store(id int, name, color, price, stock, code, published, creationDate string) (Product, error) {
	prod := Product{id, name, color, price, stock, code, published, creationDate}
	Catalog.Products = append(Catalog.Products, prod)
	return prod, nil
}

func (r *repository) Delete(id int) error {
	deleted := false
	var index int

	for i, prod := range Catalog.Products {
		if prod.Id == id {
			index = i
			deleted = true
		}
	}

	if !deleted {
		return fmt.Errorf("no existe el producto con id %d", id)
	}

	Catalog.Products = append(Catalog.Products[:index], Catalog.Products[index+1:]...)

	return nil
}

func (r *repository) Update(id int, name, color, price, stock, code, published, creationDate string) (Product, error) {
	p := Product{Name: name, Color: color, Price: price, Stock: stock, Code: code, Published: published, CreationDate: creationDate}
	updated := false

	for i, prod := range Catalog.Products {
		if prod.Id == id {
			p.Id = id
			Catalog.Products[i] = p
			updated = true
		}
	}

	if !updated {
		return Product{}, fmt.Errorf("no existe el producto con id %d", id)
	}

	return p, nil
}

func (r *repository) UpdateNameAndPrice(id int, name, price string) (Product, error) {
	var p Product
	updated := false

	for i, prod := range Catalog.Products {
		if prod.Id == id {
			prod.Name = name
			prod.Price = price
			p = prod
			Catalog.Products[i] = p
			updated = true
		}
	}

	if !updated {
		return Product{}, fmt.Errorf("no existe el producto con id %d", id)
	}

	return p, nil
}
