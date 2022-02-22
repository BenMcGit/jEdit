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
	Run: func(cmd *cobra.Command, args []string) {
		logs := parser.ParseStdin(os.Stdin)
		logs.Remove(args[0])
		logs.Print()
	},
}

func init() {
	rootCmd.AddCommand(removeKeyCmd)
}
