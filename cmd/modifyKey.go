/*
Copyright © 2022 Benjamin McAdams mcadams.benj@gmail.com

*/
package cmd

import (
	"os"

	"github.com/benmcgit/jedit/pkg/jedit"
	"github.com/spf13/cobra"
)

// modifyKeyCmd represents the modifyKey command
var modifyKeyCmd = &cobra.Command{
	Use:   "modifyKey <original_key> <new_key>",
	Short: "Modifies existing keys on object(s) in your dataset",
	Long: `Replaces the name of an existing key with a new name
on each item in the dataset.

If the key does not exist in the dataset, the original dataset will be returned.

You may use the "filter" flag to modify a key-value pair only if 
the object matches the filter criteria.

Examples:
  cat example.json | ./jedit modifyKey team club
  cat example.json | ./jedit modifyKey severity priority
  cat example.json | ./jedit modifyKey severity priority --filter "team == team-x"`,
	Args: cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		logs, err := jedit.ParseStdin(os.Stdin)
		if err != nil {
			return err
		}

		filters, err := jedit.ParseFilters(filterSlice)
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
	flags.StringSliceVarP(&filterSlice, "filter", "f", []string{}, "Apply operation to subset of data. Acceptable operators: ==, !=, <=, >=, >, <")
}
