/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	logic "note-cli/logic"
)

// updateDescCmd represents the updateDesc command
var updateDescCmd = &cobra.Command{
	Use:   "updateDesc",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		noteID, _ := cmd.Flags().GetInt("id")
		newDesc, _ := cmd.Flags().GetString("description")
		err := logic.UpdateDesc(noteID, newDesc)

		if err != nil {
			fmt.Println(err, noteID)
			fmt.Println()
		} else {
			fmt.Println("\033[34mDONE: Updated Description at id:", noteID, "\n\033[0m")
		}
	},
}

func init() {
	rootCmd.AddCommand(updateDescCmd)
	updateDescCmd.Flags().Int("id", 0, "ID at which title is to be updated")
	updateDescCmd.Flags().String("description", "", "New Description for the given ID")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// updateDescCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// updateDescCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
