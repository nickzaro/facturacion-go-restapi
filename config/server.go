package config

import (
	"fmt"

	"github.com/eduardogpg/gonv"
)

// ServerConfig parametros de servidor
type ServerConfig struct {
	host  string
	port  int
	debug bool
}

var server *ServerConfig

func init() {

	// los datos de servidor
	server = &ServerConfig{}
	server.host = gonv.GetStringEnv("HOST", "localhost")
	server.port = gonv.GetIntEnv("PORT", 3000)
}

func (srv *ServerConfig) url() string {
	return fmt.Sprintf("%s:%d", srv.host, srv.port)
}

// URLServer delvuelve la cadena de conexion del servidor
func URLServer() string {
	return server.url()
}

// PortSever devuelve el puerto del servidor
func PortSever() int {
	return server.port
}
