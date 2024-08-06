package cmd

import (
	"fmt"
	"zareix/goprojects/01-todo/internal/db"

	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove a todo",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if db.RemoveTodo(args[0]) {
			fmt.Println("Removed todo with id", args[0])
		} else {
			fmt.Println("Todo not found with id", args[0])
		}
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
}
