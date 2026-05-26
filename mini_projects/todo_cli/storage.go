package todo_cli

import (
	"encoding/json"
	"fmt"
	"os"
)

func Save(path string, list *TodoList) error {
	data, err := json.Marshal(list)
	if err != nil {
		return fmt.Errorf("failed to serialize: %w", err)
	}
	return os.WriteFile(path, data, 0644)
}

func Load(path string) (*TodoList, error) {
	var list TodoList
	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return &list, nil // return empty list if file doesn't exist
		}
		return &list, fmt.Errorf("failed to read file: %w", err)
	}
	err = json.Unmarshal(data, &list)
	if err != nil {
		return &list, fmt.Errorf("failed to parse JSON: %w", err)
	}
	return &list, nil
}
