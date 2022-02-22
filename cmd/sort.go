/*
Copyright © 2022 Benjamin McAdams mcadams.benj@gmail.com

*/
package cmd

import (
	"os"

	"github.com/benmcgit/jedit/pkg/parser"
	"github.com/spf13/cobra"
)

// sortCmd represents the sort command
var sortCmd = &cobra.Command{
	Use:   "sort <key>",
	Short: "Sorts objects in your dataset base on a user-provided key",
	Long: `Sorts a json dataset in ascending or decending order based on a user provided key.

If the user-provided key is not present in the dataset, the original dataset will be returned.

By default, the dataset will be returned in descending order. To return the 
dataset in ascending order, use the flag "asc".

Examples:
  cat example.json | ./jedit sort team
  cat example.json | ./jedit sort severity --asc`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		logs := parser.ParseStdin(os.Stdin)
		isAsc, _ := cmd.Flags().GetBool("asc")
		logs.SortBy(args[0], isAsc)
		logs.Print()
	},
}

func init() {
	rootCmd.AddCommand(sortCmd)
	sortCmd.Flags().BoolP("asc", "a", false, "Sort in ascending order")
}
