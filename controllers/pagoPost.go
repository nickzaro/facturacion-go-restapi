package controllers

import (
	"fmt"

	"github.com/nickzaro/facturacion-go-restapi/models"
)

// ProcesarPagoPost es
func ProcesarPagoPost(pagopost models.PagoPost) {
	var usr models.Usuario
	var pago models.Pago
	pago.Construir(pagopost)
	err := usr.BuscarUsuarioPorPago(pago)
	if err != nil {
		fmt.Println("NO HAY USUARIO PARA ESE PAGO")
		//return {"status":"error", "detalle":"no existe el usuario al que se realiza el pago"}
		return
	}
	//no se aceptan saldo pago positivo
	if pago.MontoPago+usr.PagoUsuario > usr.CargoUsuario { // habria saldo positivo
		fmt.Println("NO SE ACEPTA SALDO POSITIVO POR USUARIO")
		//return {"status":"error", "detalle":"no existe el usuario al que se realiza el pago"}
		return
	}
	usr.PagarFacturas(pago) // devolver el exito formateado
	usr.Almacenar()         // guardo los cambios en usuario
	// return {"status":"OK"}
	return
}
