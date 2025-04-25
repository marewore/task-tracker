package cmd

import (	
	"github.com/spf13/cobra"
	"github.com/marewore/task-tracker/internal/task"
)

var listCmd = &cobra.Command {
	Use: "list",
	Short: "List every task",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) > 0 {
			return task.ListTasks(args[0])
		}

		return task.ListTasks()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}