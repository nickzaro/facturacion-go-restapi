package models

import (
	"context"
	"fmt"
	"log"

	"github.com/nickzaro/facturacion-go-restapi/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Cargo estructura de cargos [un cargo varios pagos asociados]
type Cargo struct {
	Referencia     primitive.ObjectID   `bson:"_id" json:"_id"`
	IDUsuario      int                  `bson:"id_usuario" json:"id_usuario"`
	MesAnioFactura string               `bson:"mesanio_factura" json:"mesanio_factura"`
	FechaCargo     string               `bson:"fecha_cargo" json:"fecha_cargo"`
	Categoria      string               `bson:"categoria" json:"categoria"`
	Subcategoria   string               `bson:"subcategoria" json:"subcategoria"`
	MontoCargo     float64              `bson:"monto_cargo" json:"monto_cargo"`
	MontoPendiente float64              `bson:"monto_pendiente" json:"monto_pendiente"`
	EventoID       int                  `bson:"evento_id" json:"evento_id"`
	PagosAsociados []primitive.ObjectID `bson:"pagos_asociados" json:"pagos_asociados"`
}

// EventPost estructura de eventos solo json no se almacena en la BD
type EventPost struct {
	EventID   int     `json:"event_id"`
	Amount    float64 `json:"amount"`
	Currency  string  `json:"currency"`
	UserID    int     `json:"user_id"`
	EventType string  `json:"event_type"`
	Date      string  `json:"date"`
}

var categorias map[string]string

func init() {
	categorias = map[string]string{
		"CLASIFICADO": "MARKETPLACE",
		"VENTA":       "MARKETPLACE",
		"ENVÍO":       "MARKETPLACE",
		"PUBLICIDAD":  "SERVICIOS",
		"FIDELIDAD":   "SERVICIOS",
		"CRÉDITO":     "SERVICIOS",
		"MERCADOPAGO": "EXTERNO",
		"MERCADOSHOP": "EXTERNO",
	}
}

// Construir cargo con los datos de EventPost
func (car *Cargo) Construir(eventPost EventPost) {
	car.Referencia = primitive.NewObjectID()
	car.IDUsuario = eventPost.UserID
	car.MesAnioFactura = utils.ConvertirAnioMesString(eventPost.Date)
	car.FechaCargo = eventPost.Date
	car.Categoria = categorias[eventPost.EventType]
	car.Subcategoria = eventPost.EventType
	car.MontoCargo = utils.ConvertirAPesos(eventPost.Amount, eventPost.Currency)
	car.MontoPendiente = car.MontoCargo // al inicio el monto pendiente es el total
	car.EventoID = eventPost.EventID
	car.PagosAsociados = nil
}

// Almacenar elimina el documento previo e inserta la ultima version usando el mismo _id
func (car *Cargo) Almacenar() {
	//
	DataBaseCollection().DeleteOne(context.TODO(), bson.M{"_id": car.Referencia})

	insertResult, err := DataBaseCollection().InsertOne(context.TODO(), car)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a single document Cargo: ", insertResult.InsertedID)
}

// Pagar realiza un pago parcial o total usando un pago
func (car *Cargo) Pagar(pago *Pago) float64 {

	montoPagado := 0.00
	if car.MontoPendiente == 0 {
		return montoPagado
	}
	if car.MontoPendiente > pago.MontoPendiente {
		montoPagado = car.MontoPendiente - pago.MontoPendiente
		car.MontoPendiente = montoPagado
		pago.MontoPendiente = 0
	} else if car.MontoPendiente < pago.MontoPendiente {
		montoPagado = pago.MontoPendiente - car.MontoPendiente
		pago.MontoPendiente = montoPagado
		car.MontoPendiente = 0
	} else {
		montoPagado = pago.MontoPendiente
		car.MontoPendiente = 0
		pago.MontoPendiente = 0
	}
	fmt.Println("PAGANDO CARGO CON EL MONTO INGRESADO")
	car.AsociarPago(pago)
	pago.AsociarCargo(*car) //TODO ver si hay problemas con el puntero
	fmt.Println(pago)
	return montoPagado // el monto que paga el pago
}

func (car *Cargo) AsociarPago(pago *Pago) {
	car.PagosAsociados = append(car.PagosAsociados, pago.Referencia)
}

// CargarDeReferencia cargar un Cargo usando el _id
func (car *Cargo) CargarDeReferencia(refFac primitive.ObjectID) {
	DataBaseCollection().FindOne(context.TODO(), bson.D{{"_id", refFac}}).Decode(&car)
}
