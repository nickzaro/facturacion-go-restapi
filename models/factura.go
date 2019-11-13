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
	Referencia         primitive.ObjectID   `bson:"_id" json:"_id"`
	IDUsuario          int                  `bson:"id_usuario" json:"id_usuario"`
	MesAnioFactura     string               `bson:"mesanio_factura" json:"mesanio_factura"`
	CargoFactura       float64              `bson:"cargo_factura" json:"cargo_factura"`
	Pagofactura        float64              `bson:"pago_factura" json:"pago_factura"`
	IndiceCancelCargos int                  `bson:"indice_cancel_cargos" json:"indice_cancel_cargos"`
	Cargos             []primitive.ObjectID `bson:"cargos" json:"cargos"`
	IndiceCancelPagos  int                  `bson:"indice_cancel_pagos" json:"indice_cancel_pagos"`
	Pagos              []primitive.ObjectID `bson:"pagos" json:"pagos"`
}

// Construir construye la eestructura usando informacion de EventPost
func (fac *Factura) Construir(evn EventPost) {
	fac.Referencia = primitive.NewObjectID()
	fac.IDUsuario = evn.UserID
	fac.MesAnioFactura = utils.ConvertirAnioMesString(evn.Date)
	fac.CargoFactura = 0.00
	fac.Pagofactura = 0.00
	fac.IndiceCancelCargos = 0
	fac.IndiceCancelPagos = 0
	fac.Cargos = nil
	fac.Pagos = nil

}

// AgregarCargo agrega la referencia y el monto del cargo a factura
func (fac *Factura) AgregarCargo(car Cargo) {
	fac.Cargos = append(fac.Cargos, car.Referencia)
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
