/*
*********************************************************************

# Practica día 8

8.3 Trabajo con interfaces

Sección 2: figuras geométricas

Utilizando el trabajo previo con interfaces, se genera un programa general
para realizar ejercicios con figuras geometricas
********************************************************************
*/

package main

import (
	"fmt"
	"math"
)

type Figura interface {
	Area() float64
	Perimetro() float64
}

// ============== Circulo ==============================
type Circulo struct {
	Radio float64
}

func (c Circulo) Area() float64 {
	return math.Pi * math.Pow(c.Radio, 2)
}

func (c Circulo) Perimetro() float64 {
	return 2 * math.Pi * c.Radio
}

// ============== Cuadrado ==============================
type Cuadrado struct {
	Alto  float64
	Ancho float64
}

func (r Cuadrado) Area() float64 {
	return r.Ancho * r.Alto
}

func (r Cuadrado) Perimetro() float64 {
	return 2 * (r.Ancho + r.Alto)
}

// ============== Triangulo ==============================
type Triangulo struct {
	Base   float64
	LadoA  float64
	LadoB  float64
	Altura float64
}

func (t Triangulo) Area() float64 {
	return t.Base * t.Altura / 2
}

func (t Triangulo) Perimetro() float64 {
	return (t.Base + t.LadoA + t.LadoB)
}

func main() {
	//Lista de figuras
	figuras := []Figura{
		Circulo{Radio: 5.0},
		Cuadrado{Ancho: 4, Alto: 6},
		Triangulo{Base: 10, LadoA: 7, LadoB: 3, Altura: 10},
	}

	for _, f := range figuras {
		fmt.Printf("Tipo de figura: %T\n", f)
		fmt.Printf("Area: %.2f\n", f.Area())
		fmt.Printf("Perimetro: %.2f\n\n", f.Perimetro())
	}
}
