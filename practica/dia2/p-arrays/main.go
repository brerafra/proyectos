/*
*********************************************************************

practica día 2

2.1 practica general de arrays y slices.

*********************************************************************
*/

package main

import "fmt"

func main() {
	//Tamaño inferido
	numeros := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	//Tamaño explicito
	var vacio [5]int

	//Inicialización parcial o con índices especificos
	arr := [...]int{1: 10, 3: 30}

	fmt.Println(numeros, vacio, arr)

	//Iteracion o loops

	postres := []string{"Gelatina", "Pie", "Torta"}

	//Metodo 1: Range
	for index, postre := range postres {
		fmt.Printf("Indice %d: %s \n", index, postre)
	}

	//Metodo 2: indices standar

	for i := 0; i < len(postres); i++ {
		fmt.Println(postres[i])
	}

	//Operaciones Comunes
	//Encontrar mínimo y máximo

	min := numeros[0]
	max := numeros[0]

	for _, num := range numeros {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}

	fmt.Printf("Minimo: %d, Maximo %d \n", min, max)

	//Suma y promedio
	suma := 0
	for _, num := range numeros {
		suma += num
	}

	fmt.Printf("Promedio: %.2f\n", float64(suma)/float64(len(numeros)))

	//Comparación

	arr1 := [...]int{1, 2, 3}
	arr2 := [...]int{1, 2, 3}
	arr3 := [...]int{1, 2, 4}

	fmt.Println(arr1 == arr2)
	fmt.Println(arr1 == arr3)
}
