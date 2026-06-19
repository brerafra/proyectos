package domain

type Empleado struct {
	Id           int64  `json:"id"`
	Nombre       string `json:"nombre"`
	Departamento string `json:"departamento"`
	Salario      int    `json:"salario"`
}
