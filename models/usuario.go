package models

// Usuario estructura de los documentos en la BAse de datos
type Usuario struct {
	ID           string  `json:"id"`
	DeudaUsuario float32 `json:"deudausuario"`
	PagoUsuario  float32 `json:"pagousuario"`
	Facturas     []Factura `json:"facturas"`
}

/**
{
    "id": "1",
    "deudausuario": 1000,
    "pagousuario": 900,
    "facturas": [
        {
            "factura-fecha": "2019-06",
            "deudafactura": 200,
            "pagofactura": 100,
            "cargos": [
                {
                    "fecha": "2019-06-10-12:32",
                    "categoria": "MARKETPLACE",
                    "subcategoria": "CLASIFICADO",
                    "monto": 40
                }
            ],
            "pagos": [
                {
                    "fecha": "2019-06-10-14:35",
                    "monto": 20
                }
            ]
        }
    ]
}
*/
