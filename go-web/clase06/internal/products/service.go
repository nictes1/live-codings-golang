package products

import "github.com/nictes1/live-codings-golang/go-web/clase06/internal/domains"

type Service interface {
	GetAll() ([]domains.Product, error)
	Store(nombre, tipo string, cantidad int, precio float64) (domains.Product, error)
	Update(id int, nombre, tipo string, cantidad int, precio float64) (domains.Product, error)
	UpdateName(id int, nombre string) (domains.Product, error)
	Delete(id int) error
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{repository: r}
}

func (s *service) GetAll() ([]domains.Product, error) {
	return s.repository.GetAll()
}

func (s *service) Store(nombre, tipo string, cantidad int, precio float64) (domains.Product, error) {
	return s.repository.Store(nombre, tipo, cantidad, precio)
}

func (s *service) Update(id int, nombre, tipo string, cantidad int, precio float64) (domains.Product, error) {
	return s.repository.Update(id, nombre, tipo, cantidad, precio)
}

func (s *service) UpdateName(id int, nombre string) (domains.Product, error) {
	return s.repository.UpdateName(id, nombre)
}

func (s *service) Delete(id int) error {
	return s.repository.Delete(id)
}
