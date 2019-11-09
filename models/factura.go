package models

import "time"

// Factura estructura de una factura[son factura por mes]
type Factura struct {
	FechaFactura time.Time //seria la clave año-mes
	DeudaFactura float32
	Pagofactura  float32
	Cargos       []Cargo
	Pagos        []Pago
}
