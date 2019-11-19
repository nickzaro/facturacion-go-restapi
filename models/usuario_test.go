package models

import (
	"testing"
)

var even EventPost

func init() {
	even = EventPost{1, 100.00, "USD", 1, "CLASIFICADO", "2019-05-16T00:00:00"}
	//pagopost = PagoPost{100.00, "USD", 1}

}

func TestConstruirPorPagoUsuario(t *testing.T) {
	var usr Usuario
	var pag Pago
	pag.Construir(pagopost)
	usr.ConstruirPorPago(pag)
	if !(usr.ID == pagopost.UserID && usr.CargoUsuario == 0.00 && usr.PagoUsuario == 0.00) {
		t.Error("Error en TestConstruirPorPagoUsuario")
	}
}

func TestConstruirUsuario(t *testing.T) {
	var usr Usuario
	usr.ConstruirPorEvento(even)
	if !(usr.ID == even.UserID && usr.CargoUsuario == 0.00 && usr.PagoUsuario == 0.00) {
		t.Error("Error en TestConstruirUsuario")
	}
}

func TestAgregarFactura(t *testing.T) {
	var usr Usuario
	usr.ConstruirPorEvento(even)
	var fac Factura
	fac.Construir(even)

	usr.AgregarFactura(fac)
	existe := false
	for i := range usr.ReferenciaFacturas {
		if usr.ReferenciaFacturas[i].PeriodoFactura == fac.MesAnioFactura &&
			usr.ReferenciaFacturas[i].IDReferenciaFactura == fac.Referencia {
			existe = true
		}
	}
	if existe == false {
		t.Error("No guardo bien MesAnio o Referencia")
	}
}

/*
func TestAgregarSaldoPorCargo(t *testing.T) {
	var usr Usuario
	usr.ConstruirPorEvento(even)
	var car Cargo
	car.Construir(even)
	usr.ActualizarPorCargo(car)
	if usr.CargoUsuario != car.MontoCargo {
		t.Error("No agrego bien el saldo de cargo a usuario")
	}
	usr.ActualizarPorCargo(car)
	if usr.CargoUsuario != car.MontoCargo*2 {
		t.Error("No agrego bien el saldo de cargo*2 a usuario")
	}
}
*/
