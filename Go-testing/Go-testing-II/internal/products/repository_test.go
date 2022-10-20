package products

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tinchourteaga-ml/backpack-bcgow6-martin-urteaga/Go-testing/Go-testing-II/pkg"
)

/* =========== STUB TEST =========== */
type StubStore struct{}

func (s *StubStore) Read(data interface{}) error {
	p1 := Product{Id: 1, Name: "volleyball", Color: "white", Price: "5000", Stock: "38", Code: "AB231F", Published: "true", CreationDate: "10-09-2022"}
	p2 := Product{Id: 2, Name: "football", Color: "black", Price: "8000", Stock: "15", Code: "AB233F", Published: "false", CreationDate: "11-12-2021"}

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
	p1 := Product{Id: 1, Name: "volleyball", Color: "white", Price: "5000", Stock: "38", Code: "AB231F", Published: "true", CreationDate: "10-09-2022"}
	p2 := Product{Id: 2, Name: "football", Color: "black", Price: "8000", Stock: "15", Code: "AB233F", Published: "false", CreationDate: "11-12-2021"}
	filter := pkg.Filter{}
	stubStore := StubStore{}
	repo := NewRepository(&stubStore)
	expectedProducts := []Product{p1, p2}
	products, _ := repo.GetAll(filter)

	assert.Equal(t, expectedProducts, products)
}

/* =========== MOCK TEST =========== */
type MockedStorage struct {
	ReadWasCalled       bool
	ProductBeforeUpdate Product
}

func (m *MockedStorage) Read(data interface{}) error {
	m.ProductBeforeUpdate = Product{Id: 1, Name: "volleyball", Color: "white", Price: "5000", Stock: "38", Code: "AB231F", Published: "true", CreationDate: "10-09-2022"}

	switch t := data.(type) {
	case *[]Product:
		*t = append(*t, m.ProductBeforeUpdate)
	}

	m.ReadWasCalled = true

	return nil
}

func (m *MockedStorage) Write(data interface{}) error {
	return nil
}

func TestUpdateNameAndPrice(t *testing.T) {
	mockedStorage := MockedStorage{}
	repo := NewRepository(&mockedStorage)
	expectedProduct := Product{Id: 1, Name: "short", Color: "white", Price: "2500", Stock: "38", Code: "AB231F", Published: "true", CreationDate: "10-09-2022"}
	product, _ := repo.UpdateNameAndPrice(1, "short", "2500")

	assert.Equal(t, expectedProduct, product)
	assert.True(t, mockedStorage.ReadWasCalled)
}
