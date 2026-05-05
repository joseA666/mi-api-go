package surrealdb

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"mi-api-go/db"
	"mi-api-go/domain"

	"github.com/surrealdb/surrealdb.go"
	"github.com/surrealdb/surrealdb.go/pkg/models"
)

type surrealProducto struct {
	ID     *models.RecordID `json:"id,omitempty"`
	Nombre string           `json:"nombre"`
	Precio float64          `json:"precio"`
}

func toDomain(sp surrealProducto) domain.Producto {
	p := domain.Producto{Nombre: sp.Nombre, Precio: sp.Precio}
	if sp.ID != nil {
		switch v := sp.ID.ID.(type) {
		case int64:
			p.ID = strconv.FormatInt(v, 10)
		case int:
			p.ID = strconv.Itoa(v)
		case uint64:
			p.ID = strconv.FormatUint(v, 10)
		default:
			p.ID = fmt.Sprintf("%v", v)
		}
	}
	return p
}

// parseRecordID convierte el string ID de vuelta al tipo correcto para SurrealDB.
func parseRecordID(id string) models.RecordID {
	if intID, err := strconv.ParseInt(id, 10, 64); err == nil {
		return models.NewRecordID("productos", intID)
	}
	return models.NewRecordID("productos", id)
}

type ProductoRepository struct{}

func NewProductoRepository() *ProductoRepository {
	return &ProductoRepository{}
}

func (r *ProductoRepository) GetAll() ([]domain.Producto, error) {
	ctx := context.Background()
	resultado, err := surrealdb.Select[[]surrealProducto](ctx, db.SurrealDB, "productos")
	if err != nil {
		return nil, err
	}
	if resultado == nil {
		return []domain.Producto{}, nil
	}
	productos := make([]domain.Producto, len(*resultado))
	for i, sp := range *resultado {
		productos[i] = toDomain(sp)
	}
	return productos, nil
}

func (r *ProductoRepository) GetById(id string) (*domain.Producto, error) {
	ctx := context.Background()
	recordID := parseRecordID(id)
	resultado, err := surrealdb.Select[surrealProducto](ctx, db.SurrealDB, recordID)
	if err != nil {
		return nil, fmt.Errorf("producto no encontrado")
	}
	d := toDomain(*resultado)
	return &d, nil
}

func (r *ProductoRepository) Create(p domain.Producto) (*domain.Producto, error) {
	ctx := context.Background()
	// Usamos UnixNano como ID int64 para evitar ambigüedad de tipos en CBOR.
	recordID := models.NewRecordID("productos", time.Now().UnixNano())
	sp := surrealProducto{Nombre: p.Nombre, Precio: p.Precio}
	resultado, err := surrealdb.Create[surrealProducto](ctx, db.SurrealDB, recordID, sp)
	if err != nil {
		return nil, err
	}
	d := toDomain(*resultado)
	return &d, nil
}

func (r *ProductoRepository) Update(id string, p domain.Producto) (*domain.Producto, error) {
	ctx := context.Background()
	recordID := parseRecordID(id)
	sp := surrealProducto{Nombre: p.Nombre, Precio: p.Precio}
	resultado, err := surrealdb.Update[surrealProducto](ctx, db.SurrealDB, recordID, sp)
	if err != nil {
		return nil, err
	}
	d := toDomain(*resultado)
	return &d, nil
}

func (r *ProductoRepository) Delete(id string) error {
	ctx := context.Background()
	recordID := parseRecordID(id)
	_, err := surrealdb.Delete[surrealProducto](ctx, db.SurrealDB, recordID)
	return err
}
