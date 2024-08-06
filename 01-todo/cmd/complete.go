package cmd

import (
	"fmt"
	"zareix/goprojects/01-todo/internal/db"

	"github.com/spf13/cobra"
)

var completeCmd = &cobra.Command{
	Use:   "complete",
	Short: "Complete a todo",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		todoId := args[0]

		db.CompleteTodo(todoId)

		fmt.Println("Completed todo :", todoId)
	},
}

func init() {
	rootCmd.AddCommand(completeCmd)
}
