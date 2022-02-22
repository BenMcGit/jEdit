/*
Copyright © 2022 Benjamin McAdams mcadams.benj@gmail.com

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/benmcgit/jedit/pkg/parser"
	"github.com/spf13/cobra"
)

// removeKeyCmd represents the removeKey command
var removeKeyCmd = &cobra.Command{
	Use:   "removeKey",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
Args: func(cmd *cobra.Command, args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("removeKey command requires 1 argument. Provided arguments: %v", args)
	}
	return nil
},
Run: func(cmd *cobra.Command, args []string) {
	logs := parser.ParseStdin(os.Stdin)
	logs.Remove(args[0])
	logs.Print()
},
}

func init() {
	rootCmd.AddCommand(removeKeyCmd)
}
