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

// addKeyCmd represents the addKey command
var addKeyCmd = &cobra.Command{
	Use:   "addKey",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 2 {
			return fmt.Errorf("addKey command requires 2 arguments. Provided arguments: %v", args)
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		logs := parser.ParseStdin(os.Stdin)
		key, value := args[0], args[1]
		retain, _ := cmd.Flags().GetBool("retain")
		logs.SortBy(args[0], retain)
		logs.Add(key, value, retain)
		logs.Print()
	},
}

func init() {
	rootCmd.AddCommand(addKeyCmd)
	addKeyCmd.Flags().BoolP("retain", "r", false, "Sort in ascending order")
}
