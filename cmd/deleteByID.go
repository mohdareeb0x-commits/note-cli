/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	logic "note-cli/logic"

	"github.com/spf13/cobra"
)

// deleteByIDCmd represents the deleteByID command
var deleteByIDCmd = &cobra.Command{
	Use:   "deleteByID",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		noteID, _ := cmd.Flags().GetInt("id")
		err := logic.DeleteByID(int(noteID))

		if err != nil {
			fmt.Println(err, noteID)
			fmt.Println()
		}

		fmt.Println("\033[33mDONE: Deleted Note at id:", noteID, "\n\033[0m")

	},
}

func init() {
	rootCmd.AddCommand(deleteByIDCmd)
	deleteByIDCmd.Flags().Int("id", 0, "ID at which title is to be updated")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteByIDCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteByIDCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
