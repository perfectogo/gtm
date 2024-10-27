package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/perfectogo/gtm/template-manager/core"
)

func main() {
	if len(os.Args) == 0 {
		log.Fatalln(errors.New("not enough arguments"))
	}

	switch os.Args[1] {
	case "init":
		{

			projectName := promptInput("Enter project name: ")
			modPath := promptInput("Enter project mod path: ")

			var crudName *string
			{
				os.Stdout.Write([]byte("\nDo you want to create a new CRUD [y/n]: "))
				reader := bufio.NewReader(os.Stdin)
				input, err := reader.ReadString('\n')
				if err != nil {
					fmt.Println("Error reading input:", err)
					os.Exit(1)
				}

				if strings.ToLower(strings.TrimSpace(input)) == "y" {
					crudName = new(string)
					*crudName = promptInput("Enter CRUD name: ")
				}
			}

			core.CreateApp(projectName, modPath, crudName)
		}
	case "-ac", "--add-crud":
		{
			crudName := promptInput("Enter CRUD name: ")
			fmt.Println(crudName)
		}
	default:
		{
			log.Fatalln(errors.New("unknown flag"))
		}
	}
}
func promptInput(prompt string) string {
	fmt.Print(prompt)

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Error reading input:", err)
		os.Exit(1)
	}

	return strings.TrimSpace(input)
}
