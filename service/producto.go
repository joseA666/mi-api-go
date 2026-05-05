package service

import (
	"mi-api-go/domain"
	"mi-api-go/repository"
)

type ProductosService struct {
	repo repository.ProductosRepository
}

func NewProductosService(repo repository.ProductosRepository) *ProductosService {
	return &ProductosService{repo: repo}
}

func (s *ProductosService) GetAll() ([]domain.Producto, error) {
	return s.repo.GetAll()
}

func (s *ProductosService) GetById(id string) (*domain.Producto, error) {
	return s.repo.GetById(id)
}

func (s *ProductosService) Create(producto domain.Producto) (*domain.Producto, error) {
	return s.repo.Create(producto)
}

func (s *ProductosService) Update(id string, producto domain.Producto) (*domain.Producto, error) {
	return s.repo.Update(id, producto)
}

func (s *ProductosService) Delete(id string) error {
	return s.repo.Delete(id)
}
