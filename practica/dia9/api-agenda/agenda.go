package main

type Agenda struct {
	ID        int64  `json:"id,omitempty"`
	Nombre    string `json:"nombre"`
	Numero    string `json:"numero"`
	Direccion string `json:"direccion"`
	Edad      int    `json:"edad"`
}
