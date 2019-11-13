package models

import (
	"testing"

	"github.com/nickzaro/facturacion-go-restapi/utils"
)

func TestConstruirCargo(t *testing.T) {
	var car Cargo
	car.Construir(even)
	if !(car.IDUsuario == even.UserID &&
		car.MesAnioFactura == utils.ConvertirAnioMesString(even.Date) &&
		car.FechaCargo == even.Date &&
		car.Categoria == categorias[even.EventType] &&
		car.Subcategoria == even.EventType &&
		car.MontoCargo == even.Amount &&
		car.EventoID == even.EventID &&
		car.PagosAsociados == nil) {
		t.Error("Error en TestConstruirPago")
	}
}
