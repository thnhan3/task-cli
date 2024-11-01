package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "task",
	Short: "Task is a CLI for managing your TODOs.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to task cli")
	},
}

func init() {
	rootCmd.AddCommand(createJsonFile)
	rootCmd.AddCommand(readjsonFile)
	rootCmd.AddCommand(addTaskCommand)
	rootCmd.AddCommand(updateCommand)
	rootCmd.AddCommand(deleteCommand)
	rootCmd.AddCommand(listCommand)
	rootCmd.AddCommand(markDoneCommand)
	rootCmd.AddCommand(markInProcessCommand)
	rootCmd.AddCommand(clearTaskCommand)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
