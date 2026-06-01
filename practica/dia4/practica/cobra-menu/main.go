package main

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "myapp",
		Short: "Mi aplicacion interactiva",
	}

	var newCmd = &cobra.Command{
		Use:   "new",
		Short: "Crear nuevo elemento con selección interactiva",
		Run: func(cmd *cobra.Command, args []string) {
			//uso de survey para selección
			var language string
			prompt := &survey.Select{
				Message: "Selecciona un lenguaje:",
				Options: []string{"Go", "Python", "JavaScript"},
			}
			survey.AskOne(prompt, &language)
			fmt.Printf("Has seleccionado: %s\n", language)
		},
	}

	rootCmd.AddCommand(newCmd)
	rootCmd.Execute()
}
