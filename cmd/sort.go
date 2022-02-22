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

// sortCmd represents the sort command
var sortCmd = &cobra.Command{
	Use:   "sort",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("sort command only accepts 1 argument. Provided arguments: %v", args)
		}
		return nil
	},
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
