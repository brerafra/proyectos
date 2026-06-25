package main

import "log/slog"

func main() {
	//configurar el nivel de debu para mostrar absolutamente todos lo smensajes
	slog.SetLogLoggerLevel(slog.LevelDebug)
	slog.Debug("mensaje de depuración detallado")
	slog.Info("mensaje general del sistema")
	slog.Warn("Cuidado, espacio en disco bajo")
	slog.Error("Ocurrió un error critico")
}
