/*
Copyright © 2022 Benjamin McAdams mcadams.benj@gmail.com

*/
package cmd

import (
	"github.com/benmcgit/jedit/pkg/jedit"
	"github.com/spf13/cobra"
)

// addKeyCmd represents the addKey command
var addKeyCmd = &cobra.Command{
	Use:   "addKey <key> <value>",
	Short: "Adds an additional key to object(s) in your dataset",
	Long: `Appends a key-value pair to each object in the provided dataset.

If the key already exists in at least one entry in the dataset, the original 
value for that key will be retained by default.

To assure data is replaced, use the "replace" flag. 

You may use the "filter" flag to add a key-value pair only if 
the object matches the filter criteria.

Examples:
  cat example.json | ./jedit addKey priority high
  cat example.json | ./jedit addKey severity 10 --replace
  cat example.json | ./jedit addKey severity 10 --filter "team == team-x"`,
	Args: cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		key, value := args[0], args[1]
		replace, err := cmd.Flags().GetBool("replace")
		if err != nil {
			return err
		}

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

		logs.Add(key, value, !replace, filters)
		return writeToOutput(logs)
	},
}

func init() {
	rootCmd.AddCommand(addKeyCmd)
	flags := addKeyCmd.Flags()
	flags.BoolP("replace", "r", false, "If key already exists, replace the value with the new value")
	flags.StringSliceVarP(&filterSlice, "filter", "f", []string{}, "Apply operation to subset of data. Acceptable operators: ==, !=, <=, >=, >, <")
}
