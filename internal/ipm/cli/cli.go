// Package cli provides command-line interface utilities for the IPM application.
package cli

import (
	"os"

	"github.com/spf13/cobra"
)

// InitializeCLI initializes the CLI commands and structure.
//
// Parameters:
//   - configDir: The directory containing the configuration files for package managers.
//   - schemaFile: The path to the JSON schema file used for validation.
//   - cliCmd: The name of the command-line interface (CLI) application.
//
// This function performs the following steps:
//  1. Creates the root command for the CLI application.
//  2. Updates the completion command.
//  3. Updates the help command.
//  4. Sets up the manager commands and their subcommands.
//  5. Check required arguments for the validation command.
//  6. Sets up the default manager commands based on the OS.
//  7. Sets up dynamic manager commands based on the configuration files.
//  8. Executes the root command.
func InitializeCLI(configDir string, schemaFile string, cliCmd string) {
	// Create the root command for the CLI application
	var rootCmd = &cobra.Command{Use: cliCmd}

	// Update the completion command
	UpdateCompletionCommand(rootCmd)

	// Update the help command
	UpdateHelpCommand(rootCmd)

	// Set up the manager commands and their subcommands
	SetupManagerCommands(rootCmd, configDir, schemaFile)

	// Check required arguments for the validation command
	firstArg := ""
	secondArg := ""
	if len(os.Args) == 3 {
		firstArg = os.Args[1]
		secondArg = os.Args[2]
	}

	// Set up the default manager commands based on the OS
	SetupDefaultManagerCommands(rootCmd, configDir, schemaFile, os.Args, firstArg, secondArg)

	// Set up dynamic manager commands based on the configuration files
	SetupDynamicManagerCommands(rootCmd, configDir, schemaFile, os.Args, firstArg, secondArg)

	// Execute the root command
	rootCmd.Execute()
}
