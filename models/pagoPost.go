package models

// PagoPost estructura que representa el json que recibimos a travez de
// la peticion post de pago
type PagoPost struct {
	Amount   float32 `json:"amount,omitempty"`
	Currency string  `json:"currency,omitempty"`
	UserID   int     `json:"user_id,omitempty"`
}
