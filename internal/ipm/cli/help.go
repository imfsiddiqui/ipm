// Package cli provides command-line interface utilities for the IPM application.
package cli

import "github.com/spf13/cobra"

// UpdateHelpCommand updates the help command for the root command.
//
// This function replaces the default help command with a hidden "no-help" command.
// The purpose of this function is to disable the default help command and prevent
// it from being displayed in the CLI.
//
// Parameters:
//   - rootCmd: The root command for which the help command will be updated.
//
// Example usage:
//
//	rootCmd := &cobra.Command{Use: "ipm"}
//	cli.UpdateHelpCommand(rootCmd)
//
// This function performs the following steps:
//  1. Creates a new "no-help" command.
//  2. Sets the command to be hidden from the help output.
//  3. Replaces the default help command with the "no-help" command.
func UpdateHelpCommand(rootCmd *cobra.Command) {
	// Create the "no-help" command
	noHelpCmd := &cobra.Command{
		Use:    "no-help",
		Hidden: true,
	}

	// Replace the default help command with the "no-help" command
	rootCmd.SetHelpCommand(noHelpCmd)
}
