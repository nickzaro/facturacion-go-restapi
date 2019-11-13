package models

import (
	"context"
	"fmt"
	"log"

	"github.com/nickzaro/facturacion-go-restapi/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Trainer estrctura de prueba
type Trainer struct {
	Name string
	Age  int
	City string
}

var connection *mongo.Client
var collection *mongo.Collection

func init() {
	// Set client options
	clientOption := options.Client().ApplyURI(config.URLDatabase())
	// Connect to MongoDB
	con, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}
	connection = con

	DataBasePing()
	fmt.Println("Connected to MongoDB!")
	collection = DataBaseConnection().Database(config.DatabaseName()).Collection(config.CollectionName())
}

// DataBaseConnection devuelve la coneccion
func DataBaseConnection() *mongo.Client {
	return connection
}

// DataBaseCollection devuelve la coleccion
func DataBaseCollection() *mongo.Collection {
	return collection
}

// DataBaseDisconnect cierra la connexion de la Base de datos
func DataBaseDisconnect() {
	// desconectar el cliente
	err := DataBaseConnection().Disconnect(context.TODO())

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection to MongoDB closed.")
}

// DataBasePing ping a la Base de datos
func DataBasePing() {
	// Check the connection
	err := DataBaseConnection().Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}
}

// Xmain funciones sobre una base de datos
func Xmain() {

	//CREATING CLIENTS
	ash := Trainer{"Ash", 10, "Pallet Town"}
	misty := Trainer{"Misty", 10, "Cerulean City"}
	brock := Trainer{"Brock", 15, "Pewter City"}

	// insert un objeto
	insertResult, err := collection.InsertOne(context.TODO(), ash)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a single document: ", insertResult.InsertedID)

	//insert´s various items
	trainers := []interface{}{misty, brock}

	insertManyResult, err := collection.InsertMany(context.TODO(), trainers)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted multiple documents: ", insertManyResult.InsertedIDs)

	// Update documents
	filter := bson.D{{"name", "Ash"}}

	update := bson.D{
		{"$inc", bson.D{
			{"age", 1},
		}},
	}

	// updateResult, err := collection.UpdateMany(context.TODO(), filter, update)
	updateResult, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)

	// Find documents

	var result Trainer

	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Found a single document: %+v\n", result)

	// find multiple documents
	// Pass these options to the Find method
	findOptions := options.Find()
	findOptions.SetLimit(2)

	// Here's an array in which you can store the decoded documents
	var results []*Trainer

	// Passing bson.D{{}} as the filter matches all documents in the collection
	cur, err := collection.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Finding multiple documents returns a cursor
	// Iterating through the cursor allows us to decode documents one at a time
	for cur.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var elem Trainer
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, &elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	// Close the cursor once finished
	cur.Close(context.TODO())

	fmt.Printf("Found multiple documents (array of pointers): %+v\n", results) // ver el tema de los valores y no punteros

	// Delete all documments
	deleteResult, err := collection.DeleteMany(context.TODO(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Deleted %v documents in the trainers collection\n", deleteResult.DeletedCount)
	DataBaseDisconnect()
}
