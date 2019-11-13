package models

import (
	"github.com/nickzaro/facturacion-go-restapi/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Pago estructura de pagos [un pago varios cargos asociados]
type Pago struct {
	Referencia      primitive.ObjectID   `bson:"_id" json:"_id"`
	IDUsuario       int                  `bson:"id_usuario" json:"id_usuario"`
	MesAnioFactura  string               `bson:"mesanio_factura" json:"mesanio_factura"`
	MontoPago       float64              `bson:"monto_pago" json:"monto_pago"`
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
	pag.CargosAsociados = nil
	//MesAnioFactura no puse nada, al asociar a factura será
}
