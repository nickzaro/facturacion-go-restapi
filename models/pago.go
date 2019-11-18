package models

import (
	"context"
	"fmt"
	"log"

	"github.com/nickzaro/facturacion-go-restapi/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Pago estructura de pagos [un pago varios cargos asociados]
type Pago struct {
	Referencia      primitive.ObjectID   `bson:"_id" json:"_id"`
	IDUsuario       int                  `bson:"id_usuario" json:"id_usuario"`
	MontoPago       float64              `bson:"monto_pago" json:"monto_pago"`
	MontoPendiente  float64              `bson:"monto_pendiente" json:"monto_pendiente"`
	CargosAsociados []primitive.ObjectID `bson:"cargos_asociados" json:"cargos_asociados"`
}

// PagoPost estructura que representa el json que recibimos a travez de
// la peticion post de pago
type PagoPost struct {
	Amount   float64 `json:"amount"`
	Currency string  `json:"currency"`
	UserID   int     `json:"user_id"`
}

// Construir crea un usuario usando un id
func (pag *Pago) Construir(p PagoPost) {
	pag.Referencia = primitive.NewObjectID()
	pag.IDUsuario = p.UserID
	pag.MontoPago = utils.ConvertirAPesos(p.Amount, p.Currency)
	pag.MontoPendiente = pag.MontoPago // a l inicio el monto pendiente es el total del monto
	pag.CargosAsociados = nil
	//MesAnioFactura no puse nada, al asociar a factura será
}

func (pag *Pago) AsociarCargo(cargo Cargo) {
	pag.CargosAsociados = append(pag.CargosAsociados, cargo.Referencia)
}

func (pag *Pago) Almacenar() {
	DataBaseCollection().DeleteOne(context.TODO(), bson.M{"_id": pag.Referencia})

	insertResult, err := DataBaseCollection().InsertOne(context.TODO(), pag)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a single document Pago: ", insertResult.InsertedID)
}
