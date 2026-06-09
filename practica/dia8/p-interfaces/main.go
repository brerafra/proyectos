/*
*********************************************************************

Practica día 8

8.1 Trabajo con interfaces

Sección 1: poliformismo.

Tenemos una interfaz Figura, con método Area(). cualquier estructura (Como
rectangulo o circulo) que calcule su propia área cumple con el contrato
*********************************************************************/

package main

import (
	"fmt"
	"math"
)

// Definimos la interfaz
type figura interface {
	Area() float64
}

// Estructura rectangulo
type Rectangulo struct {
	Ancho, Alto float64
}

// Implementación implícita del metodo Area
func (r Rectangulo) Area() float64 {
	return r.Ancho * r.Alto
}

// Esctructura circulo
type Ciculo struct {
	Radio float64
}

// Implementacion implicita del metodo area
func (c Ciculo) Area() float64 {
	return math.Pi * c.Radio * c.Radio
}

// Funcion que acepta la interfaz, permitiendo polimorfismo
func ImprimirArea(f figura) {
	fmt.Printf("El area es: %.2f\n", f.Area())
}

func main() {
	r := Rectangulo{3, 4}
	c := Ciculo{5}

	ImprimirArea(r)
	ImprimirArea(c)
}
