package models

import "time"

// Cargo estructura de cargos [un cargo varios pagos asociados]
type Cargo struct {
	FechaCargo     time.Time //clave date
	Categoria      string
	Subcategoria   string
	MontoCargo     float32
	PagosAsociados []Pago
}
