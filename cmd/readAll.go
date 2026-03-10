/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	logic "note-cli/logic"
	"github.com/spf13/cobra"
)

// readAllCmd represents the readAll command
var readAllCmd = &cobra.Command{
	Use:   "readAll",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		logic.ReadNotes()
	},
}

func init() {
	rootCmd.AddCommand(readAllCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// readAllCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// readAllCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
