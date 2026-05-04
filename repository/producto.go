package repository

import "mi-api-go/domain"

type ProductosRepository interface {
	GetAll() ([]domain.Producto, error)
	GetById(id int) (*domain.Producto, error)
	Create(producto domain.Producto) (*domain.Producto, error)
	Update(id int, producto domain.Producto) (*domain.Producto, error)
	Delete(id int) error
}
