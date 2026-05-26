package todo_cli

import "errors"

type Todo struct {
	ID        int    `json:"id"`
	Text      string `json:"text"`
	Completed bool   `json:"completed"`
}

type TodoList struct {
	Todos []Todo
}

func (l *TodoList) Add(text string) error {
	if text == "" {
		return errors.New("text cannot be empty")
	}
	id := len(l.Todos) + 1
	l.Todos = append(l.Todos, Todo{ID: id, Text: text, Completed: false})
	return nil
}

func (l *TodoList) Complete(id int) bool {
	for i, t := range l.Todos {
		if t.ID == id {
			l.Todos[i].Completed = true
			return true
		}
	}
	return false
}
