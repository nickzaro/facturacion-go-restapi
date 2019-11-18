package controllers

import (
	"context"
	"testing"

	"github.com/nickzaro/facturacion-go-restapi/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/mgo.v2/bson"
)

/*
func TestXmain(t *testing.T) {
	controllers.Xmain()
}
*/

/*
func TestBuscarUsuario(t *testing.T) {
	even := models.EventPost{1, 100.00, "USD", 1, "CLASIFICADO", "2019-05-16T00:00:00"}
	var usr models.Usuario
	usr.Construir(even)
	t.Errorf("De referencia %s  %d", usr.Referencia, usr.ID)

}

func TestProcesarEventPostv1(t *testing.T) {
	even := models.EventPost{1, 100.00, "USD", 1, "CLASIFICADO", "2019-05-16T00:00:00"}
	controllers.ProcesarEventPostv1(even)
}
*/
var sub1 Sub
var sub2 Sub
var estudiante Estructura

func init() {

}

type Sub struct {
	NombreMateria string `bson:"nombre_materia" json:"nombre_materia" `
	CodigoMateria int    `bson:"codigo_materia" json:"codigo_materia"`
}
type Estructura struct {
	ID       primitive.ObjectID `bson:"_id" json:"_id"`
	Nombre   string             `bson:"nombre" json:"nombre"`
	Materias []Sub              `bson:"materias" json:"materias"`
}

func TestPruebas(t *testing.T) {
	sub1 = Sub{"Fisica", 71}
	sub2 = Sub{"Quimica", 72}
	estudiante.ID = primitive.NewObjectID()
	estudiante.Nombre = "JonasX"
	estudiante.Materias = nil

	estudiante.Materias = append(estudiante.Materias, sub1)
	estudiante.Materias = append(estudiante.Materias, sub2)
	res, err := models.DataBaseCollection().InsertOne(context.TODO(), estudiante)
	if err != nil {
		t.Fatal(estudiante)
	}
	est := Estructura{}
	models.DataBaseCollection().FindOne(context.TODO(), bson.M{"_id": res.InsertedID}).Decode(&est)
	t.Fatal(est)

}
