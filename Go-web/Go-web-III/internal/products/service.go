package products

type Service interface {
	GetAll() ([]Product, error)
	Store(name, color, price, stock, code, published, creationDate string) (Product, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetAll() ([]Product, error) {
	prods, err := s.repository.GetAll()

	if err != nil {
		return nil, err
	}

	return prods, nil
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
