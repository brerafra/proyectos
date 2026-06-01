/*
*********************************************************************

practica día 5

1.1 API Crear un middleware

# Para testear este programa se pude usar curl de la siguiente manera

- curl http://localhost:8085/foo -> para handler sin middleware
- culr http://localhost:8085/bar -> para handler con middleware

*********************************************************************
*/
package main

import (
	"log"
	"net/http"
)

func middlewareOne(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path, "excuting middlewareOne")
		next.ServeHTTP(w, r)
		log.Println(r.URL.Path, "executing middlewareOne again")
	})
}

func fooHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path, "executing fooHandler")
	w.Write([]byte("OK"))
}

func barHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path, "executing barHandler")
	w.Write([]byte("OK"))
}

func main() {
	mux := http.NewServeMux()

	mux.Handle("GET /foo", http.HandlerFunc(fooHandler))
	mux.Handle("GET /bar", middlewareOne(http.HandlerFunc(barHandler)))
	err := http.ListenAndServe(":8085", mux)
	log.Fatal(err)
}
