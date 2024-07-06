package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func displayTodos(getAllTodos bool) {
	todos := getTodos(getAllTodos)

	if getAllTodos {
		fmt.Printf("<< All todos (%d)>>\n", len(todos))
	} else {
		fmt.Printf("<< Today's todos (%d)>>\n", len(todos))
	}

	for _, todo := range todos {
		fmt.Println(todo.ToString())
	}
}

func main() {
	var newTodo string
	var idToUpdate int
	var idToDelete int
	var getAllTodos bool
	flag.StringVar(&newTodo, "c", "", "Creates a new todo for today with the description provided")
	flag.IntVar(&idToUpdate, "u", 0, "Updates a todo's status to the next logic one")
	flag.IntVar(&idToDelete, "d", 0, "Deletes a todo from the list")
	flag.BoolVar(&getAllTodos, "a", false, "Gets all todos not only from today")
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
	} else if idToDelete != 0 {
		fmt.Printf("Are you sure you want to delete a Todo? [y/N]: ")
		reader := bufio.NewReader(os.Stdin)
		opt, _, err := reader.ReadRune()
		if err != nil {
			fmt.Printf("Error trying to delete Todo: %s\n\n", err.Error())
		}

		if opt == 'y' || opt == 'Y' {
			err = deleteTodo(idToDelete)

			if err != nil {
				fmt.Printf("Error trying to delete Todo: %s\n\n", err.Error())
			}
		}
	}

	clearTerminal()
	displayTodos(getAllTodos)
}
