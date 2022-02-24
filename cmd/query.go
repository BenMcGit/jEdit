/*
Copyright © 2022 Benjamin McAdams mcadams.benj@gmail.com

*/
package cmd

import (
	"github.com/benmcgit/jedit/pkg/jedit"
	"github.com/spf13/cobra"
)

// queryCmd represents the query command
var queryCmd = &cobra.Command{
	Use:   "query",
	Short: "Reduces the amount of object(s) in your dataset",
	Long: `The query command allows the user to reduce the amount of objects 
in their JSON data set. 

IMPORTANT: You must provide at least 1 filter to use the query command.

The following filter operations are supported:
  ==, !=, <=, >=, >, <

You may provide more than 1 filter to reduce the data set even further. 

Examples:
  cat example.json | ./jedit query --filter "team == team-x"
  cat example.json | ./jedit query --filter "team != team-x"
  cat example.json | ./jedit query --filter "severity < 5"
  cat example.json | ./jedit query --filter "team == team-x" --filter "severity < 5"`,
	Args: cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		filters, err := jedit.ParseFilters(filterSlice)
		if err != nil {
			return err
		}

		inputFilePath, err := getInputFilePath()
		if err != nil {
			return err
		}

		logs, err := jedit.ParseFile(inputFilePath)
		if err != nil {
			return err
		}

		logs.Filter(filters)
		return writeToOutput(logs)
	},
}

func init() {
	rootCmd.AddCommand(queryCmd)
	flags := queryCmd.Flags()
	flags.StringSliceVarP(&filterSlice, "filter", "f", []string{}, "Filter out objects in dataset. Acceptable operators: ==, !=, <=, >=, >, <")
	cobra.MarkFlagRequired(flags, "filter")
}
