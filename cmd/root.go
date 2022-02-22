/*
Copyright © 2022 Benjamin McAdams mcadams.benj@gmail.com

*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// support filter across commands
var filterSlice []string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "jedit <command>",
	Short: "Edits a list of JSON objects in a user provided dataset",
	Long: `Parsing and editing JSON in bulk can be time-consuming
and difficult. jEdit aims to help engineers save time by providing them 
a tool that can reduce, filter, and modify their existing JSON dataset.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
