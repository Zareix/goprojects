/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"zareix/goprojects/01-todo/internal/db"

	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new todo",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		title := args[0]
		db.CreateTodo(title)
		fmt.Println("Created todo :", title)
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
}
