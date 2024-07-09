package models

import (
	"fmt"
	"log"
	"time"
)

const (
	StillTodo = iota
	Doing
	Done
)

type Todo struct {
	Id          int8   `json:"id"`
	Description string `json:"description"`
	Date        string `json:"date"`
	Status      int8   `json:"status"`
	Note        string `json:"note"`
}

func (t *Todo) GetStatus() string {
	var status string

	switch t.Status {
	case StillTodo:
		status = "[ Fazer ]"
	case Doing:
		status = "[Fazendo]"
	case Done:
		status = "[ Feito ]"
	}

	return status
}

func (t *Todo) UpdateStatus() {
	var status int8

	switch t.Status {
	case StillTodo:
		status = Doing
	case Doing:
		status = Done
	case Done:
		status = StillTodo
	}

	t.Status = status
}

func (t *Todo) UpdateNote(note string) {
	t.Note = note
}

func (t *Todo) ToString(flags Flags) string {
	date, err := time.Parse("2006-01-02", t.Date)

	if err != nil {
		log.Fatal(err)
	}

	todo := fmt.Sprintf("%-5d << [ %s ] %s >> %s", t.Id, date.Format("02-Jan-2006"), t.GetStatus(), t.Description)
	if flags.ShowTodosNote && t.Note != "" {
		todo += fmt.Sprintf("\n                                   └> Nota: %s", t.Note)
	}

	return todo
}
