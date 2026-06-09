/*
*********************************************************************

practica 1 semana 2

1.2 Creación de un sistema de tareas (TIPO TODO) CLI

Caracteristicas:

- Golang
- SQLITE3

esta aplicación puede crear y listar tareas así como
eliminar una tarea por ID.

*********************************************************************
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	var operacion string
	var input string

	fmt.Println("Programa Tareas CLI")
	fmt.Println("Se puede crear, litar y eliminar tareas ")

	for {
		fmt.Printf("Ingresa opcion 1.Crear 2.Listar 3.Eliminar: ")
		fmt.Scan(&operacion)
		switch operacion {
		case "3":
			fmt.Printf("borrar una actividad\n")
			fmt.Printf("ingrese id: ")
			var id string
			fmt.Scan(&id)

			id_int, err := strconv.ParseInt(id, 10, 64)
			if err != nil {
				fmt.Println("Error:", err)
				break
			}
			err = delete(id_int)
			if err != nil {
				fmt.Println("tarea no existe")
				break
			}
			fmt.Println("tarea eliminada")
		case "2":
			t := new(Todo)
			todos, _ := t.GetTodos()

			fmt.Println("Lista de actividades activas guardadas")
			for _, k := range todos {
				fmt.Printf("id:%d :%s,%s,%s\n", k.ID, k.Title, k.Description, k.UpdatedAt)
			}
		case "1":
			fmt.Printf("Agregar una actividad: ")
			reader := bufio.NewReader(os.Stdin)
			title, _ := reader.ReadString('\n')
			title = strings.TrimSuffix(title, "\n")

			fmt.Printf("Agregar descripción: ")
			reader = bufio.NewReader(os.Stdin)
			description, _ := reader.ReadString('\n')
			description = strings.TrimSuffix(description, "\n")

			fmt.Printf("Tipo 1.Escuela, 2.Trabajo, 3.Deporte: ")
			fmt.Scan(&input)

			kind, err := strconv.ParseInt(input, 10, 10)
			if err != nil {
				fmt.Println("Error:", err)
				break

			}

			todo := Todo{}
			todo.Title = title
			todo.Description = description
			todo.Kind = int(kind)
			todo.CreatedAt = time.Now()
			fmt.Println(todo.CreatedAt)

			err = todo.create()
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				break
			}
			fmt.Println("tarea creada correctamente.")

		default:
			fmt.Println("Opcion invalida")
		}
	}
}
