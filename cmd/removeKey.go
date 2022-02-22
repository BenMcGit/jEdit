/*
Copyright © 2022 Benjamin McAdams mcadams.benj@gmail.com

*/
package cmd

import (
	"os"

	"github.com/benmcgit/jedit/pkg/parser"
	"github.com/spf13/cobra"
)

// removeKeyCmd represents the removeKey command
var removeKeyCmd = &cobra.Command{
	Use:   "removeKey <key>",
	Short: "Removes existing keys on object(s) in your dataset",
	Long: `Removes a key-value pair to each object in the provided dataset.

If the key under consideration for removal is not found on an object
the object will be returned without modification. 

Examples:
  cat example.json | ./jedit removeKey team
  cat example.json | ./jedit removeKey severity`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		key := args[0]
		logs, err := parser.ParseStdin(os.Stdin)
		if err != nil {
			return err
		}

		filters, err := parser.ParseFilters(filterSlice)
		if err != nil {
			return err
		}

		logs.Remove(key, filters)
		logs.Print()
		return nil
	},
}

func init() {
	rootCmd.AddCommand(removeKeyCmd)
	flags := removeKeyCmd.Flags()
	flags.StringSliceVarP(&filterSlice, "filter", "f", []string{}, "Reduce the data set. Acceptable operators: ==, !=, <=, >=, >, <")
}
