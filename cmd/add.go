/*
Copyright © 2026 note-cli mohdareeb0x@gmail.com
*/
package cmd

import (
	"fmt"
	logic "note-cli/logic"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new note",
	Long: `Add a new note to the notes collection with a title and description.
The note will be assigned an auto-incremented ID.

Example:
  note-cli add --title "Meeting Notes" --description "Discussed project timeline"`,
	Run: func(cmd *cobra.Command, args []string) {

		noteTitle, _ := cmd.Flags().GetString("title")
		noteDescription, _ := cmd.Flags().GetString("description")
		err := logic.AddNote(noteTitle, noteDescription)

		if err != nil {
			fmt.Println(err.Error())

		} else {
			fmt.Println("\n\033[32mDONE: Note Added Successfully!\n\033[0m")
		}

	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().String("title", "", "Note Title")
	addCmd.Flags().String("description", "", "Note desciption")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
