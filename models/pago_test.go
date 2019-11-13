package models

import (
	"testing"

	"github.com/nickzaro/facturacion-go-restapi/utils"
)

var pagopost PagoPost

func init() {
	pagopost = PagoPost{150.10, "USD", 1}
}

func TestConstruirPago(t *testing.T) {
	var pago Pago
	pago.Construir(pagopost)
	if !(pago.IDUsuario == pagopost.UserID &&
		pago.MontoPago == utils.ConvertirAPesos(pagopost.Amount, pagopost.Currency)) {
		t.Error("Error en TestConstruirPago")
	}
}
