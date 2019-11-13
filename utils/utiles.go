package utils

import (
	"fmt"
	"strings"
)

var tipoDeCambio float64

func init() {
	// ver si podemos tomar de una api externa
	tipoDeCambio = 65.0 // pesos/dolar
}

// ConvertirAPesos convierte dolares a pesos por el tipo de cambio
func ConvertirAPesos(dol float64, curr string) float64 {
	if strings.ToUpper(curr) == "USD" {
		dol = dol * tipoDeCambio
	}
	//dol = math.Floor(dol*100) / 100 // el redondeo sera para emviar al usuario respuesta
	return dol

}

// ConvertirAnioMes convierte Time a "####-##" año-mes
//func ConvertirAnioMes(date time.Time) string {
//	return fmt.Sprintf("%d-%s", date.Year(), date.Month().String())
//}

// ConvertirAnioMesString convierte string a "####-##" año-mes
func ConvertirAnioMesString(date string) string {
	arr := strings.Split(date, "-")
	return fmt.Sprintf("%s-%s", arr[0], arr[1])
}
