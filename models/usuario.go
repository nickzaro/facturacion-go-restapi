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
	Referencia      primitive.ObjectID   `bson:"_id" json:"_id"`
	ID              int                  `bson:"id" json:"id"`
	CargoUsuario    float64              `bson:"cargo_usuario" json:"cargo_usuario"`
	PagoUsuario     float64              `bson:"pago_usuario" json:"pago_usuario"`
	PeriodoFacturas []string             `bson:"periodo_facturas" json:"periodo_facturas"`
	Facturas        []primitive.ObjectID `bson:"facturas" json:"facturas"`
}

// Construir crea un usuario usando un id
func (usr *Usuario) Construir(evn EventPost) {
	usr.Referencia = primitive.NewObjectID()
	usr.ID = evn.UserID
	usr.CargoUsuario = 0.00
	usr.PagoUsuario = 0.00
	usr.PeriodoFacturas = nil
	usr.Facturas = nil
}

// AgregarFactura agrega una factura a usuario
func (usr *Usuario) AgregarFactura(fac Factura) {
	usr.PeriodoFacturas = append(usr.PeriodoFacturas, fac.MesAnioFactura)
	usr.Facturas = append(usr.Facturas, fac.Referencia)
}

// AgregarSaldoPorCargo agrega a la deuda del usuario el monto del cargo ingresado
func (usr *Usuario) AgregarSaldoPorCargo(car Cargo) {
	usr.CargoUsuario += car.MontoCargo
}

// BuscarUsuario busca el usuario en la base de datos si no la encuentro, contruyo con los datos
func (usr *Usuario) BuscarUsuario(ev EventPost) {
	err := DataBaseCollection().FindOne(context.TODO(), bson.D{{"id", ev.UserID}}).Decode(&usr)
	if err != nil {
		usr.Construir(ev)
	}
}

// BuscarFactura busca la factura en memoria segun el vector de facturas, luego en BD segun _id
func (usr *Usuario) BuscarFactura(even EventPost) Factura {
	var fac Factura

	periodo := utils.ConvertirAnioMesString(even.Date)
	index := -1
	for i := range usr.PeriodoFacturas {
		if usr.PeriodoFacturas[i] == periodo {
			index = i // existe la factura en usuario
			break
		}
	}
	if index != -1 { // cargo de la BD
		DataBaseCollection().FindOne(context.TODO(), bson.D{{"_id", usr.Facturas[index]}}).Decode(&fac)
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
