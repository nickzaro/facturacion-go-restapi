package main

import (
	"fmt"

	routes "github.com/nickzaro/facturacion-go-restapi/routers"
)

func main() {

	//controllers.Xmain()
	fmt.Printf("Iniciando el servidor\n")
	routes.IniciarServidor()

}
