package products

import (
	"fmt"

	"github.com/nictes1/live-codings-golang/go-web/clase04/pkg/store"
)

type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"nombre"`
	Type  string  `json:"tipo"`
	Count int     `json:"cantidad"`
	Price float64 `json:"precio"`
}

var ps []Product
var lastID int

//***Importado**//
type Repository interface {
	GetAll() ([]Product, error)
	Store(id int, name, productType string, count int, price float64) (Product, error)
	LastID() (int, error)
	Update(id int, name, productType string, count int, price float64) (Product, error)
}

type repository struct {
	db store.Store
} //struct implementa los metodos de la interfaz

func NewRepository(db store.Store) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Store(id int, nombre, tipo string, cantidad int, precio float64) (Product, error) {

	var ps []Product

	err := r.db.Read(&ps)
	if err != nil {
		return Product{}, err
	}

	p := Product{id, nombre, tipo, cantidad, precio}

	ps = append(ps, p)
	if err := r.db.Write(ps); err != nil {
		return Product{}, err
	}

	return p, nil
}

func (r *repository) GetAll() (products []Product, err error) {

	err = r.db.Read(&products)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (r *repository) LastID() (int, error) {
	var ps []Product
	err := r.db.Read(ps)
	if err != nil {
		return 0, err
	}
	if len(ps) == 0 {
		return 0, nil
	}

	return ps[len(ps)-1].ID, nil
}

func (r *repository) Update(id int, name, productType string, count int, price float64) (Product, error) {
	p := Product{Name: name, Type: productType, Count: count, Price: price}
	updated := false
	for i := range ps {
		if ps[i].ID == id {
			p.ID = id
			ps[i] = p
			updated = true
		}
	}
	if !updated {
		return Product{}, fmt.Errorf("Producto %d no encontrado", id)
	}
	return p, nil
}
