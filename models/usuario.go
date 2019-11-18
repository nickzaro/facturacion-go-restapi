package models

import (
	"context"
	"fmt"
	"log"

	"github.com/nickzaro/facturacion-go-restapi/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Usuario estructura de los documentos en la base de datos
type Usuario struct {
	Referencia   primitive.ObjectID `bson:"_id" json:"_id"`
	ID           int                `bson:"id" json:"id"`
	CargoUsuario float64            `bson:"cargo_usuario" json:"cargo_usuario"`
	PagoUsuario  float64            `bson:"pago_usuario" json:"pago_usuario"`
	// PagoPendiente  si se acepta saldo positivo
	ReferenciaFacturas []RefFactura `bson:"referencia_facturas" json:"referencia_facturas"`
}

// RefFactura agrupar valores para una factura
type RefFactura struct {
	EstaCancelado       bool               `bson:"esta_cancelado" json:"esta_cancelado"`
	PeriodoFactura      string             `bson:"periodo_factura" json:"periodo_factura"`
	IDReferenciaFactura primitive.ObjectID `bson:"id_referencia_factura" json:"id_referencia_factura"`
}

// ConstruirPorEvento crea un usuario usando un id
func (usr *Usuario) ConstruirPorEvento(evn EventPost) {
	usr.Referencia = primitive.NewObjectID()
	usr.ID = evn.UserID
	usr.CargoUsuario = 0.00
	usr.PagoUsuario = 0.00
	usr.ReferenciaFacturas = nil
}

// ConstruirPorPago crea un usuario usando un id
func (usr *Usuario) ConstruirPorPago(pag Pago) {
	usr.Referencia = primitive.NewObjectID()
	usr.ID = pag.IDUsuario
	usr.CargoUsuario = 0.00
	usr.PagoUsuario = 0.00
	usr.ReferenciaFacturas = nil
}

// AgregarFactura agrega una factura a usuario
func (usr *Usuario) AgregarFactura(fac Factura) {
	refFac := RefFactura{}
	refFac.EstaCancelado = false //siempre que se agrega una factura en false
	refFac.PeriodoFactura = fac.MesAnioFactura
	refFac.IDReferenciaFactura = fac.Referencia
	usr.ReferenciaFacturas = append(usr.ReferenciaFacturas, refFac)
}

// AgregarCambiosPorCargo agrega a la deuda del usuario el monto del cargo ingresado
func (usr *Usuario) AgregarSaldoPorCargo(car Cargo) {
	usr.CargoUsuario += car.MontoCargo
}

// BuscarUsuario busca el usuario en la base de datos si no la encuentro, contruyo con los datos
func (usr *Usuario) BuscarUsuario(ev EventPost) error {
	err := DataBaseCollection().FindOne(context.TODO(), bson.D{{"id", ev.UserID}}).Decode(&usr)
	if err != nil {
		usr.ConstruirPorEvento(ev)
	}
	return err
}

// BuscarUsuarioPorPago busca el usuario en la base de datos si no la encuentro, contruyo con los datos
func (usr *Usuario) BuscarUsuarioPorPago(pag Pago) error {
	err := DataBaseCollection().FindOne(context.TODO(), bson.D{{"id", pag.IDUsuario}}).Decode(&usr)
	if err != nil {
		usr.ConstruirPorPago(pag)
	}
	return err
}

// BuscarFactura busca la factura en memoria segun el vector de facturas, luego en BD segun _id
func (usr *Usuario) BuscarFactura(even EventPost) Factura {
	var fac Factura

	periodo := utils.ConvertirAnioMesString(even.Date)
	index := -1
	for i := range usr.ReferenciaFacturas {
		if usr.ReferenciaFacturas[i].PeriodoFactura == periodo {
			index = i // existe la factura en usuario
			break
		}
	}
	if index != -1 { // cargo de la BD
		DataBaseCollection().FindOne(context.TODO(), bson.D{{"_id", usr.ReferenciaFacturas[index].IDReferenciaFactura}}).Decode(&fac)
	} else {
		fac.Construir(even)     // si no esta en la BD entonces creo fac
		usr.AgregarFactura(fac) // agrego la referencia de fac en usuario
	}
	return fac
}

// Almacenar elimina el documento previo e inserta la ultima version usando el mismo _id
func (usr *Usuario) Almacenar() {
	//
	DataBaseCollection().DeleteOne(context.TODO(), bson.M{"_id": usr.Referencia})

	insertResult, err := DataBaseCollection().InsertOne(context.TODO(), usr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a single document Usuario: ", insertResult.InsertedID)
}

// BuscarFacturaNoCancelada TODO
func (usr *Usuario) BuscarFacturaNoCancelada() Factura {
	var fac Factura

	return fac
}

// PagarFacturas realiza pagos a las facturas
func (usr *Usuario) PagarFacturas(pago Pago) {
	var fac Factura
	for idx, refFac := range usr.ReferenciaFacturas {
		if refFac.EstaCancelado == false {
			fac.CargarDeReferencia(refFac.IDReferenciaFactura)
			fac.Pagar(&pago)
			if fac.CargoFactura == fac.Pagofactura { // si cumple es porque esta cancelada la factura
				usr.ReferenciaFacturas[idx].EstaCancelado = true
			}
			fac.AsociarPago(pago)
			fac.Almacenar()
		}
		// TODO:modificar en caso se acepten saldo positivo
		// agregarPagoPendiente
	}
	usr.PagoUsuario = usr.PagoUsuario + pago.MontoPago - pago.MontoPendiente

	pago.Almacenar()
}
