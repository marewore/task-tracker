package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tasktracker",
	Short: "A simple and efficient task tracker CLI tool",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(args)
	},
}


func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}