/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	logic "note-cli/logic"
	"github.com/spf13/cobra"
)

// readByIDCmd represents the readByID command
var readByIDCmd = &cobra.Command{
	Use:   "readByID",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		noteID, _ := cmd.Flags().GetInt("id")
		logic.ReadNotesID(noteID)
	},
}

func init() {
	rootCmd.AddCommand(readByIDCmd)
	readByIDCmd.Flags().Int("id", 0, "ID at which title is to be updated")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// readByIDCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// readByIDCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
