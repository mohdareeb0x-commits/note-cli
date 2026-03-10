/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	logic "note-cli/logic"

	"github.com/spf13/cobra"
)

// addNoteCmd represents the addNote command
var addNoteCmd = &cobra.Command{
	Use:   "addNote",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		noteTitle, _ := cmd.Flags().GetString("title")
		noteDescription, _ := cmd.Flags().GetString("description")
		err := logic.AddNote(noteTitle, noteDescription)

		if err != nil {
			fmt.Println(err, noteTitle)
			fmt.Println()
		}

		fmt.Println("\033[32mDONE: Note Added Successfully!\n\033[0m")

	},
}

func init() {
	rootCmd.AddCommand(addNoteCmd)
	addNoteCmd.Flags().String("title", "", "Note Title")
	addNoteCmd.Flags().String("description", "", "Note desciption")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addNoteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addNoteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
