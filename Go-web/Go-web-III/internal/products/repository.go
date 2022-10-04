package products

import (
	"encoding/json"

	"github.com/tinchourteaga-ml/backpack-bcgow6-martin-urteaga/Go-web/Go-web-III/pkg"
)

var Catalog = ProductsCatalog{}
var lastId int

type Repository interface {
	GetAll() ([]Product, error)
	Store(id int, name, color, price, stock, code, published, creationDate string) (Product, error)
	LastId() (int, error)
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

func (r *repository) GetAll() ([]Product, error) {
	return Catalog.Products, nil
}

func (r *repository) LastId() (int, error) {
	return lastId, nil
}

func (r *repository) Store(id int, name, color, price, stock, code, published, creationDate string) (Product, error) {
	prod := Product{id, name, color, price, stock, code, published, creationDate}
	Catalog.Products = append(Catalog.Products, prod)
	lastId = prod.Id
	return prod, nil
}
