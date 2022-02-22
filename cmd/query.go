/*
Copyright © 2022 Benjamin McAdams mcadams.benj@gmail.com

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/benmcgit/jedit/pkg/filter"
	"github.com/benmcgit/jedit/pkg/parser"
	"github.com/spf13/cobra"
)

var filterSlice []string

// queryCmd represents the query command
var queryCmd = &cobra.Command{
	Use:   "query",
	Short: "Isolates object(s) in your dataset",
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
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) > 0 {
			return fmt.Errorf("query command does not accept arguments. Provided arguments: %v", args)
		}
		err := filter.ValidateFilters(filterSlice)
		if err != nil {
			return err
		}
		return nil
	},	
	Run: func(cmd *cobra.Command, args []string) {
		logs := parser.ParseStdin(os.Stdin)
		filters := filter.GetFilters(filterSlice)
		logs.Filter(filters)
		logs.Print()
	},
}

func init() {
	rootCmd.AddCommand(queryCmd)
	flags := queryCmd.Flags()
	flags.StringSliceVarP(&filterSlice, "filter", "f", []string{}, "Reduce the data set. Acceptable operators: ==, !=, <=, >=, >, <")
	cobra.MarkFlagRequired(flags, "filter")
}
