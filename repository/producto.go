package repository

import "mi-api-go/domain"

type ProductosRepository interface {
	GetAll() ([]domain.Producto, error)
	GetById(id string) (*domain.Producto, error)
	Create(producto domain.Producto) (*domain.Producto, error)
	Update(id string, producto domain.Producto) (*domain.Producto, error)
	Delete(id string) error
}
