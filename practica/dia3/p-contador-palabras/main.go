/*
*********************************************************************

practica día 3

2.1 Contador de palabras usando slices

*********************************************************************
*/

package main

import "fmt"

func main() {
	frase := []string{"Hola", "mundo", "como", "estas", " ", "adios", "."}

	fmt.Printf("Palabras dentro del slice: %v\n", len(frase))

	i := 0
	for _, k := range frase {
		i++
		fmt.Println(k)
	}
	fmt.Printf("Palagras dentro del slice: %d\n", i)

}
