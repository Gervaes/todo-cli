package main

import (
	"flag"
	"fmt"
)

func displayTodos() {
	todos := getTodos()

	for _, todo := range todos {
		fmt.Println(todo.ToString())
	}
	fmt.Println("...")
}

func main() {
	var newTodo string
	var idToUpdate int
	flag.StringVar(&newTodo, "c", "", "Creates a new todo for today with the description provided")
	flag.IntVar(&idToUpdate, "u", 0, "Updates a todo's status to the next logic one")
	flag.Parse()

	if newTodo != "" {
		err := createTodo(newTodo)

		if err != nil {
			fmt.Printf("Error trying to create Todo: %s\n\n", err.Error())
		}
	} else if idToUpdate != 0 {
		todo, err := getTodo(idToUpdate)

		if err != nil {
			fmt.Printf("Error trying to update Todo: %s\n\n", err.Error())
		}

		todo.UpdateStatus()
		err = updateTodo(todo)

		if err != nil {
			fmt.Printf("Error trying to update Todo: %s\n\n", err.Error())
		}
	}

	displayTodos()
}
