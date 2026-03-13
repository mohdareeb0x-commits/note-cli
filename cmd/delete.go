/*
Copyright © 2026 note-cli mohdareeb0x@gmail.com
*/
package cmd

import (
	"fmt"
	logic "note-cli/logic"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete notes",
	Long: `Delete one or all notes from the collection.
Use --id to delete a specific note by its ID, or --a to delete all notes.

Examples:
  note-cli delete --id 1
  note-cli delete --a`,

	Run: func(cmd *cobra.Command, args []string) {

		aFlag, _ := cmd.Flags().GetBool("a")
		noteID, _ := cmd.Flags().GetInt("id")

		if aFlag {
			err := logic.DeleteAll()

			if err != nil {
				fmt.Println()
				fmt.Println(err.Error())
				fmt.Println()

			} else {
				fmt.Println("\n\033[33mDONE: Deleted All Notes\n\033[0m")
			}


		} else {
			err := logic.DeleteByID(noteID)

			if err != nil {
				fmt.Println()
				fmt.Println(err.Error())
				fmt.Println()

			} else {
				fmt.Println("\033[33mDONE: Deleted Note at id:", noteID, "\n\033[0m")
			}
		}

	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
	deleteCmd.Flags().Bool("a", false, "To delete all")
	deleteCmd.Flags().Int("id", 0, "To delete by id")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
