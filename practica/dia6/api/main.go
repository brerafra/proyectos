/*
*********************************************************************

# API día 6

5.2 API Crear un request timer middleware

# Para testear este programa se pude usar curl de la siguiente manera

- curl http://localhost:8085/hello -> para ver el logger

********************************************************************
*/
package main

import (
	"log"
	"net/http"
	"time"
)

// TimerMiddleware envuelve un http.Handler para medir el tiempo de respuesta
func TimerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		inicio := time.Now()
		next.ServeHTTP(w, r)
		//Calcula el tiempo transcurrido
		duracion := time.Since(inicio)
		log.Printf("[%s] %s tardó %v", r.Method, r.URL.Path, duracion)
	})
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/data", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(150 * time.Millisecond) //simula procesamiento
		w.Write([]byte("Datos procesados"))
	})

	//envolvemos el mux con nuestro middleware
	loggedMux := TimerMiddleware(mux)

	log.Println("Servidor escuchado en :8085")
	http.ListenAndServe(":8085", loggedMux)
}
