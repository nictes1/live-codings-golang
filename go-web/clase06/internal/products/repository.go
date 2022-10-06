package products

import (
	"fmt"

	"github.com/nictes1/live-codings-golang/go-web/clase06/internal/domains"
)

type Repository interface {
	GetAll() ([]domains.Product, error)
	Store(nombre, tipo string, cantidad int, precio float64) (domains.Product, error)
	Update(id int, nombre, tipo string, cantidad int, precio float64) (domains.Product, error)
	UpdateName(id int, nombre string) (domains.Product, error)
	Delete(id int) error
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

var (
	products []domains.Product
	lastID   int = 0
)

// Get products
func (r *repository) GetAll() ([]domains.Product, error) {
	return products, nil
}

// Store de products
func (r *repository) Store(nombre, tipo string, cantidad int, precio float64) (domains.Product, error) {
	p := domains.Product{
		Nombre:   nombre,
		Tipo:     tipo,
		Cantidad: cantidad,
		Precio:   precio,
	}

	lastID++
	p.Id = lastID
	products = append(products, p)
	return p, nil
}

// Update product
func (r *repository) Update(id int, nombre, tipo string, cantidad int, precio float64) (domains.Product, error) {
	p := domains.Product{
		Nombre:   nombre,
		Tipo:     tipo,
		Cantidad: cantidad,
		Precio:   precio,
	}
	var updated bool
	for i := range products {
		if products[i].Id == id {
			p.Id = products[i].Id
			products[i] = p
			updated = true
		}
	}

	if !updated {
		return domains.Product{}, fmt.Errorf("product id %d not exists", id)
	}
	return p, nil
}

// Update Name product
func (r *repository) UpdateName(id int, nombre string) (domains.Product, error) {
	var updated bool
	p := domains.Product{}
	for i := range products {
		if products[i].Id == id {
			products[i].Nombre = nombre
			p = products[i]
			updated = true
		}
	}

	if !updated {
		return domains.Product{}, fmt.Errorf("product id %d not exists", id)
	}
	return p, nil
}

// Delete product
func (r *repository) Delete(id int) error {
	var deleted bool
	var pos int
	for i := range products {
		if products[i].Id == id {
			pos = i
			deleted = true
		}
	}

	if !deleted {
		return fmt.Errorf("product id %d not exists", id)
	}

	products = append(products[:pos], products[pos+1:]...)
	return nil
}
