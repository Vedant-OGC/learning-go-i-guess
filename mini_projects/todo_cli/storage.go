package todo_cli

import (
	"encoding/json"
	"os"
)

func Save(path string, list *TodoList) error {
	data, err := json.Marshal(list)
	if err != nil {
		return err
	}
	// save to file, with error checking
	return os.WriteFile(path, data, 0644)
}

func Load(path string) (*TodoList, error) {
	var list TodoList
	data, err := os.ReadFile(path)
	if err != nil {
		return &list, err
	}
	err = json.Unmarshal(data, &list)
	return &list, err
}
