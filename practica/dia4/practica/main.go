/*
*********************************************************************

practica día 4

4.1 Aplicación CLI para control de sistema de usuarios (agregar, ver y borrar)

*********************************************************************
*/

package main

import (
	"fmt"
	"strconv"
)

type User struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Status   bool   `json:"status"`
}

var Users []User

func main() {
	var operacion string
	var user User
	var id string
	fmt.Println("Programa de gestion de usuarios:")
	for {
		fmt.Printf("Ingresa una opción: 1.Agregar, 2.Borrar. 3.Listar: ")
		fmt.Scan(&operacion)
		fmt.Println()

		switch operacion {
		case "1":
			fmt.Printf("Ingresa nombre: ")
			fmt.Scan(&user.Name)
			fmt.Printf("Ingresa email: ")
			fmt.Scan(&user.Email)
			fmt.Printf("Ingresa Password: ")
			fmt.Scan(&user.Password)

			if len(Users) == 0 {
				user.Id = 0
			} else {
				user.Id = Users[len(Users)-1].Id + 1
			}
			user.Status = true

			Users = append(Users, user)
			fmt.Println("usuario agregado correctamente")
		case "2":
			fmt.Printf("Da un id de usuario a eliminar: ")
			fmt.Scan(&id)

			id_int, err := strconv.ParseInt(id, 10, 64)
			if err != nil {
				fmt.Println("Error:", err)

			} else {
				for _, k := range Users {
					if k.Id == id_int {
						Users[k.Id].Status = false
					}
				}
			}
		case "3":
			fmt.Println("Usuarios agregados: ")
			for _, k := range Users {
				if k.Status == true {
					fmt.Printf("id: %d, nombre: %s, email: %s, password: %s\n", k.Id, k.Name, k.Email, k.Password)
				}
			}

		default:
			fmt.Println("operacion no valida.")
		}
	}

}
