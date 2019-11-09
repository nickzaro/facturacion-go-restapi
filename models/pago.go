package models

import "time"

// Pago estructura de pagos [un pago varios cargos asociados]
type Pago struct {
	FechaPago time.Time //clave date
	MontoPago float32
	//CargosAsociados []Cargo
}
