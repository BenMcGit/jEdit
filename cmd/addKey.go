/*
Copyright © 2022 Benjamin McAdams mcadams.benj@gmail.com

*/
package cmd

import (
	"os"

	"github.com/benmcgit/jedit/pkg/parser"
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

Examples:
  cat example.json | ./jedit addKey priority high
  cat example.json | ./jedit addKey severity 10 --replace`,
	Args: cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		key, value := args[0], args[1]
		replace, err := cmd.Flags().GetBool("replace")
		if err != nil {
			return err
		}

		filters, err := parser.ParseFilters(filterSlice)
		if err != nil {
			return err
		}

		logs, err := parser.ParseStdin(os.Stdin)
		if err != nil {
			return err
		}

		logs.Add(key, value, !replace, filters)
		logs.Print()
		return nil
	},
}

func init() {
	rootCmd.AddCommand(addKeyCmd)
	flags := addKeyCmd.Flags()
	flags.BoolP("replace", "r", false, "If key already exists, replace the value with the new value")
	flags.StringSliceVarP(&filterSlice, "filter", "f", []string{}, "Reduce the data set. Acceptable operators: ==, !=, <=, >=, >, <")
}
