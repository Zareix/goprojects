package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"
	"zareix/goprojects/01-todo/internal/db"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all todos",
	Run: func(cmd *cobra.Command, args []string) {
		all, _ := cmd.Flags().GetBool("all")

		todos := db.GetTodos()

		w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
		defer w.Flush()
		fmt.Fprintf(w, "Id\tTitle\tCreated\tDone\n")
		for _, todo := range todos {
			if all || !todo.Done {
				fmt.Fprint(w, todo)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.Flags().BoolP("all", "a", false, "List all todos")
}
