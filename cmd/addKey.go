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
	Run: func(cmd *cobra.Command, args []string) {
		logs := parser.ParseStdin(os.Stdin)
		key, value := args[0], args[1]
		replace, _ := cmd.Flags().GetBool("replace")
		logs.Add(key, value, !replace)
		logs.Print()
	},
}

func init() {
	rootCmd.AddCommand(addKeyCmd)
	addKeyCmd.Flags().BoolP("replace", "r", false, "If key already exists, replace the value with the new value")
}
