/*
Copyright © 2022 Benjamin McAdams mcadams.benj@gmail.com

*/
package cmd

import (
	"os"

	"github.com/benmcgit/jedit/pkg/parser"
	"github.com/spf13/cobra"
)

// modifyKeyCmd represents the modifyKey command
var modifyKeyCmd = &cobra.Command{
	Use:   "modifyKey <original_key> <new_key>",
	Short: "Modifies existing keys on object(s) in your dataset",
	Long: `Replaces the name of an existing key with a new name
on each item in the dataset.

If the key does not exist in the dataset, the original dataset will be returned.

Examples:
  cat example.json | ./jedit modifyKey team club
  cat example.json | ./jedit modifyKey severity priority`,
	Args: cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		logs, err := parser.ParseStdin(os.Stdin)
		if err != nil {
			return err
		}

		filters, err := parser.ParseFilters(filterSlice)
		if err != nil {
			return err
		}

		logs.Modify(args[0], args[1], filters)
		logs.Print()
		return nil
	},
}

func init() {
	rootCmd.AddCommand(modifyKeyCmd)
	flags := modifyKeyCmd.Flags()
	flags.StringSliceVarP(&filterSlice, "filter", "f", []string{}, "Reduce the data set. Acceptable operators: ==, !=, <=, >=, >, <")
}
