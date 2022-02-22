/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/benmcgit/jedit/pkg/parser"
	"github.com/spf13/cobra"
)

// modifyKeyCmd represents the modifyKey command
var modifyKeyCmd = &cobra.Command{
	Use:   "modifyKey",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 2 {
			return fmt.Errorf("modifyKey command requires 2 arguments. Provided arguments: %v", args)
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		logs := parser.ParseStdin(os.Stdin)
		logs.Modify(args[0], args[1])
		logs.Print()
	},
}

func init() {
	rootCmd.AddCommand(modifyKeyCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// modifyKeyCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// modifyKeyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
