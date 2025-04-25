package cmd

import (
	"fmt"
	"strconv"
	
	"github.com/spf13/cobra"
	"github.com/marewore/task-tracker/internal/task"
)

var update = &cobra.Command {
	Use: "update",
	Short: "Update a task",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 2 {
			return fmt.Errorf("Missing id or description")
		}

		ID, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			return err
		}

		desc := args[1]
		return task.UpdateDescription(ID, desc)
	},
}

var statusProg = &cobra.Command {
	Use: "mark-in-progress",
	Short: "Mark a task as in-progress",
	RunE: func(cmd *cobra.Command, args []string) error {
		return RunUpdateStatus(args, "in-progress")
	},
}

var statusDone = &cobra.Command {
	Use: "mark-done",
	Short: "Mark a task as done",
	RunE: func(cmd *cobra.Command, args []string) error {
		return RunUpdateStatus(args, "done")
	},
}

var statusTodo = &cobra.Command {
	Use: "mark-todo",
	Short: "Mark a task as todo",
	RunE: func(cmd *cobra.Command, args []string) error {
		return RunUpdateStatus(args, "todo")
	},
}

func RunUpdateStatus(args []string, status string) error {
	if len(args) == 0 {
		return fmt.Errorf("Missing task ID")
	}

	ID, err := strconv.ParseInt(args[0], 10, 64)
	if err != nil {
		return err
	}

	return task.UpdateStatus(ID, status)
}

func init() {
	rootCmd.AddCommand(update)
	rootCmd.AddCommand(statusProg)
	rootCmd.AddCommand(statusDone)
	rootCmd.AddCommand(statusTodo)
}