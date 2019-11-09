package config

import (
	"fmt"

	"github.com/eduardogpg/gonv"
)

// DatabaseConfig representa los parametros de conexion a la base de datos
type DatabaseConfig struct {
	username   string
	password   string
	host       string
	port       int
	database   string
	collection string
	debug      bool
}

var database *DatabaseConfig

func init() {
	// los datos de conexion de la base de datos desde Env
	database = &DatabaseConfig{}
	database.username = gonv.GetStringEnv("USERNAME", "root")
	database.password = gonv.GetStringEnv("PASSWORD", "")
	database.host = gonv.GetStringEnv("HOST", "localhost")
	database.port = gonv.GetIntEnv("PORT", 27017)
	database.database = gonv.GetStringEnv("DATABASE", "facturacion_go_restapi")
	database.collection = gonv.GetStringEnv("COLLECTION", "facturacion")
	database.debug = gonv.GetBoolEnv("DEBUG", true)

}

func (db *DatabaseConfig) url() string {
	return fmt.Sprintf("mongodb://%s:%d", db.host, db.port)
}

func (db *DatabaseConfig) collectionName() string {
	return db.collection
}

func (db *DatabaseConfig) databaseName() string {
	return db.database
}

// DatabaseName el nombre de la base de datos
func DatabaseName() string {
	return database.databaseName()
}

// CollectionName el nombre de la coleccion
func CollectionName() string {
	return database.collectionName()
}

// URLDatabase devuelve la cadena de conexion de la base de datos
func URLDatabase() string {
	return database.url()
}
