package models

import (
	"testing"

	"github.com/nickzaro/facturacion-go-restapi/utils"
)

func TestConstruirFactura(t *testing.T) {
	var fac Factura
	fac.Construir(even)
	if !(fac.IDUsuario == even.UserID &&
		fac.MesAnioFactura == utils.ConvertirAnioMesString(even.Date)) {
		t.Error("Error en TestConstruirPago")
	}
}

func TestAgregarCargoFactura(t *testing.T) {
	var fac Factura
	fac.Construir(even) // reseteo la factura
	var car Cargo
	car.Construir(even)
	fac.AgregarCargo(car)
	if fac.CargoFactura != car.MontoCargo {
		t.Error("No agrego bien el saldo de cargo a usuario")
	}
	existe := false
	for i := range fac.Cargos {
		if fac.Cargos[i] == car.Referencia {
			existe = true
		}
	}
	if existe == false {
		t.Error("No guardo bien MesAnio o Referencia")
	}

}
