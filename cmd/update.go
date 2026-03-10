/*
Copyright © 2026 note-cli mohdareeb0x@gmail.com
*/
package cmd

import (
	"fmt"
	logic "note-cli/logic"

	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update an existing note",
	Long: `Update the title and/or description of an existing note by its ID.
You can update just the title, just the description, or both.

Examples:
  note-cli update --id 1 --title "New Title"
  note-cli update --id 1 --description "New Description"
  note-cli update --id 1 --title "New Title" --description "New Description"`,
	Run: func(cmd *cobra.Command, args []string) {

		noteID, _ := cmd.Flags().GetInt("id")
		newTitle, _ := cmd.Flags().GetString("title")
		newDesc, _ := cmd.Flags().GetString("description")
		err := logic.Update(noteID, newTitle, newDesc)

		if err != nil {
			fmt.Println(err.Error())
			fmt.Println()
		} else {
			fmt.Println("\033[34mDONE: Updated Notes \n\033[0m")
		}
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
	updateCmd.Flags().String("title", "", "New Title for the given ID")
	updateCmd.Flags().Int("id", 0, "ID at which title is to be updated")
	updateCmd.Flags().String("description", "", "New Description for the given ID")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// updateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// updateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
