package models

import (
	"testing"
)

var even EventPost

func init() {
	even = EventPost{1, 100.00, "USD", 1, "CLASIFICADO", "2019-05-16T00:00:00"}

}

func TestConstruirUsuario(t *testing.T) {
	var usr Usuario
	usr.Construir(even)
	if !(usr.ID == even.UserID && usr.CargoUsuario == 0.00 && usr.PagoUsuario == 0.00) {
		t.Error("Error en TestConstruirUsuario")
	}
}

func TestAgregarFactura(t *testing.T) {
	var usr Usuario
	usr.Construir(even)
	var fac Factura
	fac.Construir(even)

	usr.AgregarFactura(fac)
	existe := false
	for i := range usr.PeriodoFacturas {
		if usr.PeriodoFacturas[i] == fac.MesAnioFactura && usr.Facturas[i] == fac.Referencia {
			existe = true
		}
	}
	if existe == false {
		t.Error("No guardo bien MesAnio o Referencia")
	}
}

func TestAgregarSaldoPorCargo(t *testing.T) {
	var usr Usuario
	usr.Construir(even)
	var car Cargo
	car.Construir(even)
	usr.AgregarSaldoPorCargo(car)
	if usr.CargoUsuario != car.MontoCargo {
		t.Error("No agrego bien el saldo de cargo a usuario")
	}
	usr.AgregarSaldoPorCargo(car)
	if usr.CargoUsuario != car.MontoCargo*2 {
		t.Error("No agrego bien el saldo de cargo*2 a usuario")
	}
}
