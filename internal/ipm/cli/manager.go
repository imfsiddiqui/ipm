// Package cli provides command-line interface utilities for the IPM application.
package cli

import (
	"github.com/spf13/cobra"
)

// SetupManagerCommands sets up the manager commands and their subcommands.
//
// Parameters:
//   - rootCmd: The root command to which the manager commands will be added.
//   - configDir: The directory containing the configuration files for package managers.
//   - schemaFile: The path to the JSON schema file used for validation.
//
// This function performs the following steps:
//  1. Creates the manager command.
//  2. Adds the validate command to the manager command.
//  3. Adds the enable command to the manager command.
//  4. Adds the disable command to the manager command.
//  5. Adds the list command to the manager command.
//  6. Adds the generate command to the manager command.
//  7. Adds the delete command to the manager command.
//  8. Adds the manager command to the root command.
func SetupManagerCommands(rootCmd *cobra.Command, configDir string, schemaFile string) {
	// Create the manager command
	var managerCmd = &cobra.Command{
		Use:   "manager",
		Short: "Manage package manager configurations",
	}

	// Add the validate command to the manager command
	AddValidateCommand(managerCmd, configDir, schemaFile)

	// Add the enable command to the manager command
	AddEnableCommand(managerCmd, configDir)

	// Add the disable command to the manager command
	AddDisableCommand(managerCmd, configDir)

	// Add the list command to the manager command
	AddListCommand(managerCmd, configDir)

	// Add the generate command to the manager command
	AddGenerateCommand(managerCmd, configDir)

	// Add the delete command to the manager command
	AddDeleteCommand(managerCmd, configDir)

	// Add the manager command to the root command
	rootCmd.AddCommand(managerCmd)
}
