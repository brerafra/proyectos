/*
*********************************************************************

# Proyecto día 5

5.2 API Crear un mini banco cli

# En esta aplicación interactiva, se podra añadir saldo, ver saldo y quitar saldo
a un usuario en especifico

********************************************************************
*/
package main

import (
	"fmt"
	"strconv"
)

type User struct {
	Id     int64   `json:"id"`
	Name   string  `json:"name"`
	Email  string  `json:"email"`
	Status bool    `json:"status"`
	Saldo  float64 `json:"saldo"`
}

func main() {
	var funcion string
	var cantidad string
	var Id string
	Users := []User{
		{Id: 1, Name: "Brenthon", Email: "brerafra@gmail.com", Status: true, Saldo: 0},
		{Id: 2, Name: "Manuel", Email: "correo1@gmail.com", Status: true, Saldo: 100},
		{Id: 3, Name: "Israel", Email: "correo2@gmail.com", Status: true, Saldo: 500},
		{Id: 4, Name: "Irma", Email: "correo3@gmail.com", Status: true, Saldo: 0},
		{Id: 5, Name: "Adriana", Email: "correo4@gmail.com", Status: true, Saldo: 359.59},
	}

	fmt.Printf("Programa CLI mini Banco\n")
	for {
		fmt.Printf("Ingresa una opción: 1.Usuarios, 2.Añadir, 3.Quitar:")
		fmt.Scan(&funcion)
		fmt.Println()

		switch funcion {
		case "1":
			fmt.Printf("Usuarios: \n")
			for _, k := range Users {
				if k.Status == true {
					fmt.Printf("id: %d, nombre: %s, email: %s, saldo: %.2f\n", k.Id, k.Name, k.Name, k.Saldo)
				}
			}
		case "2":
			fmt.Println("Agregar saldo:")
			fmt.Printf("Ingrese Id usuario:")
			fmt.Scan(&Id)
			fmt.Println()
			user := User{}
			id_int, err := strconv.ParseInt(Id, 10, 64)
			if err != nil {
				fmt.Println("Error:", err)
				break

			} else {
				for _, k := range Users {
					if k.Id == id_int {
						user = k
					}
				}
			}
			fmt.Printf("Usuario:%s, Saldo: %.2f\n", user.Name, user.Saldo)
			fmt.Printf("Ingrese cantidad:")
			fmt.Scan(&cantidad)
			cantidad_float, err := strconv.ParseFloat(cantidad, 64)
			if err != nil {
				fmt.Println("Error:", err)
				break
			}

			fmt.Printf("Cantidad: %.2f\n", cantidad_float)

			saldo_final := user.Saldo + cantidad_float

			fmt.Printf("Saldo final: %.2f\n", saldo_final)

			Users[user.Id-1].Saldo = saldo_final

		case "3":
			fmt.Println("Agregar saldo:")
			fmt.Printf("Ingrese Id usuario:")
			fmt.Scan(&Id)
			fmt.Println()
			user := User{}
			id_int, err := strconv.ParseInt(Id, 10, 64)
			if err != nil {
				fmt.Println("Error:", err)
				break

			} else {
				for _, k := range Users {
					if k.Id == id_int {
						user = k
					}
				}
			}
			fmt.Printf("Usuario:%s, Saldo: %.2f\n", user.Name, user.Saldo)
			fmt.Printf("Ingrese cantidad:")
			fmt.Scan(&cantidad)
			cantidad_float, err := strconv.ParseFloat(cantidad, 64)
			if err != nil {
				fmt.Println("Error:", err)
				break
			}

			fmt.Printf("Cantidad: %.2f\n", cantidad_float)

			if (user.Saldo - cantidad_float) < 0 {
				fmt.Println("No existe dinero suficiente para retirar")
				break
			}

			saldo_final := user.Saldo - cantidad_float

			fmt.Printf("Saldo final: %.2f\n", saldo_final)

			Users[user.Id-1].Saldo = saldo_final

		default:
			fmt.Println("operacion no valida")
		}
	}
}
