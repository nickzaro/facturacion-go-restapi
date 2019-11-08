package config

import (
	"fmt"

	"github.com/eduardogpg/gonv"
)

// DatabaseConfig representa los parametros de conexion a la base de datos
type DatabaseConfig struct {
	username string
	password string
	host     string
	port     int
	database string
	debug    bool
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
	database.debug = gonv.GetBoolEnv("DEBUG", true)

}

func (db *DatabaseConfig) url() string {
	return fmt.Sprintf("mongodb://%s:%d", db.host, db.port)
}

// URLDatabase devuelve la cadena de conexion de la base de datos
func URLDatabase() string {
	return database.url()
}
