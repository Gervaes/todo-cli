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
	flag.StringVar(&newTodo, "c", "", "Creates a new todo for today with the description provided")
	flag.Parse()

	if newTodo != "" {
		err := createTodo(newTodo)

		if err != nil {
			fmt.Printf("Error trying to create Todo: %s\n================\n", err.Error())
		}
	}

	displayTodos()
}
