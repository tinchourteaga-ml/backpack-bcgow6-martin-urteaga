package pkg

import (
	"io/ioutil"
)

type Filter struct {
	Id           int
	Name         string
	Color        string
	Price        string
	Stock        string
	Code         string
	Published    string
	CreationDate string
}

func ReadJSON() ([]byte, error) {
	file, err := ioutil.ReadFile("/Users/MURTEAGA/Documents/Bootcamp GO/GITHUB/backpack-bcgow6-martin-urteaga/Go-web/Go-web-III/pkg/products.json")

	if err != nil {
		return nil, err
	}

	return file, nil
}
