package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nickzaro/facturacion-go-restapi/config"
	"github.com/nickzaro/facturacion-go-restapi/handlers"
)

var router *mux.Router

// IniciarRutas inicia las routers para atencion de peticiones
func iniciarRutas() *mux.Router {

	mux := mux.NewRouter()
	mux.HandleFunc("/api/evento/", handlers.PostDeEvento).Methods("POST")
	mux.HandleFunc("/api/pago/", handlers.PostDePago).Methods("POST")
	mux.HandleFunc("/api/cargos/{id:[0-9]+}", handlers.GetDeCargos).Methods("GET")
	mux.HandleFunc("/api/facturas/{id:[0-9]+}", handlers.GetDeFacturas).Methods("GET")
	mux.HandleFunc("/api/pagos/{id:[0-9]+}", handlers.GetDePagos).Methods("GET")
	mux.HandleFunc("/api/usuario/{id:[0-9]+}", handlers.GetStatusUsuario).Methods("GET")
	return mux
}

// Router devuelve el router de la conexion
func Router() *mux.Router {
	return router
}

//IniciarServidor carga las rutas e inicia el servidor
func IniciarServidor() {
	log.Fatal(http.ListenAndServe(config.URLServer(), iniciarRutas()))
}
