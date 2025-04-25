package cmd

import (
	"fmt"
	
	"github.com/spf13/cobra"
	"github.com/marewore/task-tracker/internal/task"
)

var addCmd = &cobra.Command {
	Use: "add",
	Short: "Add a new task",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return fmt.Errorf("Missing task description")
		}

		return task.AddTask(args[0])
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}