package models

import "flag"

type Flags struct {
	NewDescription string
	IdToUpdate     int
	IdToDelete     int
	GetAllTodos    bool
}

func NewFlags() Flags {
	flags := Flags{}

	flag.StringVar(&flags.NewDescription, "c", "", "Creates a new todo for today with the description provided")
	flag.IntVar(&flags.IdToUpdate, "u", 0, "Updates a todo's status to the next logic one")
	flag.IntVar(&flags.IdToDelete, "d", 0, "Deletes a todo from the list")
	flag.BoolVar(&flags.GetAllTodos, "a", false, "Gets all todos not only from today")
	flag.Parse()

	return flags
}

func (f *Flags) HasNewDescription() bool {
	return f.NewDescription != ""
}

func (f *Flags) HasIdToUpdate() bool {
	return f.IdToUpdate != 0
}

func (f *Flags) HasIdToDelete() bool {
	return f.IdToDelete != 0
}
