package main

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

func (t *Todo) ToString() string {
	date, err := time.Parse("2006-01-02", t.Date)

	if err != nil {
		log.Fatal(err)
	}

	return fmt.Sprintf("%-5d << [ %s ] %s >> %s", t.Id, date.Format("02-Jan-2006"), t.GetStatus(), t.Description)
}
