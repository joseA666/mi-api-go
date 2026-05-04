package surrealdb

import (
	"context"
	"fmt"

	"mi-api-go/db"
	"mi-api-go/domain"

	"github.com/surrealdb/surrealdb.go"
	"github.com/surrealdb/surrealdb.go/pkg/models"
)

type ProductoRepository struct{}

func NewProductoRepository() *ProductoRepository {
	return &ProductoRepository{}
}

func (r *ProductoRepository) GetAll() ([]domain.Producto, error) {
	ctx := context.Background()
	resultado, err := surrealdb.Select[[]domain.Producto](ctx, db.SurrealDB, "productos")
	if err != nil {
		return nil, err
	}
	return *resultado, nil
}

func (r *ProductoRepository) GetById(id int) (*domain.Producto, error) {
	ctx := context.Background()
	recordID := models.NewRecordID("productos", id)
	resultado, err := surrealdb.Select[domain.Producto](ctx, db.SurrealDB, recordID)
	if err != nil {
		return nil, fmt.Errorf("producto no encontrado")
	}
	return resultado, nil
}

func (r *ProductoRepository) Create(p domain.Producto) (*domain.Producto, error) {
	ctx := context.Background()
	resultado, err := surrealdb.Create[domain.Producto](ctx, db.SurrealDB, "productos", p)
	if err != nil {
		return nil, err
	}
	return resultado, nil
}

func (r *ProductoRepository) Update(id int, p domain.Producto) (*domain.Producto, error) {
	ctx := context.Background()
	recordID := models.NewRecordID("productos", id)
	resultado, err := surrealdb.Update[domain.Producto](ctx, db.SurrealDB, recordID, p)
	if err != nil {
		return nil, err
	}
	return resultado, nil
}

func (r *ProductoRepository) Delete(id int) error {
	ctx := context.Background()
	recordID := models.NewRecordID("productos", id)
	_, err := surrealdb.Delete[domain.Producto](ctx, db.SurrealDB, recordID)
	return err
}
