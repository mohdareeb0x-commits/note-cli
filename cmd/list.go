/*
Copyright © 2026 note-cli mohdareeb0x@gmail.com
*/
package cmd

import (
	"fmt"
	logic "note-cli/logic"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all notes",
	Long: `Display all notes in the collection.
Each note will show its ID, title, and description.

Example:
  note-cli list`,
	Run: func(cmd *cobra.Command, args []string) {

		noteList, err := logic.ListNotes()

		if err != nil {
			fmt.Println(err.Error())
			fmt.Println()
		} else {
			fmt.Print(noteList)
			fmt.Println("\033[34mDONE: All Notes Listed.\n\033[0m")
		}

	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
