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

	flag.StringVar(&flags.NewDescription, "c", "", `Creates a new Todo for today with the description provided. Example: todo -c "review today's PRs"`)
	flag.StringVar(&flags.NewNote, "m", "", `Adds a note to a Todo when used together with -u. Example: todo -u 8 -m "he's only available after 3PM"`)
	flag.IntVar(&flags.IdToUpdate, "u", 0, `Updates a Todo's status to the next logic one, unless when used with -m, that updates the Todo's note instead. Example: todo -u 8`)
	flag.IntVar(&flags.IdToDelete, "d", 0, `Deletes a Todo from the list. Example: todo -d 8`)
	flag.BoolVar(&flags.GetAllTodos, "a", false, `Gets all Todos not only from today. Example: todo -a`)
	flag.BoolVar(&flags.ShowTodosNote, "n", false, `Gets Todos with their note's information. Example: todo -n`)
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
