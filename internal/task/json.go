package task

import (
	"encoding/json"
	"errors"
	"os"
)

const fileName = "tasks.json"

func ReadTasksFromFile() ([]Task, error) {
	if _, err := os.Stat(fileName); errors.Is(err, os.ErrNotExist) {
		if err := os.WriteFile(fileName, []byte("[]"), 0644); err != nil {
			return nil, err
		}
		return []Task{}, nil
	}

	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var tasks []Task
	if err := json.NewDecoder(file).Decode(&tasks); err != nil {
		return nil, err
	}

	return tasks, nil
}

func WriteTasksToFile(tasks []Task) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	return json.NewEncoder(file).Encode(tasks)
}