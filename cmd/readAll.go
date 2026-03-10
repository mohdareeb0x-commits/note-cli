/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	logic "note-cli/logic"
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
		err := logic.ReadNotes()

		if err != nil {
			fmt.Println(err.Error())
			fmt.Println()
		} else {
			fmt.Println("\033[34mDONE: All Notes Readed.\n\033[0m")
		}

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
