package postgres

import (
	"context"
	"errors"
	"strconv"

	"mi-api-go/db"
	"mi-api-go/domain"

	"github.com/jackc/pgx/v5"
)

type ProductoRepository struct{}

func NewProductoRepository() *ProductoRepository {
	return &ProductoRepository{}
}

func (r *ProductoRepository) GetAll() ([]domain.Producto, error) {
	rows, err := db.PostgresPool.Query(context.Background(),
		"SELECT id::text, nombre, precio FROM productos")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var productos []domain.Producto
	for rows.Next() {
		var p domain.Producto
		if err := rows.Scan(&p.ID, &p.Nombre, &p.Precio); err != nil {
			return nil, err
		}
		productos = append(productos, p)
	}
	return productos, nil
}

func (r *ProductoRepository) GetById(id string) (*domain.Producto, error) {
	intID, err := strconv.Atoi(id)
	if err != nil {
		return nil, errors.New("ID inválido")
	}
	var p domain.Producto
	var rawID int
	err = db.PostgresPool.QueryRow(context.Background(),
		"SELECT id, nombre, precio FROM productos WHERE id = $1", intID).
		Scan(&rawID, &p.Nombre, &p.Precio)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}
	p.ID = strconv.Itoa(rawID)
	return &p, err
}

func (r *ProductoRepository) Create(p domain.Producto) (*domain.Producto, error) {
	var rawID int
	err := db.PostgresPool.QueryRow(context.Background(),
		"INSERT INTO productos (nombre, precio) VALUES ($1, $2) RETURNING id, nombre, precio",
		p.Nombre, p.Precio).
		Scan(&rawID, &p.Nombre, &p.Precio)
	p.ID = strconv.Itoa(rawID)
	return &p, err
}

func (r *ProductoRepository) Update(id string, p domain.Producto) (*domain.Producto, error) {
	intID, err := strconv.Atoi(id)
	if err != nil {
		return nil, errors.New("ID inválido")
	}
	var rawID int
	err = db.PostgresPool.QueryRow(context.Background(),
		"UPDATE productos SET nombre=$1, precio=$2 WHERE id=$3 RETURNING id, nombre, precio",
		p.Nombre, p.Precio, intID).
		Scan(&rawID, &p.Nombre, &p.Precio)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}
	p.ID = strconv.Itoa(rawID)
	return &p, err
}

func (r *ProductoRepository) Delete(id string) error {
	intID, err := strconv.Atoi(id)
	if err != nil {
		return errors.New("ID inválido")
	}
	_, err = db.PostgresPool.Exec(context.Background(),
		"DELETE FROM productos WHERE id=$1", intID)
	return err
}
