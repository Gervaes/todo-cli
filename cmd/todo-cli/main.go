package main

import (
	"bufio"
	"fmt"
	"os"
	"todo-cli/internal/models"
	"todo-cli/internal/storage"
)

func displayTodos(flags models.Flags) {
	store := storage.NewStorage()
	todos := store.GetTodos(flags.GetAllTodos)

	if flags.GetAllTodos {
		fmt.Printf("<< All todos (%d)>>\n", len(todos))
	} else {
		fmt.Printf("<< Today's todos (%d)>>\n", len(todos))
	}

	for _, todo := range todos {
		fmt.Println(todo.ToString(flags))
	}
}

func main() {
	flags := models.NewFlags()
	store := storage.NewStorage()

	if flags.HasNewDescription() {
		err := store.CreateTodo(flags.NewDescription)

		if err != nil {
			fmt.Printf("Error trying to create Todo: %s\n\n", err.Error())
		}
	} else if flags.HasIdToUpdate() {
		todo, err := store.GetTodo(flags.IdToUpdate)

		if err != nil {
			fmt.Printf("Error trying to update Todo: %s\n\n", err.Error())
		}

		if flags.HasNewNote() {
			todo.UpdateNote(flags.NewNote)
		} else {
			todo.UpdateStatus()
		}

		err = store.UpdateTodo(todo)

		if err != nil {
			fmt.Printf("Error trying to update Todo: %s\n\n", err.Error())
		}
	} else if flags.HasIdToDelete() {
		fmt.Printf("Are you sure you want to delete a Todo? [y/N]: ")
		reader := bufio.NewReader(os.Stdin)
		opt, _, err := reader.ReadRune()
		if err != nil {
			fmt.Printf("Error trying to delete Todo: %s\n\n", err.Error())
		}

		if opt == 'y' || opt == 'Y' {
			err = store.DeleteTodo(flags.IdToDelete)

			if err != nil {
				fmt.Printf("Error trying to delete Todo: %s\n\n", err.Error())
			}
		}
	}

	displayTodos(flags)
}
