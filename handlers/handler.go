package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/nickzaro/facturacion-go-restapi/controllers"
	"github.com/nickzaro/facturacion-go-restapi/models"
)

// PostDeEvento maneja el post de los eventos
func PostDeEvento(w http.ResponseWriter, r *http.Request) {
	// recibimos el json
	var eventpost models.EventPost
	_ = json.NewDecoder(r.Body).Decode(&eventpost)
	// se procesaria la informacion y devuelve
	controllers.ProcesarEventPost(eventpost)
	json.NewEncoder(w).Encode(eventpost)
	//Aca procesar

	fmt.Fprintf(w, "PostDeEvento")
}

// PostDePago maneja el post de pagos
func PostDePago(w http.ResponseWriter, r *http.Request) {

	// recibimos el json
	var pagopost models.PagoPost
	_ = json.NewDecoder(r.Body).Decode(&pagopost)

	// se procesaria la informacion y devuelve
	//otroPagopost, _ := controllers.ProcesarPagoPost(pagopost)
	json.NewEncoder(w).Encode(pagopost)
	//Aca procesar

	//fmt.Fprintf(w, "PostDePagos")
}

// GetDeCargos maneja el get de cargos
func GetDeCargos(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "GetDeCargos")
}

// GetDeFacturas maneja el get de facturas
func GetDeFacturas(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "GetDeFacturas")
}

// GetDePagos maneja el get de pagos
func GetDePagos(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "GetDePagos")
}

// GetStatusUsuario manejo el get del status del usuario
func GetStatusUsuario(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "GetStatusUsuario")
}
