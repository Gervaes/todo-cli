package models

import "flag"

type Flags struct {
	NewDescription string
	NewNote        string
	IdToUpdate     int
	IdToDelete     int
	GetAllTodos    bool
	ShowTodosNote  bool
}

func NewFlags() Flags {
	flags := Flags{}

	flag.StringVar(&flags.NewDescription, "c", "", "Creates a new todo for today with the description provided")
	flag.StringVar(&flags.NewNote, "m", "", "Adds a note to a todo. Used together with -u")
	flag.IntVar(&flags.IdToUpdate, "u", 0, "Updates a todo's status to the next logic one")
	flag.IntVar(&flags.IdToDelete, "d", 0, "Deletes a todo from the list")
	flag.BoolVar(&flags.GetAllTodos, "a", false, "Gets all todos not only from today")
	flag.BoolVar(&flags.ShowTodosNote, "n", false, "Gets todos with notes")
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

func (f *Flags) HasNewNote() bool {
	return f.NewNote != ""
}
