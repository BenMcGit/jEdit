/*
Copyright © 2022 Benjamin McAdams mcadams.benj@gmail.com

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/benmcgit/jedit/pkg/jedit"
	"github.com/spf13/cobra"
)

// support filter across commands
var filterSlice []string

// support input & output flags for all commands
var inputFileLocation string
var outputFileLocation string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "jedit <command>",
	Short: "Edits a list of JSON objects in a user provided dataset",
	Long: `Parsing and editing JSON in bulk can be time-consuming
and difficult. jEdit aims to help engineers save time by providing them 
a tool that can reduce, filter, and modify their existing JSON dataset.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&inputFileLocation, "input", "", "Path to JSON file to be parsed")
	rootCmd.PersistentFlags().StringVar(&outputFileLocation, "output", "", "Path to file to write resulting JSON to. If not existent, it will be created.")
}

func getInputFilePath() (string,error) {
	input, err := rootCmd.Flags().GetString("input")
	stat, _ := os.Stdin.Stat()
	if err != nil {
		return "", err
	} else if input != "" {
		return input, nil
	} else if (stat.Mode() & os.ModeCharDevice) == 0 {
		// data is being piped to stdin
		return os.Stdin.Name(), nil
	}
	return "", fmt.Errorf("No input file found. Please provide input using Stdin or the --input flag.")
}

func writeToOutput(logs jedit.Logs) error {
	output, err := rootCmd.Flags().GetString("output")
	if err != nil {
		return err
	} else if output != "" {
		// write to specified output file if defined
		return logs.WriteToFile(output)
	}
	// if not output file is defined print to stdout
	logs.Print()
	return nil
}
