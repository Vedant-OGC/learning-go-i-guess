package todo_cli

type Todo struct {
	ID        int    `json:"id"`
	Text      string `json:"text"`
	Completed bool   `json:"completed"`
}

type TodoList struct {
	Todos []Todo
}

func (l *TodoList) Add(text string) {
	id := len(l.Todos) + 1
	l.Todos = append(l.Todos, Todo{ID: id, Text: text, Completed: false})
}
