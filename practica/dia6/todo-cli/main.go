/*
*********************************************************************

# Proyecto día 6

5.2 APP Crear un mini banco cli

# En esta aplicación interactiva, es una aplicación TODO donde
el "status" = s - > iniciado
			  t - > terminado
			  e - > borrado
******************************************************************/

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

type Todo struct {
	Id       int64  `json:"id"`
	Activity string `json:"activity"`
	Status   string `json:"status"`
}

var todos []Todo

func main() {
	//var actividad string

	data, err := ioutil.ReadFile("todo.js")
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(data, &todos)

	var funcion string
	fmt.Println("Programa CLI tipo TODO")
	for {
		fmt.Printf("Ingresa una opcion: 1.Listar, 2.Agregar, 3.Borrar 4.Guardar: ")
		fmt.Scan(&funcion)
		fmt.Println()

		switch funcion {
		case "1":
			fmt.Println("Lista de actividades activas guardadas")
			for _, k := range todos {
				if k.Status == "s" {
					fmt.Printf("id:%d :%s\n", k.Id, k.Activity)
				}
			}
		case "2":
			fmt.Printf("Agregar una actividad: ")
			reader := bufio.NewReader(os.Stdin)
			//fmt.Scanln(&actividad)
			actividad, _ := reader.ReadString('\n')
			actividad = strings.TrimSuffix(actividad, "\n")
			todo := Todo{}

			if len(todos) == 0 {
				todo.Id = 0
			} else {
				todo.Id = todos[len(todos)-1].Id + 1
			}

			todo.Activity = actividad
			todo.Status = "s"

			todos = append(todos, todo)
			fmt.Println("actividad agregada correctamente")

		case "3":
			fmt.Printf("borrar una actividad\n")
			fmt.Printf("ingrese id: ")
			var id string
			todo := Todo{}
			fmt.Scan(&id)

			id_int, err := strconv.ParseInt(id, 10, 64)
			if err != nil {
				fmt.Println("Error:", err)
				break

			}
			for _, k := range todos {
				if k.Id == id_int {
					todo = k
				}
			}

			todos[todo.Id].Status = "e"

		case "4":
			fmt.Printf("Guardando todas las actividades.\n")
			jsonData, err := json.Marshal(todos)
			if err != nil {
				log.Fatal(err)
			}
			os.WriteFile("todo.js", jsonData, 0644)

		default:
			fmt.Println("Función no valida")
		}

	}
}
