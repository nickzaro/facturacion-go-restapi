package main

import (
	"fmt"

	controllers "github.com/nickzaro/facturacion-go-restapi/controllers"
	//routes "github.com/nickzaro/facturacion-go-restapi/routers"
)

func main() {

	controllers.Xmain()
	fmt.Printf("Iniciando el servidor")
	//routes.IniciarServidor()

}
