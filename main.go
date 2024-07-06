package main

import (
	"fmt"
)

func main() {
	todos := getTodos()

	for _, todo := range todos {
		fmt.Println(todo.ToString())
	}
}
