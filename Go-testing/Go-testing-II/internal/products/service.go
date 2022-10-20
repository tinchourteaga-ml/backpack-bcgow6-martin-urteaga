package products

import "github.com/tinchourteaga-ml/backpack-bcgow6-martin-urteaga/Go-testing/Go-testing-II/pkg"

type Service interface {
	GetAll(filter pkg.Filter) ([]Product, error)
	GetSpecific(id int) (Product, error)
	Store(name, color, price, stock, code, published, creationDate string) (Product, error)
	Delete(id int) error
	Update(id int, name, color, price, stock, code, published, creationDate string) (Product, error)
	UpdateNameAndPrice(id int, name, price string) (Product, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetAll(filter pkg.Filter) ([]Product, error) {
	prods, err := s.repository.GetAll(filter)

	if err != nil {
		return nil, err
	}

	return prods, nil
}

func (s *service) GetSpecific(id int) (Product, error) {
	prod, err := s.repository.GetSpecific(id)

	if err != nil {
		return Product{}, err
	}

	return prod, nil
}

func (s *service) Store(name, color, price, stock, code, published, creationDate string) (Product, error) {
	lastId, err := s.repository.LastId()

	if err != nil {
		return Product{}, err
	}

	lastId++
	prod, err := s.repository.Store(lastId, name, color, price, stock, code, published, creationDate)

	if err != nil {
		return Product{}, err
	}

	return prod, nil
}

func (s *service) Delete(id int) error {
	return s.repository.Delete(id)
}

func (s *service) Update(id int, name, color, price, stock, code, published, creationDate string) (Product, error) {
	return s.repository.Update(id, name, color, price, stock, code, published, creationDate)
}

func (s *service) UpdateNameAndPrice(id int, name, price string) (Product, error) {
	return s.repository.UpdateNameAndPrice(id, name, price)
}
