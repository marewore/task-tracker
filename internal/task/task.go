package task

import (
	"fmt"
)

type Task struct {
	ID          int64  `json:"id"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

func AddTask(description string) error {
	tasks, err := ReadTasksFromFile()
	if err != nil {
		return err
	}

	id := int64(len(tasks) + 1)
	task := Task{
		ID:          id,
		Description: description,
		Status:      "todo",
	}

	tasks = append(tasks, task)

	fmt.Printf("Task added successfully (ID=%d) \n", id)

	return WriteTasksToFile(tasks)
}

func DeleteTask(id int64) error {
	tasks, err := ReadTasksFromFile()
	if err != nil {
		return err
	}

	var updatedTasks []Task
	found := false
	for _, task := range tasks {
		if task.ID == id {
			found = true
			continue
		}
		updatedTasks = append(updatedTasks, task)
	}

	if !found {
		return fmt.Errorf("no task found (ID=%d) \n", id)
	}

	fmt.Printf("Task deleted successfully (ID=%d) \n", id)

	return WriteTasksToFile(updatedTasks)
}

func UpdateDescription(id int64, description string) error {
	tasks, err := ReadTasksFromFile()
	if err != nil {
		return err
	}

	var updatedTasks []Task
	found := true
	for _, task := range tasks {
		if task.ID == id {
			found = true
			task.Description = description
		}
		updatedTasks = append(updatedTasks, task)
	}

	if !found {
		return fmt.Errorf("no task found (ID=%d) \n", id)
	}

	fmt.Printf("Task updated successfully (ID=%d) \n", id)

	return WriteTasksToFile(updatedTasks)
}

func UpdateStatus(id int64, status string) error {
	tasks, err := ReadTasksFromFile()
	if err != nil {
		return err
	}

	var updatedTasks []Task
	found := false
	for _, task := range tasks {
		if task.ID == id {
			task.Status = status
			found = true
		}
		updatedTasks = append(updatedTasks, task)
	}

	if !found {
		return fmt.Errorf("no task found (ID=%d)", id)
	}

	fmt.Printf("Task status updated successfully (ID=%d, Status=%s)\n", id, status)

	return WriteTasksToFile(updatedTasks)
}

func ListTasks(status ...string) error {
	tasks, err := ReadTasksFromFile()
	if err != nil {
		return err
	}

	if len(status) > 0 && status[0] != "" {
		found := false

		for _, task := range tasks {
			if task.Status == status[0] {
				fmt.Printf("%-30s | %-6d | %-6s\n", task.Description, task.ID, task.Status)
				found = true
			}
		}

		if !found {
			fmt.Println("No tasks found with status:", status[0])
		}

		return nil
	}

	for _, task := range tasks {
		fmt.Printf("%-30s | %-6d | %-6s\n", task.Description, task.ID, task.Status)
	}

	return nil
}