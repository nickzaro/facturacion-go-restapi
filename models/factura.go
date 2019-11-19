package models

import (
	"context"
	"fmt"
	"log"

	"github.com/nickzaro/facturacion-go-restapi/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Factura estructura de una factura[son factura por mes]
// los indices son las posiciones de las direcciones de la ultima cancelacion
// en la version sin saldo a favor pagos indice apuntaria
type Factura struct {
	Referencia       primitive.ObjectID `bson:"_id" json:"_id"`
	IDUsuario        int                `bson:"id_usuario" json:"id_usuario"`
	MesAnioFactura   string             `bson:"mesanio_factura" json:"mesanio_factura"`
	CargoFactura     float64            `bson:"cargo_factura" json:"cargo_factura"`
	Pagofactura      float64            `bson:"pago_factura" json:"pago_factura"`
	ReferenciaCargos []ItemFacturaCargo `bson:"referencia_cargos" json:"referencia_cargos"`
	ReferenciaPagos  []ItemFacturaPago  `bson:"referencia_pagos" json:"referencia_pagos"`
}

// Separados por si tengo que poner mas atributos en alguna estructura

// ItemFacturaCargo para representar la referencia y el estado cancelado o no
type ItemFacturaCargo struct {
	EstaCancelado     bool               `bson:"esta_cancelado" json:"esta_cancelado"`
	IDReferenciaCargo primitive.ObjectID `bson:"id_referencia_cargo" json:"id_referncia_cargo"`
}

// ItemFacturaPago para representar la referencia y el estado cancelado o no
type ItemFacturaPago struct {
	EstaCancelado    bool               `bson:"esta_cancelado" json:"esta_cancelado"`
	IDReferenciaPago primitive.ObjectID `bson:"id_referencia_pago" json:"id_referncia_pago"`
}

// Construir construye la eestructura usando informacion de EventPost
func (fac *Factura) Construir(evn EventPost) {
	fac.Referencia = primitive.NewObjectID()
	fac.IDUsuario = evn.UserID
	fac.MesAnioFactura = utils.ConvertirAnioMesString(evn.Date)
	fac.CargoFactura = 0.00
	fac.Pagofactura = 0.00
	fac.ReferenciaCargos = nil
	fac.ReferenciaPagos = nil

}

// AgregarCargo un NUEVO cargo a la factura
func (fac *Factura) AgregarCargo(car Cargo) {
	refcar := ItemFacturaCargo{}
	refcar.EstaCancelado = false
	refcar.IDReferenciaCargo = car.Referencia

	fac.ReferenciaCargos = append(fac.ReferenciaCargos, refcar)
	fac.CargoFactura += car.MontoCargo

}

// Almacenar borra la version anterior y guarda la version actual
func (fac *Factura) Almacenar() {
	//
	DataBaseCollection().DeleteOne(context.TODO(), bson.M{"_id": fac.Referencia})

	insertResult, err := DataBaseCollection().InsertOne(context.TODO(), fac)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a single document Factura: ", insertResult.InsertedID)
}

// Pagar pago
func (fac *Factura) Pagar(pago *Pago) float64 {
	var cargo Cargo
	montoPagado := 0.00
	for idx, refCargo := range fac.ReferenciaCargos {
		if refCargo.EstaCancelado == false {
			cargo.CargarDeReferencia(refCargo.IDReferenciaCargo)
			montoPagado = montoPagado + cargo.Pagar(pago)
			cargo.Almacenar()
			fmt.Printf("ESTO AGREGA A FAC.PAGOFACTURA %f\n", montoPagado)
			if cargo.MontoPendiente == 0 {
				fac.ReferenciaCargos[idx].EstaCancelado = true
			}

		}
		if pago.MontoPendiente == 0 {
			break
		}
	}
	fac.Pagofactura = fac.Pagofactura + montoPagado

	return montoPagado
}

// CargarDeReferencia cago una factura usando el _id
func (fac *Factura) CargarDeReferencia(refFac primitive.ObjectID) {
	DataBaseCollection().FindOne(context.TODO(), bson.D{{"_id", refFac}}).Decode(&fac)
}

// AsociarPago asocia un pago a una factura
func (fac *Factura) AsociarPago(pago Pago) {
	refpago := ItemFacturaPago{}
	if pago.MontoPendiente == 0 {
		refpago.EstaCancelado = true
	} else {
		refpago.EstaCancelado = false // deberia se false sin esto
	}
	refpago.IDReferenciaPago = pago.Referencia
	fac.ReferenciaPagos = append(fac.ReferenciaPagos, refpago)
}

//TODO falta asociar cargo para facturas que no son nuevas
