/*
*********************************************************************

API día 5

5.2 API Crear un request logger

# Para testear este programa se pude usar curl de la siguiente manera

- curl http://localhost:8085/hello -> para ver el logger


*********************************************************************/

package main

import (
	"log"
	"net/http"
	"time"
)

// responseWriter es un wrapper para capturar el codigo de estado http
type ResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (rw *ResponseWriter) writeHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

// logginMiddleware envuelve un hanlder para registrar los detalles de la solicitud
func logginMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		wrapped := &ResponseWriter{ResponseWriter: w, statusCode: http.StatusOK}

		//Ejecutar el handler original
		next.ServeHTTP(w, r)

		//registrar los detalles de completar la solicitud
		duration := time.Since(start)
		log.Printf(
			"%s %s %d %s",
			r.Method,
			r.RequestURI,
			wrapped.statusCode,
			duration,
		)
	})
}

func main() {
	//Handler de ejemplo
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hola mundo"))
	})

	//aplicar el middleware al servidor http
	addr := ":8085"
	log.Printf("Servidor escuchando en %s", addr)
	if err := http.ListenAndServe(addr, logginMiddleware(http.DefaultServeMux)); err != nil {
		log.Fatal(err)
	}
}
