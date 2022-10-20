package products

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tinchourteaga-ml/backpack-bcgow6-martin-urteaga/Go-testing/Go-testing-II/pkg"
)

type StubStore struct{}

func (s *StubStore) Read(data interface{}) error {
	p1 := Product{Name: "volleyball", Color: "white", Price: "5000", Stock: "38", Code: "AB231F", Published: "true", CreationDate: "10-09-2022"}
	p2 := Product{Name: "football", Color: "black", Price: "8000", Stock: "15", Code: "AB233F", Published: "false", CreationDate: "11-12-2021"}

	// Como recibimos una interface{}, necesitamos identificar que estrcutura es la que nos llega. Armamos este switch para hacer algo en función de lo que nos llega.
	// En este caso, nos llega una lista de productos del catalogo. Identificamos que es una lista de productos y hacemos el append de estos productos hardcodeados para el test.
	switch t := data.(type) {
	case *[]Product:
		*t = append(*t, p1, p2)
	}

	return nil
}

func (s *StubStore) Write(data interface{}) error {
	return nil
}

func TestGetAll(t *testing.T) {
	p1 := Product{Name: "volleyball", Color: "white", Price: "5000", Stock: "38", Code: "AB231F", Published: "true", CreationDate: "10-09-2022"}
	p2 := Product{Name: "football", Color: "black", Price: "8000", Stock: "15", Code: "AB233F", Published: "false", CreationDate: "11-12-2021"}
	filter := pkg.Filter{}
	stubStore := StubStore{}
	repo := NewRepository(&stubStore)
	expectedProducts := []Product{p1, p2}
	products, _ := repo.GetAll(filter)

	assert.Equal(t, expectedProducts, products)
}
