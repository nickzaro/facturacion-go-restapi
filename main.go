package main

import (
	"fmt"

	routes "github.com/nickzaro/facturacion-go-restapi/routers"
)

func main() {
	// routers.IniciarRutas()
	fmt.Printf("Iniciando el servidor")
	routes.IniciarServidor()

}
