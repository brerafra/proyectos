/*
*********************************************************************

# Practica día 8

8.2 Trabajo con interfaces

Sección 2: poliformismo - desacoplamiento (patrón de diseño).

Las interfaces son buenas para separar la lógica de negocio de los detalles de
la implementación, ejemplo: systemas de pago, conexión a db
********************************************************************
*/
package main

import "fmt"

// Contrato para procesar pagos
type ProcesarPago interface {
	Pagar(monto float64) error
}

// implementacion 1: pago por paypal
type PayPal struct {
	Email string
}

func (p PayPal) Pagar(monto float64) error {
	fmt.Printf("Pagando $%.2f usando la cuenta paypal: %s\n", monto, p.Email)
	return nil
}

// implmentación 2: pago con tarjeta de credito
type TarjetaCredito struct {
	Numero string
}

func (t TarjetaCredito) Pagar(monto float64) error {
	fmt.Printf("Pagando $%.2f con la tarjeta terminada en: %s\n", monto, t.Numero)
	return nil
}

// la funcion recibe la interz, no le importa si es pyalpal o tarjeta
func ReralizarCompra(p ProcesarPago, monto float64) {
	err := p.Pagar(monto)
	if err != nil {
		fmt.Println("Error al procesar el pago.")
	}
}

func main() {
	miPaypal := PayPal{Email: "usuario@dominio.com"}
	miTarjeta := TarjetaCredito{Numero: "1234123412341234"}

	ReralizarCompra(miPaypal, 295.00)
	ReralizarCompra(miTarjeta, 1999.00)
}
