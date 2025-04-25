package cmd

import (
	"fmt"
	"strconv"
	
	"github.com/spf13/cobra"
	"github.com/marewore/task-tracker/internal/task"
)

var delCmd = &cobra.Command {
	Use: "delete",
	Short: "Delete a task",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return fmt.Errorf("Missing task id")
		}
		
		ID, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			return err
		}

		return task.DeleteTask(ID)
	},
}

func init() {
	rootCmd.AddCommand(delCmd)
}