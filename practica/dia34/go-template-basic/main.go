package main

import (
	"html/template"
	"log"
	"net/http"
)

type DatosUsuario struct {
	Nombre string
	Cargo  string
}

func main() {
	//1. parsear la plantilla HTML
	//Asume que tienes un archivo llamado "index.html" en el mismo directorio
	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		log.Fatal(err)
	}

	//2. Definir el manejador de rutas
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//Datos estaticos o dinamicos (ej. desde una base de datos)
		datos := DatosUsuario{
			Nombre: "Brenthon",
			Cargo:  "Desarrollador Go",
		}

		//3. Ejecutar la plantilla con los datos
		err := tmpl.Execute(w, datos)
		if err != nil {
			log.Panicln("Error ejecutando plantilla: ", err)
		}
	})

	//4. Iniciar el servidor en el purto 8085
	log.Println("Servidor iniciado en http://localhost:8085")
	log.Fatal(http.ListenAndServe(":8085", nil))
}
