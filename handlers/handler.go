package handlers

import (
	"fmt"
	"net/http"
)

// PostDeEvento maneja el post de los eventos
func PostDeEvento(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "PostDeEvento")
}

// PostDePagos maneja el post de pagos
func PostDePagos(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "PostDePagos")
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
