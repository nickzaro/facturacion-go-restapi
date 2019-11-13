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
	fac.AgregarCargo(car)         //guarda la ref de cargo en factura
	usr.AgregarSaldoPorCargo(car) //actualiza monto cargo en usuario
	//-A
	// este orden de almacenamiento es indistinto pero tienen que guardarse los 3
	car.Almacenar() //guarda en la BD el cargo
	fac.Almacenar() //guarda en la BD el cargo
	usr.Almacenar() //guarda en la BD el cargo
	return eventpost
	//Escribir estos metodos
}

/*
// ProcesarEventPostv1 procesa el pago en la version 1, sin saldos positivos
func ProcesarEventPostv1(eventpost models.EventPost) models.EventPost {
	//procesar y devolver la respuesta
	usr, errusr := BuscarUsuario(eventpost.UserID)
	if errusr != nil {
		usr.Construir(eventpost)
		AlmacenarUsuario(usr)
	}
	fac, err := BuscarFactura(usr, eventpost) // en factura.go en controllers
	if err != nil {
		fac.Construir(eventpost)
		AlmacenarFactura(fac)   // almacena la factura en la BD
		usr.AgregarFactura(fac) //guarda la ref de factura en usuario
	}
	var car models.Cargo
	car.Construir(eventpost)
	AlmacenarCargo(car)           //guarda cargo en la BD
	fac.AgregarCargo(car)         //guarda la ref de cargo en factura
	usr.AgregarSaldoPorCargo(car) //actualiza monto cargo en usuario
	//ActualizarFactura(fac)
	//ActualizarUsuario(usr)
	return eventpost
}

// Creo que el almacenar de usr, fac y car deberia de hacerse al final
// osea el Almacenar al final para escribir los cambios o el
*/
