package domain

type Producto struct {
	ID     string  `json:"id,omitempty"`
	Nombre string  `json:"nombre"`
	Precio float64 `json:"precio"`
}
