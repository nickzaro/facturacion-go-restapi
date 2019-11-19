package controllers

import "github.com/nickzaro/facturacion-go-restapi/models"

// ProcesarEventPost procesa el pago en la version 1, sin saldos positivos
func ProcesarEventPost(eventpost models.EventPost) models.EventPost {
	//procesar y devolver la respuesta
	var usr models.Usuario
	usr.BuscarUsuario(eventpost) // retorno un usuario nuevo o de la BD

	var fac models.Factura
	fac = usr.BuscarFactura(eventpost) //retorno la factura, nueva o de la BD

	var car models.Cargo
	car.Construir(eventpost)
	// A
	fac.AgregarCargo(car)            //guarda la ref de cargo en factura
	usr.ActualizarPorCargo(fac, car) //actualiza monto cargo en usuario
	//-A
	// este orden de almacenamiento es indistinto pero tienen que guardarse los 3
	car.Almacenar() //guarda en la BD el cargo
	fac.Almacenar() //guarda en la BD el cargo
	usr.Almacenar() //guarda en la BD el cargo
	return eventpost
}
