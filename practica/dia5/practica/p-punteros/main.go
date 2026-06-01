/*
*********************************************************************

practica día 5

1.1 Practica con punteros

-> Los punteros son variables que almacenan la dirección de memoria de otro valor
permiten compartir datos entre funciones de forma eficiente y modificar variables
directamente desde otras funciones. se utiliza & para obtener la dirección y
* para acceder al valor almacenado.

*********************************************************************
*/
package main

import "fmt"

// funcion recibe un puntero a un tenero (*int)
func duplicar(valor *int) {
	//se usa el asterisco para desreferenciar y modificar el valor original
	*valor = *valor * 2
}

// intercambiar valor de dos variables (swapping)
func intercambiar(x, y *int) {
	temp := *x
	*x = *y
	*y = temp
}

//Cuando usas estructuras grandes pasarlas completas como valor a una funcion genera
//una copia de todos sus datos. Usar un puntero a struct es mas eficiente por que solo
//copia la dirección de memoria

type Rectangulo struct {
	Ancho float64
	Alto  float64
}

func (r *Rectangulo) Escalar(factor float64) {
	r.Ancho = r.Ancho * factor
	r.Alto = r.Alto * factor
}

func main() {
	numero := 10
	fmt.Println("Valor original:", numero)

	//pasamos la dirección de memoria de "numero" usando el simbolo "&"
	duplicar(&numero)

	fmt.Println("Valor modificado: ", numero)

	//----- Intercambiando valores
	a, b := 5, 10
	fmt.Printf("Antes: a = %d, b = %d\n", a, b)

	intercambiar(&a, &b)

	fmt.Printf("Despues a= %d, b =%d\n", a, b)

	//punteros en structs

	rect := Rectangulo{Ancho: 3, Alto: 4}
	fmt.Printf("Area inicial: %.2f\n", rect.Ancho*rect.Alto)

	//modificamos el rectangulo pasandolo por referencia
	rect.Escalar(2.0)
	fmt.Printf("Area modificada: %.2f\n", rect.Ancho*rect.Alto)
}
