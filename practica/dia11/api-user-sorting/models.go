package main

type Producto struct {
	ID     int64  `json:"id"`
	Nombre string `json:"nombre"`
	Precio int64  `json:"precio"`
	Status bool   `json:"status"`
}
