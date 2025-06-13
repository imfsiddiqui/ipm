// Package cli provides command-line interface utilities for the IPM application.
package cli

import "github.com/spf13/cobra"

// UpdateCompletionCommand adds the completion command to the root command.
//
// This function creates a hidden "completion" command that generates
// the autocompletion script for the specified shell. The completion
// command is useful for enabling shell autocompletion for the CLI.
//
// Parameters:
//   - rootCmd: The root command to which the completion command will be added.
//
// Example usage:
//
//	rootCmd := &cobra.Command{Use: "ipm"}
//	cli.UpdateCompletionCommand(rootCmd)
//
// This function performs the following steps:
//  1. Creates a new "completion" command.
//  2. Sets the command to be hidden from the help output.
//  3. Adds the "completion" command to the root command.
func UpdateCompletionCommand(rootCmd *cobra.Command) {
	// Create the "completion" command
	completion := &cobra.Command{
		Use:   "completion",
		Short: "Generate the autocompletion script for the specified shell",
	}

	// Set the command to be hidden from the help output
	completion.Hidden = true

	// Add the "completion" command to the root command
	rootCmd.AddCommand(completion)
}
