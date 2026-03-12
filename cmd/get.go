/*
Copyright © 2026 note-cli mohdareeb0x@gmail.com
*/
package cmd

import (
	"fmt"
	logic "note-cli/logic"

	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a specific note by ID",
	Long: `Retrieve and display a single note by its ID.
The note's ID, title, and description will be printed to the console.

Example:
  note-cli get --id 1`,

	Run: func(cmd *cobra.Command, args []string) {

		noteID, _ := cmd.Flags().GetInt("id")
		err := logic.GetNotes(noteID)

		if err != nil {
			fmt.Println(err.Error())
			fmt.Println()
		}

	},
}

func init() {
	rootCmd.AddCommand(getCmd)
	getCmd.Flags().Int("id", 0, "ID at which title is to be updated")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
