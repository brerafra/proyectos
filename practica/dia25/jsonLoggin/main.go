package main

import (
	"log/slog"
	"os"
)

//En ambientes de producción es muy común enviar los registros a herramientas
//como cloudwatch, dataDog o Grafana, esto se logra fácilmente cambiando el handler a json

func main() {
	//Crear un manejador que imprima en formato json hacia la salida standar
	handler := slog.NewJSONHandler(os.Stdout, nil)

	//Crea un logger personalizado
	logger := slog.New(handler)

	//Registrar un evento, que se imprimirá como una línea json structurada
	logger.Info("Usuario inició sesión", "userId", 42, "status", "activo")
}
