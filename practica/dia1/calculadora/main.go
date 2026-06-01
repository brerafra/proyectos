/*
*********************************************************************

practica día 1

1.2 Creación de una calculadora simple CLI con unicamente las operaciones
+,-,*,/ integradas.

*********************************************************************
*/

package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a, b string
	var operacion string

	fmt.Println("Programa calculadora sencilla")

	fmt.Print("Ingresa operación: ")
	fmt.Scan(&operacion)
	if operacion != "+" {
		if operacion != "-" {
			if operacion != "*" {
				if operacion != "/" {
					fmt.Println("Operacion invalida")
					return
				}
			}
		}
	}

	fmt.Print("Ingresa primer factor: ")
	fmt.Scan(&a)
	fa, err := strconv.ParseFloat(a, 64)
	if err != nil {
		fmt.Println("Primer factor invalido")
		return
	}
	fmt.Print("Ingresa el segundo factor:")
	fmt.Scan(&b)

	fb, err := strconv.ParseFloat(b, 64)
	if err != nil {
		fmt.Println("Segundo factor invalido")
		return
	}
	var resultado float64

	if operacion == "+" {
		resultado = fa + fb
	}

	if operacion == "-" {
		resultado = fa - fb
	}

	if operacion == "*" {
		resultado = fa * fb
	}

	if operacion == "/" {
		if fb == 0 {
			fmt.Println("No se puede dividir entre cero")
			return
		}
		resultado = fa / fb
	}

	fmt.Printf("Resultado: %v\n", resultado)

}
