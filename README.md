# todo-cli
todo app for the cli

```
Usage of todo:
  -a    Gets all Todos not only from today. Example: todo -a
  -c string
        Creates a new Todo for today with the description provided. Example: todo -c "review today's PRs"
  -d int
        Deletes a Todo from the list. Example: todo -d 8
  -m string
        Adds a note to a Todo when used together with -u. Example: todo -u 8 -m "he's only available after 3PM"
  -n    Gets Todos with their note's information. Example: todo -n
  -u int
        Updates a Todo's status to the next logic one, unless when used with -m, that updates the Todo's note instead. Example: todo -u 8
```
