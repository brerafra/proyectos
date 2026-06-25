package main

import "log/slog"

func main() {
	//imprimir un mensaje informativo
	slog.Info("Servidor iniciado exitosamente", "puerto", 8080)

	//imprimir un mensaje de error con datos estructurados
	slog.Error("No se puede conectar a la bade de datos", "error", "timeout", "intento", 3000)
}
